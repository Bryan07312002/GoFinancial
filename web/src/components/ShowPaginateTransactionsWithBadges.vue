<template>
    <div class="flex flex-col gap-4 h-full">
        <div class="flex justify-between">
            <b class="text-2xl">
                Transactions
            </b>
            <Button text="New Transaction" />
        </div>

        <div class="flex gap-4">
            <drop-down-input :value="''" :options="fields" placeholder="Order by" />

            <div class="flex gap-2 justify-center items-center">
                <drop-down-input :value="''" :options="fields" placeholder="field" />
                <b>:</b>
                <Input :value="''" />
            </div>

            <div class="flex gap-4 items-center">
                <div class="flex gap-2">
                    <calendar-icon class="h-auto text-[var(--neutral-400)]" />
                    <date-picker-input :value="new Date()" class="w-26" placeholder="start" />
                </div>
                <b>to</b>
                <div class="flex gap-2">
                    <calendar-icon class="h-auto text-[var(--neutral-400)]" />
                    <date-picker-input :value="new Date()" class="w-26" placeholder="finish" />
                </div>
                <Button text="Search" />
            </div>
        </div>

        <card class="h-full overflow-y-auto">
            <div class=" flex flex-col shadow-lg h-full rounded-md overflow-y-auto">
                <div @click="emits('openTransaction', transaction)" v-for="transaction, i in transactions?.data"
                    :class="i % 2 == 0 ? '' : 'bg-[var(--neutral-700)]'"
                    class="h-full cursor-pointer flex justify-between p-4 hover:bg-[var(--neutral-700)] rounded-[var(--radius)]">
                    <div class="flex gap-4 items-center">

                        <arrow-up-icon class="w-6 h-6 text-green-300"
                            v-if="transaction.type == TransactionType.Income" />
                        <arrow-down-icon class="w-6 h-6 stroke-red-300 " v-else />

                        <div>
                            <div>{{ transaction.establishment }}</div>
                            <div class="text-[var(--neutral-400)]">{{ formatDate(transaction.date) }}</div>
                        </div>
                    </div>

                    <div class="flex text-[var(--neutral-100)] gap-4 justify-center items-center">
                        <bank-icon class="text-[var(--neutral-400)]" />
                        {{ transaction.bank_account.name }}
                    </div>

                    <div class="flex gap-4 justify-center items-end">
                        <credit-card-icon class="text-[var(--neutral-400)]" v-if="transaction.credit" />
                        <coin-icon class="text-[var(--neutral-400)]" v-else />

                        <div v-if="transaction.type == TransactionType.Income" class="text-green-300">
                            R$ {{ transaction.value }}
                        </div>
                        <div v-else class="text-[var(--des)]"> R$ {{ transaction.value }} </div>
                    </div>
                </div>
            </div>
        </card>
        <div class="mt-auto flex items-center">
            <div v-if="transactions">
                {{ transactions.current_page * transactions.page_size }}
                of
                {{ transactions.total }}
            </div>
            <div class="ml-auto flex gap-2 text-[var(--neutral-100)] items-center">
                <div class="flex text-[var(--neutral-400)] gap-3 items-center">
                    Rows per page:
                    <Input class="w-10 mr-20" :value="'10'" />
                </div>
                <div @click="emits('previousPage')"
                    class="bg-[var(--neutral-700)] border rounded-md text-center w-10 text-2xl p-1 aspect-square border-[var(--secondary-1)]">
                    ←
                </div>
                <div @click="emits('nextPage')"
                    class="bg-[var(--neutral-700)] border rounded-md text-center w-10 text-2xl p-1 aspect-square border-[var(--secondary-1)]">
                    →
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { type TransactionWithDetails, TransactionType } from '../services/transactions/transaction';
import { formatDateShort, type PaginateResult } from '../services/api/client';
import DatePickerInput from './DatePickerInput.vue';
import ArrowUpIcon from '../assets/ArrowUpIcon.vue';
import ArrowDownIcon from '../assets/ArrowDownIcon.vue';
import BankIcon from '../assets/BankIcon.vue';
import CreditCardIcon from '../assets/CreditCardIcon.vue';
import CoinIcon from '../assets/CoinIcon.vue';
import Button from './Button.vue';
import CalendarIcon from '../assets/CalendarIcon.vue';
import DropDownInput from './DropDownInput.vue';
import Input from './Input.vue';
import { type Ref, ref } from 'vue';
import Card from './Card.vue';

defineProps<{
    transactions: PaginateResult<TransactionWithDetails> | null,
    isLoding: boolean,
}>();
const emits = defineEmits(["openTransaction", "nextPage", "previousPage"]);
const fields: Ref<
    (string | { value: string | number, label: string, disabled?: boolean })[]
> = ref([{ label: '', value: '' }])

function formatDate(date: Date) {
    const today = new Date();
    const yesterday = new Date(today);
    yesterday.setDate(today.getDate() - 1);

    if (isSameDay(date, today)) {
        return `Hoje às ${formatTime(date)}`;
    }

    if (isSameDay(date, yesterday)) {
        return `Ontem às ${formatTime(date)}`;
    }

    return formatDateShort(date)
}

function isSameDay(date1: Date, date2: Date): boolean {
    return date1.getDate() === date2.getDate() &&
        date1.getMonth() === date2.getMonth() &&
        date1.getFullYear() === date2.getFullYear();
}

function formatTime(date: Date) {
    return date.toLocaleTimeString('pt-BR', {
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false,
    });
}
</script>
