import apiClient from '../api/client';

interface LoginRequest {
    email: string;
    password: string;
}

interface AuthResponse {
    token: string;
}

interface Register {
    email: string;
    password: string;
}

export const authService = {
    async signIn(credentials: LoginRequest): Promise<AuthResponse> {
        const response = await apiClient.post<AuthResponse>(`/login`, credentials);
        return response;
    },

    async signUp(credentials: Register): Promise<void> {
        console.log(credentials)
        await apiClient.post(`/register`, credentials);
        return;
    },

    logout(): void {
        localStorage.removeItem('authToken');
    }
};
