import apiClient from "../api/client";

export type Badge = {
    id: number;
    name: string;
    color: string;
}

export type BadgeWithValue = Badge & {
    value: number
}

export const BadgeService = {
    getMostExpansive: (): Promise<BadgeWithValue[]> => {
        return apiClient.get("/badges/most-expansive");
    }
}
