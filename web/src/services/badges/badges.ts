import apiClient, { addPaginationQuery, type PaginateOptions, type PaginateResult } from "../api/client";

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
    },

    getPaginate: async (
        options?: PaginateOptions,
    ): Promise<PaginateResult<Badge>> => {
        return apiClient.get(addPaginationQuery('/badges', options ?? {}))
            .then((res: any) => ({
                ...res,
                data: res.data.map((badge: any) => ({
                    ...badge,
                }))
            }));
    },
}
