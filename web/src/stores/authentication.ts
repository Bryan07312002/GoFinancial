import { defineStore } from 'pinia';
import { computed } from 'vue';
import { jwtDecode } from 'jwt-decode';
import Cookies from 'js-cookie';

interface JwtPayload {
    exp: number;
}

export const useAuthStore = defineStore('auth', () => {
    // State
    const token = computed(() => {
        return Cookies.get('jwt')
    });

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
            secure: import.meta.env.VITE_API_ENV === 'production',
            sameSite: 'Strict',
        });
    };

    const clearToken = () => {
        Cookies.remove('jwt');
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
    };
});
