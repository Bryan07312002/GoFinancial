import apiClient from "../api/client";
import type { BankAccount } from "../bankAccounts/bankAccounts";
import type { Item, ItemWithBadges } from "../items";

export enum PaymentMethod {
    CreditCard = "credit_card",
    DebitCard = "debit_card",
    Other = "other",
}

export enum TransactionType {
    Income = "income",
    Expense = "expense",
    Transfer = "transfer",
}

export type Transaction = {
    id: number;
    type: TransactionType;
    method: PaymentMethod;
    establishment: string;
    credit: boolean;
    value: number;
    date: Date;
    cardId: number | null;
    bankAccountId: number;
}

export type TransactionWithDetails = Transaction & {
    bank_account: BankAccount;
    items: ItemWithBadges[];
}

export type Badge = {
    id: number;
    name: string;
}

export type RecentResponse = Omit<Transaction, "date"> & {
    badges: Badge[];
    date: string;
}

export type TransactionWithBadges = Transaction & {
    badges: Badge[];
}

export type Balance = {
    balance: number;
    credit: number;
}

export type NewTransactionRequest = {
    type: TransactionType;
    method: PaymentMethod;
    establishment: string;
    credit: boolean;
    value: number;
    date: string;
    bank_account_id: number;
    card_id: number | null;
};

export const TransactionService = {
    create: async (transaction: Omit<Transaction, "id">) => {
        const newTransaction: NewTransactionRequest = {
            type: transaction.type,
            method: transaction.method,
            establishment: transaction.establishment,
            credit: transaction.credit,
            value: transaction.value,
            date: formatDate(transaction.date),
            bank_account_id: transaction.bankAccountId,
            card_id: transaction.cardId,
        };

        await apiClient.post<void>("/transactions", newTransaction);
    },

    getRecent: async (): Promise<TransactionWithBadges[]> => {
        const result = await apiClient
            .get<RecentResponse[]>("/transactions/recent");

        return result.map(el => ({
            ...el,
            date: parseDate(el.date)
        }))
    },

    getById: async (id: number): Promise<TransactionWithDetails> => {
        return apiClient.get(`/transactions/${id}`).then((res: any) => ({
            ...res,
            date: parseDate(res.date)
        }))
    },

    getBalance: async (): Promise<Balance> => {
        const res = await apiClient
            .get<{ balance: string, credit: string }>("/transactions/balance");

        return {
            balance: parseFloat(res.balance),
            credit: parseFloat(res.credit),
        }
    },
}

function parseDate(isoString: string) {
    // Truncate microseconds to milliseconds and handle timezone
    const corrected = isoString.replace(/(\.\d{3})\d+(Z|[+-]\d{2}:\d{2})/, '$1$2');
    return new Date(corrected);
}

function formatDate(date: Date): string {
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
