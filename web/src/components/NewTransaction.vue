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

        <transaction-form v-model:transactions-with-details="newTransaction" />

        <div class="">
            <items-table v-model:new-item="newItem" :items="items" />
            <div v-if="newItem == null" @click="openNewItem"
                class="mx-auto p-2 w-8 h-8 flex justify-center items-center rounded-lg bg-[var(--primary)]">+</div>
            <div @click="handleAddItem" v-else
                class="mx-auto p-2 w-8 h-8 flex justify-center items-center rounded-lg bg-[var(--primary)]"> v
            </div>
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
import Button from './Button.vue';
import ArrowLeftRight from '../assets/ArrowLeftRight.vue';
import ItemsTable from './ItemsTable.vue';
import TransactionForm from './TransactionForm.vue';
import {
    TransactionType,
    PaymentMethod,
    TransactionService,
    type TransactionWithDetails,
} from '../services/transactions/transaction';
import { BankAccountService, type BankAccount } from '../services/bankAccounts/bankAccounts';
import type { Item, ItemWithBadges } from '../services/items';

defineProps<{
    shouldHaveCloseButton?: boolean,
}>();

const newTransaction: Ref<Omit<TransactionWithDetails, "bank_account" | "id">> = ref({
    type: TransactionType.Expense,
    method: PaymentMethod.DebitCard,
    establishment: '',
    credit: false,
    value: 0,
    date: new Date(),
    cardId: null,
    bankAccountId: 1,
    items: [] as ItemWithBadges[],
});


const newItem: Ref<Omit<Item, "id" | "transaction_id"> | null> = ref(null);

const itemsToSave = ref([]);

const items = computed(() => [
    ...newTransaction.value.items,
    ...itemsToSave.value
])

function openNewItem() {
    newItem.value = {
        name: "",
        value: 0,
        quantity: 0,
    };
}

function handleAddItem() {
    if (newItem.value != null)
        itemsToSave.value.push({ ...newItem.value })
}

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

async function handleSave() {
    isLoading.value = true;
    await TransactionService.create({ ...newTransaction.value });
    isLoading.value = false;

    emits("created")
}
</script>
