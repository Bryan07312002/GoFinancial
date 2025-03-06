<template>
    <div class="flex flex-col gap-4">
        <div class="flex justify-between font-bold">
            <transfer-icon class="w-6 h-6" />
            Transaction
            <div class="text-xl hover:text-red-300 cursor-pointer text-end" @click="emits('close')">
                X
            </div>
        </div>

        <div class="h-[1px] mb-1 w-full bg-[var(--neutral-600)]" />

        <div class="w-full flex flex-col gap-4">
            <div class="flex-10 flex flex-col gap-4">
                <show-field :is-loading="isLoading" name="Type" :value="transaction?.type" />
                <show-field :is-loading="isLoading" name="Method" :value="transaction?.method" />
                <show-field :is-loading="isLoading" name="Establishment" :value="transaction?.establishment" />
                <show-field :is-loading="isLoading" name="Value" :value="transaction?.value" />
                <show-field :is-loading="isLoading" name="Date" :value="formatDateShort(transaction?.date)" />
                <show-field :is-loading="isLoading" name="Bank Account" :value="transaction?.bank_account.name" />
            </div>
            <div class="flex-1 flex gap-4 flex-col">
                <items-table v-if="transaction?.items" :items="transaction.items" />
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import ShowField from './ShowField.vue';
import TransferIcon from '../assets/TransferIcon.vue';
import { type TransactionWithDetails, formatDateShort } from '../services/transactions/transaction';
import ItemsTable from './ItemsTable.vue';

const emits = defineEmits(["close"]);

defineProps<{
    transaction: TransactionWithDetails | null;
    isLoading: boolean;
}>();
</script>
