import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import Cookies from 'js-cookie';
import { jwtDecode } from 'jwt-decode';

interface JwtPayload {
    exp: number;
    // Add other claims as needed
}

export const useAuthStore = defineStore('auth', () => {
    // State
    const token = ref<string | null>(null);

    // Helper function to validate token
    const isTokenValid = (t: string) => {
        try {
            const decoded = jwtDecode<JwtPayload>(t);
            return decoded.exp * 1000 > Date.now();
        } catch {
            return false;
        }
    };

    // Actions
    const setToken = (newToken: string) => {
        const decoded = jwtDecode<JwtPayload>(newToken);
        const expires = new Date(decoded.exp * 1000);

        Cookies.set('jwt', newToken, {
            expires,
            // FIXME: get real env
            secure: import.meta.env.VITE_API_ENV === 'production',
            sameSite: 'Strict',
        });

        token.value = newToken;
    };

    const clearToken = () => {
        Cookies.remove('jwt');
        token.value = null;
    };

    const initializeFromCookie = () => {
        const cookieToken = Cookies.get('jwt');
        if (cookieToken) {
            if (isTokenValid(cookieToken)) {
                token.value = cookieToken;
            } else {
                clearToken();
            }
        }
    };

    // Getters
    const isAuthenticated = computed(() =>
        !!token.value && isTokenValid(token.value)
    );

    return {
        token,
        isAuthenticated,
        setToken,
        clearToken,
        initializeFromCookie,
    };
});
