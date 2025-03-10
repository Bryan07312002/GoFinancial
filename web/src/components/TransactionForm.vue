<template>
    <div class="flex flex-col gap-4">
        <toggle-group label="Transaction Type" :disabled="isLoading" :options="[
            { name: TransactionType.Income, value: TransactionType.Income },
            { name: TransactionType.Expense, value: TransactionType.Expense }
        ]" v-model:value="transactionsWithDetails.type" />

        <toggle-group label="Payment Method" :disabled="isLoading" :options="[
            { name: 'Credit Card', value: PaymentMethod.CreditCard },
            { name: 'Debit Card', value: PaymentMethod.DebitCard },
            { name: 'Other', value: PaymentMethod.Other }
        ]" v-model:value="transactionsWithDetails.method" />

        <Input :disabled="isLoading" v-model:value="transactionsWithDetails.establishment"
            placeholder="Establishment" />

        <monetary-input :disabled="isLoading" v-model:value="transactionsWithDetails.value" placeholder="value" />

        <lazy-drop-down-input :disabled="isLoading" @search="getBankAccounts($event)" placeholder="Bank Account"
            v-model:value="transactionsWithDetails.bank_account_id" :options="bankAccountsOptions" />

        <date-picker-input :disabled="isLoading" v-model:value="transactionsWithDetails.date" />

        <div class="flex gap-4">
            <check-box :disabled="isLoading" v-model:value="transactionsWithDetails.credit" />
            <div>Credit transaction</div>
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
import {
    TransactionType,
    PaymentMethod,
    type Transaction,
} from '../services/transactions/transaction';
import { BankAccountService, type BankAccount } from '../services/bankAccounts/bankAccounts';

defineProps<{
    transactionsWithDetails: Omit<Transaction, "id">,
}>();

const emits = defineEmits([
    'close',
    'cancel',
    'created',
]);
const isLoading = ref(false);
const bankAccounts: Ref<BankAccount[]> = ref([]);
const bankAccountsOptions = computed(() =>
    bankAccounts.value.map(account => ({ name: account.name, value: account.id })));

onMounted(() => {
    getBankAccounts('');
});

// TODO: apply pagination find here
async function getBankAccounts(search: string) {
    bankAccounts.value = (await BankAccountService.getPaginate(0, 0)).data;
}
</script>
