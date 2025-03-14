import type { BankAccount } from "../bankAccounts/bankAccounts";
import type { ItemWithBadges } from "../items";
import {
    apiClient,
    addPaginationQuery,
    parseDate,
    formatDateForSaving,
    type PaginateResult,
    type PaginateOptionsWithTime,
} from "../api/client";

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
    card_id: number | null;
    bank_account_id: number;
}

export type TransactionWithDetails = Transaction & {
    bank_account: BankAccount;
    items: ItemWithBadges[];
}

export type Badge = {
    id: number;
    name: string;
    color: string;
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
            date: formatDateForSaving(transaction.date),
            bank_account_id: transaction.bank_account_id,
            card_id: transaction.card_id,
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

    getPaginateTransactions: async (options?: PaginateOptionsWithTime): Promise<PaginateResult<TransactionWithDetails>> => {
        return apiClient.get(addPaginationQuery('/transactions', options ?? {}))
            .then((res: any) => ({
                ...res,
                data: res.data.map((transaction: any) => ({
                    ...transaction,
                    date: parseDate(transaction.date),
                    value: parseFloat(transaction.value)
                }))
            }));
    },

    updateTransaction: async (transaction: Transaction) => {
        const newTransaction: NewTransactionRequest = {
            type: transaction.type,
            method: transaction.method,
            establishment: transaction.establishment,
            credit: transaction.credit,
            value: transaction.value,
            date: formatDateForSaving(transaction.date),
            bank_account_id: transaction.bank_account_id,
            card_id: transaction.card_id,
        };

        await apiClient.patch<void>(`/transactions/${transaction.id}`, newTransaction);
    },
}

