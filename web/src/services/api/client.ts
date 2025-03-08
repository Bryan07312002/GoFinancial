import { useAuthStore } from '../../stores/authentication';
import router from '../../router';

export type PaginateResult<T> = {
    data: T[];
    total: number;
    current_page: number;
    page_size: number;
    total_pages: number;
}

export type PaginateOptions = {
    page?: number,
    take?: number,
    sort_by?: string,
    sort_desc?: boolean,
    start?: string,
    finish?: string,
}

export function addPaginationQuery(
    url: string,
    options: PaginateOptions,
): string {
    if (!url.includes('?'))`${url}?`
    if (options.page) url = `${url}&page=${options.page}`
    if (options.take) url = `${url}&take=${options.take}`
    if (options.sort_by) url = `${url}&sort_by=${options.sort_by}`
    if (options.sort_desc) url = `${url}&sort_desc=${options.sort_desc}`
    if (options.finish) url = `${url}&finish=${options.finish}`
    if (options.start) url = `${url}&finish=${options.start}`

    return url
}

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
    router.replace({ name: "Login" });
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

export function parseDate(isoString: string) {
    // Truncate microseconds to milliseconds and handle timezone
    const corrected = isoString.replace(/(\.\d{3})\d+(Z|[+-]\d{2}:\d{2})/, '$1$2');
    return new Date(corrected);
}

export function formatDateForSaving(date: Date): string {
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');
    const milliseconds = String(date.getMilliseconds()).padStart(3, '0');

    const fractionalSeconds = `${milliseconds}000000`;

    const offsetMinutes = date.getTimezoneOffset();
    const sign = offsetMinutes > 0 ? '-' : '+'; // Invert sign for correct timezone representation
    const absOffset = Math.abs(offsetMinutes);
    const offsetHours = String(Math.floor(absOffset / 60)).padStart(2, '0');
    const offsetMinutesPart = String(absOffset % 60).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}.${fractionalSeconds} ${sign}${offsetHours}${offsetMinutesPart}`;
}

export function formatDateShort(date: Date | undefined) {
    if (!date) return ''

    const day = String(date.getDate()).padStart(2, '0');
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const year = date.getFullYear();
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');

    return `${day}/${month}/${year} - ${hours}:${minutes}`;
}

export default apiClient;
