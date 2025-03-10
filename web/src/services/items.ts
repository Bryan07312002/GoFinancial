import apiClient from "./api/client";
import type { Badge } from "./transactions/transaction"

export type Item = {
    id: number;
    transaction_id: number;
    name: string;
    value: number;
    quantity: number;
}

export type ItemWithBadges = Item & {
    badges: Badge[];
}

export const ItemService = {
    addMultiplesToTransaction: async (
        transactionId: number,
        items: Omit<ItemWithBadges, "id" | "transaction_id">[],
    ) => {
        const requestBody = {
            transaction_id: transactionId,
            items: items.map(item => ({
                ...item,
                quantity: parseInt(item as any),
                badges: item.badges.map(badge => badge.id),
            }))
        }

        await apiClient.post('/items', requestBody);
    },

    delete: (id: number) => {
        return apiClient.delete(`/items/${id}`);
    }
}

