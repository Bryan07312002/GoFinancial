<template>
    <card class="min-w-80 p-4 flex flex-col gap-4">
        <div class="flex justify-between items-center">
            <div class="text-2xl">Recent Activities</div>
            <Button @click="emits('ClickedOnNewBankAccount')" text="New Bank Account" />
        </div>

        <div v-if="isLoading" class="flex items-center justify-center h-full">
            <loading-icon />
        </div>

        <div v-else class="flex flex-col gap-4 h-full">
            <card @click="emits('ClickedOnTransaction', transaction)" v-for="transaction in transactions"
                class="cursor-pointer flex justify-between p-4 hover:brightness-50">
                <div class="flex gap-4">

                    <arrow-up-icon class="w-6 h-6 text-green-300" v-if="transaction.type == TransactionType.Income" />
                    <arrow-down-icon class="w-6 h-6 stroke-red-300 " v-else />

                    <div>
                        <div>{{ transaction.establishment }}</div>
                        <div class="text-primary">{{ formatDate(transaction.date) }}</div>
                    </div>
                </div>

                <div v-if="transaction.type == TransactionType.Income" class="text-green-300"> R$ {{ transaction.value
                    }} </div>
                <div v-else class="text-red-300"> R$ {{ transaction.value }} </div>
            </card>
        </div>
    </card>
</template>

<script setup lang="ts">
import Card from "./Card.vue";
import Button from "./Button.vue";
import ArrowUpIcon from "../assets/ArrowUpIcon.vue";
import LoadingIcon from '../assets/LoadingIcon.vue';
import ArrowDownIcon from "../assets/ArrowDownIcon.vue";
import {
    type Transaction,
    TransactionType,
} from '../services/transactions/transaction';

defineProps<{ transactions: Transaction[], isLoading?: boolean }>();
const emits = defineEmits(["ClickedOnTransaction", "ClickedOnNewBankAccount"]);

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

    return date.toLocaleString('pt-BR', {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
        second: '2-digit',
        hour12: false,
    });
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
