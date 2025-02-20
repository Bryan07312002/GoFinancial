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
    cardId: number;
    bankAccountId: number;
}
