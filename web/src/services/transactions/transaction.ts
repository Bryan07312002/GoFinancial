import apiClient from "../api/client";

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

function formatDate(date: Date): string {
    // Date components
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0'); // Months are 0-based
    const day = String(date.getDate()).padStart(2, '0');
    const hours = String(date.getHours()).padStart(2, '0');
    const minutes = String(date.getMinutes()).padStart(2, '0');
    const seconds = String(date.getSeconds()).padStart(2, '0');
    const milliseconds = String(date.getMilliseconds()).padStart(3, '0');

    // Fractional seconds (milliseconds + 6 zeros to simulate nanoseconds)
    const fractionalSeconds = `${milliseconds}000000`;

    // Timezone offset (±HHMM)
    const offsetMinutes = date.getTimezoneOffset();
    const sign = offsetMinutes > 0 ? '-' : '+'; // Invert sign for correct timezone representation
    const absOffset = Math.abs(offsetMinutes);
    const offsetHours = String(Math.floor(absOffset / 60)).padStart(2, '0');
    const offsetMinutesPart = String(absOffset % 60).padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}.${fractionalSeconds} ${sign}${offsetHours}${offsetMinutesPart}`;
}

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
    }
}
