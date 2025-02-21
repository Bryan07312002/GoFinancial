import apiClient from "../api/client";

export type BankAccount = {
    id: number;
    name: number;
    credit: number;
    debt: number;
}

export type CreateBankAccount = {
    name: string
    description: string
}

export const BankAccountService = {
    getPaginate: async (page: number, take: number) => {
        //page //page_size
        const bankAccounts = await apiClient.get('/bank_accounts');
        return bankAccounts;
    },

    createBankAccount: async (bankAccount: CreateBankAccount) => {
        await apiClient.post('/bank_accounts', bankAccount);
    },

    deleteBankAccount: async (id: number) => {
        await apiClient.post(`/bank_accounts/${id}`);
    },
}
