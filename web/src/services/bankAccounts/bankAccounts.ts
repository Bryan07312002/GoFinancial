import apiClient from "../api/client";

export type BankAccount = {
    id: number;
    name: string;
    credit: number;
    debt: number;
}

export type CreateBankAccount = {
    name: string
    description: string
}
export type PaginateResult<T> = {
    data: T[];
    total: number;
    current_page: number;
    page_size: number;
    total_pages: number;
}

export const BankAccountService = {
    getPaginate: async (page: number, page_size: number): Promise<PaginateResult<BankAccount>> => {
        const bankAccounts = await apiClient.get('/bank_accounts');
        return bankAccounts as any;
    },

    createBankAccount: async (bankAccount: CreateBankAccount) => {
        await apiClient.post('/bank_accounts', bankAccount);
    },

    deleteBankAccount: async (id: number) => {
        await apiClient.post(`/bank_accounts/${id}`);
    },
}
