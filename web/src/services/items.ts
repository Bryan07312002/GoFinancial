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
