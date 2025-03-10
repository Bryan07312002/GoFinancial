<template>
    <div class="flex max-h-screen ">
        <sidebar />
        <div class="p-6 max-w-[1440px] mx-auto w-full h-screen">
            <show-paginate-transactions-with-badges @open-transaction="openShow($event)"
                :paginate-options="paginateOptions" :is-loding="isTransactionsLoading" :transactions="transactions" />
        </div>
    </div>
    <modal v-if="selectedTransaction">
        <card class="p-6">
            <show-transaction-with-details @open-update="openUpdate" v-if="modalState == ModalStates.Show"
                @close="selectedTransaction = null" :transaction="selectedTransaction" :is-loading="false" />

            <edit-transaction :update="selectedTransaction?.id" v-if="modalState == ModalStates.Edit" @close="close()"
               :transaction="selectedTransaction" :transactions-with-details="selectedTransaction" />
        </card>
    </modal>
</template>

<script setup lang="ts">
//<card class="max-w-[1440px] overflow-x-auto mx-auto m-6 flex flex-col gap-6 w-full p-6">
import Sidebar from '../components/Sidebar.vue';
import Card from '../components/Card.vue';
import ShowPaginateTransactionsWithBadges from '../components/ShowPaginateTransactionsWithBadges.vue';
import ShowTransactionWithDetails from '../components/ShowTransactionWithDetails.vue';
import NewTransaction from '../components/NewTransaction.vue';
import EditTransaction from '../components/EditTransaction.vue';
import Modal from '../components/Modal.vue';
import { TransactionService, type TransactionWithDetails } from '../services/transactions/transaction';
import { onMounted, ref, type Ref } from 'vue';
import type { PaginateOptions, PaginateResult } from '../services/api/client';

const isTransactionsLoading = ref(true);
const transactions: Ref<
    PaginateResult<TransactionWithDetails
    > | null> = ref(null);

async function getTransactions() {
    isTransactionsLoading.value = true;
    transactions.value = await TransactionService.getPaginateTransactions();
    isTransactionsLoading.value = false;
}

enum ModalStates {
    Show,
    Edit,
    Create,
}

const modalState: Ref<ModalStates> = ref(ModalStates.Show);

const paginateOptions: Ref<PaginateOptions> = ref({

});

onMounted(() => {
    getTransactions()
})

function openCreate(t: TransactionWithDetails) {
    selectedTransaction.value = t;
    modalState.value = ModalStates.Create
}

function openUpdate() {
    console.log("aaaa")
    modalState.value = ModalStates.Edit
}

function openShow(t: TransactionWithDetails) {
    selectedTransaction.value = t;
    modalState.value = ModalStates.Show
}

function close() {
    selectedTransaction.value = null;
}

const selectedTransaction: Ref<TransactionWithDetails | null> = ref(null)

</script>
