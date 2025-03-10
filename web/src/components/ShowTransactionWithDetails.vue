<template>
    <div class="flex flex-col gap-4 min-w-72">
        <div class="flex justify-between font-bold">
            <transfer-icon class="w-6 h-6" />
            Transaction
            <div class="text-xl hover:text-red-300 cursor-pointer text-end" @click="emits('close')">
                X
            </div>
        </div>

        <div class="h-[1px] mb-1 w-full bg-[var(--neutral-600)]" />

        <div class="w-full relative flex flex-col gap-4">
            <div
                @click="emits('openUpdate')"
                class="absolute right-0 ml-auto bg-[var(--primary)] rounded-sm hover:brightness-50 cursor-pointer flex justify-center items-center w-8 h-8">
                <pencil-icon class="w-5" />
            </div>

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
import ItemsTable from './ItemsTable.vue';
import PencilIcon from '../assets/PencilIcon.vue';
import { formatDateShort } from '../services/api/client';
import { type TransactionWithDetails } from '../services/transactions/transaction';

const emits = defineEmits(["close", "openUpdate"]);

defineProps<{
    transaction: TransactionWithDetails | null;
    isLoading: boolean;
}>();
</script>
