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

        <transaction-form v-model:transactions-with-details="updated" />

        <div class="">
            {{ toDeleteItems }}
            <items-table v-model:delete-items="toDeleteItems" edit-mode class="mb-4" v-model:new-item="newItem"
                :items="transaction.items" :added-items="itemsToSave" :badge-options="badges" />

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
import Button from './Button.vue';
import ItemsTable from './ItemsTable.vue';
import { ItemService, type ItemWithBadges } from '../services/items';
import TransactionForm from './TransactionForm.vue';
import { ref, onMounted, computed, type Ref } from 'vue';
import ArrowLeftRight from '../assets/ArrowLeftRight.vue';
import { BankAccountService, type BankAccount } from '../services/bankAccounts/bankAccounts';
import { TransactionService, type Badge, type TransactionWithDetails, } from '../services/transactions/transaction';
import { BadgeService } from '../services/badges/badges';

const props = defineProps<{
    shouldHaveCloseButton?: boolean,
    transaction: Omit<TransactionWithDetails, "bank_account" | "id">
    update?: number;
}>();

const updated = ref({ ...props.transaction });

type newItem = Omit<ItemWithBadges, "id" | "transaction_id">

const newItem: Ref<newItem | null> = ref(null);
const itemsToSave: Ref<newItem[]> = ref([]);

const toDeleteItems: Ref<number[]> = ref([]);

const emits = defineEmits([
    'close',
    'cancel',
    'created',
]);

function openNewItem() {
    newItem.value = {
        name: "",
        value: 0,
        quantity: 0,
        badges: [],
    };
}

function handleAddItem() {
    if (newItem.value != null)
        itemsToSave.value.push({ ...newItem.value })

    newItem.value = null
}

const isLoading = ref(false);
const bankAccounts: Ref<BankAccount[]> = ref([]);

const bankAccountsOptions = computed(() =>
    bankAccounts.value.map(account => ({ name: account.name, value: account.id })));

onMounted(() => {
    getBankAccounts("");
    getBadges("");
});

const isBadgeLoading = ref(false);
const badges: Ref<{ name: string, value: Badge }[]> = ref([]);

// TODO: apply pagination find here
async function getBadges(_: string) {
    badges.value = (await BadgeService.getPaginate()).data
        .map(badge => ({ name: badge.name, value: badge }));
}

// TODO: apply pagination find here
async function getBankAccounts(_: string) {
    bankAccounts.value = (await BankAccountService.getPaginate(0, 0)).data;
}

async function handleSave() {
    if (props.update) {
        //await TransactionService.updateTransaction({
        //    id: props.update,
        //    ...updated.value,
        //})

        await ItemService.addMultiplesToTransaction(
            props.update,
            itemsToSave.value,
        )

        toDeleteItems.value.forEach(async item => await ItemService.delete(item))
    }
}
</script>
