<template>
    <div class="flex flex-col gap-4">
        <div class="flex gap-4 font-bold ">
            <div class="flex-1">
                <arrow-left-right class=" stroke-primary" />
            </div>
            <div class="flex-8 text-center">New Transaction</div>
            <div v-if="shouldHaveCloseButton" class="flex-1 text-xl hover:text-red-300 cursor-pointer text-end"
                @click="emits('closeButton')">
                X
            </div>
        </div>

        <Input v-model:value="newTransaction.establishment" placeholder="Establishment" />

        <div class="flex items-end">
            <div class="pr-2 font-light">R$</div>
            <Input class="w-full" v-model:value="newTransaction.establishment" placeholder="value" />
            <check-box />
        </div>

        <drop-down-input placeholder="Bank Account" v-model:value="newTransaction.bankAccountId"
            :options="[{ label: 'Hell', value: 1 }]" />

        <div class="flex gap-4 justify-between w-full">
            <Button :is-loading="isLoading" bottom-color-type="secondary" class="flex-1" text="Cancel"
                @click="emits('cancelButton')" />
            <Button :is-loading="isLoading" class="flex-1" text="Save" @click="handleSave" />
        </div>
    </div>

</template>

<script setup lang="ts">
import { ref } from 'vue';
import Input from './Input.vue';
import DropDownInput from './DropDownInput.vue';
import CheckBox from './CheckBox.vue';
import ArrowLeftRight from '../assets/ArrowLeftRight.vue';
import {
    TransactionType,
    PaymentMethod,
    type Transaction,
} from '../services/transactions/transaction';
import type { Ref } from 'vue';

defineProps<{ shouldHaveCloseButton?: boolean }>();

const emits = defineEmits(['closeButton', 'cancelButton']);
const isLoading = ref(false);

// @ts-ignore
const newTransaction: Ref<Transaction> = ref({
    type: TransactionType.Expense,
    method: PaymentMethod.Other,
    establishment: '',
    credit: false,
    value: 0,
    date: Date(),
    cardId: 1,
    bankAccountId: 1,
})

async function handleSave() { }
</script>
