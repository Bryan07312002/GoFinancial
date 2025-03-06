<template>
    <div class="flex flex-col gap-4">
        <div class="flex justify-between font-bold">
            <transfer-icon class="w-6 h-6" />
            Transaction
            <div @click="emits('close')" class="text-red-300">X</div>
        </div>

        <div class="w-full h-[1px] bg-[var(--foreground)]" />

        <div class="w-full flex">
            <div class="flex-4 flex flex-col gap-4">
                <show-field :is-loading="isLoading" name="Type" :value="transaction?.type" />
                <show-field :is-loading="isLoading" name="Method" :value="transaction?.method" />
                <show-field :is-loading="isLoading" name="Establishment" :value="transaction?.establishment" />
                <show-field :is-loading="isLoading" name="Value" :value="transaction?.value" />
                <show-field :is-loading="isLoading" name="Date" :value="formatDateShort(transaction?.date)" />
                <show-field :is-loading="isLoading" name="Bank Account" :value="transaction?.bank_account.name" />
            </div>
            <div class="flex-1 flex gap-4 flex-col">
                <div v-for="item in transaction?.items">
                    <div class="flex text-[var(--muted-foreground)] justify-between w-full">
                        <div>{{ item.name }}</div>
                        <div v-if="item.quantity > 1">x{{ item.quantity }}</div>
                        <div>R$ {{ item.value }}</div>
                    </div>

                    <div class="flex gap-2">
                        <div v-for="badge in item.badges"
                            class="p-1 text-center text-ellipsis overflow-hidden text-nowrap border rounded-[var(--radius)]">
                            {{ badge.name }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import ShowField from './ShowField.vue';
import type TransferIcon from '../assets/TransferIcon.vue';
import { type TransactionWithDetails, formatDateShort } from '../services/transactions/transaction';

const emits = defineEmits(["close"]);

defineProps<{
    transaction: TransactionWithDetails | null;
    isLoading: boolean;
}>();


</script>
