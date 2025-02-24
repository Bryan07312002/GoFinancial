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

export type Badge = {
    id: number,
    name: string,
}

export type RecentResponse = Omit<Transaction, "date"> & {
    badges: Badge[]
    date: string
}

export type TransactionWithBadges = Transaction & {
    badges: Badge[]
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

function parseFormattedDate(formattedStr: string): Date {
    console.log(formattedStr)
    // Split the formatted string into its components
    const [dateTimePart, timezonePart] = formattedStr.split(/(?= [+-])/); // Split at the space before timezone
    const [datePart, timePart] = dateTimePart.split(' ');

    // Extract date components
    const [year, month, day] = datePart.split('-').map(Number);

    // Extract time components
    const [time, fractional] = timePart.split('.');
    const [hours, minutes, seconds] = time.split(':').map(Number);
    const milliseconds = parseInt(fractional.substring(0, 3), 10); // Use first 3 digits (ignore nanoseconds)

    // Extract timezone offset
    const sign = timezonePart[0] === '+' ? 1 : -1;
    const offsetHours = parseInt(timezonePart.substring(1, 3), 10);
    const offsetMinutes = parseInt(timezonePart.substring(3, 5), 10);
    const totalOffsetMinutes = sign * (offsetHours * 60 + offsetMinutes);

    // Create a Date object in UTC
    const utcDate = new Date(
        Date.UTC(
            year,
            month - 1, // Months are 0-based in JS
            day,
            hours,
            minutes,
            seconds,
            milliseconds
        )
    );

    // Adjust for the original timezone offset
    return new Date(utcDate.getTime() + totalOffsetMinutes * 60 * 1000);
}

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
    },

    getRecent: async (): Promise<TransactionWithBadges[]> => {
        const result = await apiClient.get<RecentResponse[]>("/transactions/recent");
        return result.map(el => ({
            ...el,
            date: parseFormattedDate(el.date)
        }))
    }
}
