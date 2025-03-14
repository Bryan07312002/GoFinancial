import {
    apiClient,
    addPaginationQuery,
    type PaginateOptions,
    type PaginateResult,
} from "../api/client";

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

export const BankAccountService = {
    getPaginate: async (paginateOpt: PaginateOptions): Promise<PaginateResult<BankAccount>> => {
        const bankAccounts = await apiClient
            .get(addPaginationQuery('/bank_accounts', paginateOpt));

        return bankAccounts as any;
    },

    createBankAccount: async (bankAccount: CreateBankAccount) => {
        await apiClient.post('/bank_accounts', bankAccount);
    },

    deleteBankAccount: async (id: number) => {
        await apiClient.post(`/bank_accounts/${id}`);
    },
}
