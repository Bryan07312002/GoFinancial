<template>
    <card class="min-w-80 p-4 flex flex-col gap-4">
        <div class="flex justify-between items-center">
            <div class="text-2xl">Recent Activities</div>
            <Button @click="emits('newTransaction')" text="New Transaction" />
        </div>

        <div v-if="isLoading" class="flex items-center justify-center h-full">
            <loading-icon />
        </div>

        <div v-else class="flex flex-col shadow-lg rounded-md overflow-hidden h-full">
            <div @click="emits('openTransaction', transaction.id)" v-for="transaction, i in transactions"
                :class="i % 2 == 0 ? '' : 'bg-[var(--neutral-800)]'"
                class="cursor-pointer h-full flex justify-between p-4 hover:bg-[var(--neutral-700)] rounded-[var(--radius)]">
                <div class="flex gap-4 items-center">

                    <arrow-up-icon class="w-6 h-6 text-green-300" v-if="transaction.type == TransactionType.Income" />
                    <arrow-down-icon class="w-6 h-6 stroke-red-300 " v-else />

                    <div>
                        <div>{{ transaction.establishment }}</div>
                        <div class="text-[var(--muted-foreground)]">{{ formatDate(transaction.date) }}</div>
                    </div>
                </div>

                <div class="flex flex-col items-end">
                    <div v-if="transaction.type == TransactionType.Income" class="text-green-300">
                        R$ {{ transaction.value }}
                    </div>
                    <div v-else class="text-[var(--des)]"> R$ {{ transaction.value }} </div>

                    <div class="flex gap-1 max-w-40 overflow-hidden">
                        <badge v-for="badge in transaction.badges" :badge="badge" />
                    </div>
                </div>
            </div>
        </div>
    </card>
</template>

<script setup lang="ts">
import Card from "./Card.vue";
import Button from "./Button.vue";
import ArrowUpIcon from "../assets/ArrowUpIcon.vue";
import LoadingIcon from '../assets/LoadingIcon.vue';
import ArrowDownIcon from "../assets/ArrowDownIcon.vue";
import { formatDateShort } from "../services/transactions/transaction"
import Badge from "./Badge.vue";
import {
    TransactionType,
    type TransactionWithBadges,
} from '../services/transactions/transaction';

defineProps<{ transactions: TransactionWithBadges[], isLoading?: boolean }>();
const emits = defineEmits(["openTransaction", "newTransaction"]);

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
