<template>
    <div class="flex flex-col gap-4">
        <div class="flex gap-4 font-bold ">
            <div class="flex-1">
                <arrow-left-right />
            </div>
            <div class="flex-8 text-center">New Transaction</div>
            <div v-if="shouldHaveCloseButton" class="flex-1 text-xl hover:text-red-300 cursor-pointer text-end"
                @click="emits('close')">
                X
            </div>
        </div>

        <div class="h-[1px] mb-1 w-full bg-[var(--neutral-600)]" />

        <toggle-group label="Transaction Type" :disabled="isLoading" :options="[
            { name: TransactionType.Income, value: TransactionType.Income },
            { name: TransactionType.Expense, value: TransactionType.Expense }
        ]" v-model:value="newTransaction.type" />

        <toggle-group label="Payment Method" :disabled="isLoading" :options="[
            { name: 'Credit Card', value: PaymentMethod.CreditCard },
            { name: 'Debit Card', value: PaymentMethod.DebitCard },
            { name: 'Other', value: PaymentMethod.Other }
        ]" v-model:value="newTransaction.method" />

        <Input :disabled="isLoading" v-model:value="newTransaction.establishment" placeholder="Establishment" />

        <monetary-input :disabled="isLoading" v-model:value="newTransaction.value" placeholder="value" />

        <lazy-drop-down-input :disabled="isLoading" @search="getBankAccounts($event)" placeholder="Bank Account"
            v-model:value="newTransaction.bankAccountId" :options="bankAccountsOptions" />

        <date-picker-input :disabled="isLoading" v-model:value="newTransaction.date" />

        <div class="flex gap-4">
            <check-box :disabled="isLoading" v-model:value="newTransaction.credit" />
            <div>Credit transaction</div>
        </div>

        <div class="flex gap-4 justify-between w-full">
            <Button :is-loading="isLoading" bottom-color-type="secondary" class="flex-1" text="Cancel"
                @click="emits('cancel')" />
            <Button :is-loading="isLoading" class="flex-1" text="Save" @click="handleSave" />
        </div>
    </div>

</template>

<script setup lang="ts">
import { ref, onMounted, computed, type Ref } from 'vue';
import Input from './Input.vue';
import MonetaryInput from './MonetaryInput.vue';
import LazyDropDownInput from './LazyDropDownInput.vue';
import DatePickerInput from './DatePickerInput.vue';
import ToggleGroup from './ToggleGroup.vue';
import CheckBox from './CheckBox.vue';
import Button from './Button.vue';
import ArrowLeftRight from '../assets/ArrowLeftRight.vue';
import {
    TransactionType,
    PaymentMethod,
    type Transaction,
    TransactionService,
} from '../services/transactions/transaction';
import { BankAccountService, type BankAccount } from '../services/bankAccounts/bankAccounts';

defineProps<{ shouldHaveCloseButton?: boolean }>();

const emits = defineEmits(['close', 'cancel', 'created']);
const isLoading = ref(false);
const bankAccounts: Ref<BankAccount[]> = ref([]);
const bankAccountsOptions = computed(() =>
    bankAccounts.value.map(account => ({ name: account.name, value: account.id })));

onMounted(() => getBankAccounts(''));

async function getBankAccounts(search: string) {
    bankAccounts.value = (await BankAccountService.getPaginate(0, 0)).data;
}


// @ts-ignore lsp being crazy
const newTransaction: Ref<Transaction> = ref({
    type: TransactionType.Expense,
    method: PaymentMethod.DebitCard,
    establishment: '',
    credit: false,
    value: 0,
    date: new Date(),
    cardId: null,
    bankAccountId: 1,
})

async function handleSave() {
    isLoading.value = true;
    await TransactionService.create({ ...newTransaction.value });
    isLoading.value = false;

    emits("created")
}
</script>
