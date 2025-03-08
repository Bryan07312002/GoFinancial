<template>
    <div class="flex max-h-screen max-w-screen">
        <sidebar />
        <div class="p-6 w-full h-screen">
            <show-paginate-transactions-with-badges @open-transaction="selectedTransaction = $event"
                :is-loding="isTransactionsLoading" :transactions="transactions" />
        </div>
    </div>
    <modal v-if="selectedTransaction">
        <card class="p-6">
            <show-transaction-with-details @close="selectedTransaction = null" :transaction="selectedTransaction"
                :is-loading="false" />
        </card>
    </modal>
</template>

<script setup lang="ts">
//<card class="max-w-[1440px] overflow-x-auto mx-auto m-6 flex flex-col gap-6 w-full p-6">
import Sidebar from '../components/Sidebar.vue';
import Card from '../components/Card.vue';
import ShowPaginateTransactionsWithBadges from '../components/ShowPaginateTransactionsWithBadges.vue';
import ShowTransactionWithDetails from '../components/ShowTransactionWithDetails.vue';
import Modal from '../components/Modal.vue';
import { TransactionService, type TransactionWithDetails } from '../services/transactions/transaction';
import { onMounted, ref, type Ref } from 'vue';
import type { PaginateResult } from '../services/api/client';

const isTransactionsLoading = ref(true);
const transactions: Ref<
    PaginateResult<TransactionWithDetails
    > | null> = ref(null);
async function getTransactions() {
    isTransactionsLoading.value = true;
    transactions.value = await TransactionService.getPaginateTransactions();
    isTransactionsLoading.value = false;
}

onMounted(() => {
    getTransactions()
})

const selectedTransaction: Ref<TransactionWithDetails | null> = ref(null)

</script>
