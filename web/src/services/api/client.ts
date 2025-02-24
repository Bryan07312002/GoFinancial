import { useAuthStore } from '../../stores/authentication';

type HttpMethod = 'GET' | 'POST' | 'PUT' | 'PATCH' | 'DELETE';

const baseUrl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

const getAuthHeader = () => {
    const authStore = useAuthStore();
    const token = authStore.token;
    return token ? { 'Authorization': `Bearer ${token}` } : {};
};

const handleUnauthorized = () => {
    const authStore = useAuthStore();
    authStore.clearToken();
};

const createRequest = async <T>(
    method: HttpMethod,
    url: string,
    data?: unknown,
    config?: RequestInit
): Promise<T> => {
    try {
        const headers: HeadersInit = {
            'Content-Type': 'application/json',
            ...getAuthHeader(),
            ...config?.headers
        } as any;

        const response = await fetch(`${baseUrl}${url}`, {
            method,
            headers,
            body: data ? JSON.stringify(data) : undefined,
            ...config
        });

        if (!response.ok) {
            if (response.status === 401) {
                handleUnauthorized();
            }

            throw response;
        }

        // Handle empty response
        const text = await response.text();
        return text ? JSON.parse(text) : {} as T;

    } catch (error) {
        if (error instanceof Error && error.message.includes('401')) {
            handleUnauthorized();
        }
        throw error;
    }
};

export const apiClient = {
    get: <T>(url: string, config?: RequestInit) =>
        createRequest<T>('GET', url, undefined, config),

    post: <T>(url: string, data?: unknown, config?: RequestInit) =>
        createRequest<T>('POST', url, data, config),

    put: <T>(url: string, data?: unknown, config?: RequestInit) =>
        createRequest<T>('PUT', url, data, config),

    delete: <T>(url: string, config?: RequestInit) =>
        createRequest<T>('DELETE', url, undefined, config),

    patch: <T>(url: string, data?: unknown, config?: RequestInit) =>
        createRequest<T>('PATCH', url, data, config)
};

export default apiClient;
