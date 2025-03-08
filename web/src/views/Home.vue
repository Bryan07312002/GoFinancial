<template>
    <div class="flex max-h-screen max-w-screen">
        <sidebar />
        <div class="max-w-[1440px] overflow-x-auto mx-auto flex flex-col gap-6 w-full p-6">
            <div class="flex flex-wrap gap-6 w-full flex-1">
                <card class="flex-1 min-w-[300px] flex flex-col gap-6 justify-between p-6">
                    <div class="flex gap-6 text-[var(--neutral-400)]">
                        <coin-icon />
                        <div>Total Network</div>
                    </div>
                    <div class="text-3xl">R$ {{ balance?.balance + balance?.credit }}</div>
                </card>
                <card class="flex-1 flex min-w-[300px] gap-6 flex-col justify-between p-6">
                    <div class="flex gap-6 text-[var(--neutral-400)]">
                        <wallet-icon />
                        <div>Balance</div>
                    </div>
                    <div class="text-3xl">R$ {{ balance?.balance }}</div>
                </card>
                <card class="flex-1 flex min-w-[200px] flex-col gap-6 justify-between p-6">
                    <div class="flex gap-6 text-[var(--neutral-400)]">
                        <credit-card-icon />
                        <div>Debt</div>
                    </div>
                    <div class="text-3xl">R$ {{ balance?.credit }}</div>
                </card>
            </div>

            <card class="w-full p-6 flex flex-wrap flex-5 gap-6">
                <most-expansive-badges class="flex-1 min-h-[300px]" :badges="mostExpansiveBadges" />
                <recent-transfer-activities @new-transaction="handleOpenNewTransactionModal"
                    @open-transaction="handleOpenTransactionWithDetails($event)" class="flex-1 min-h-[300px]"
                    :transactions="transactions" />
            </card>
        </div>
    </div>
    <modal v-if="modalState.isOpen">
        <card class="p-6">
            <show-transaction-with-details @close="closeModal"
                v-if="modalState.state === ModalState.ShowTransactionWithDetails" :transaction="transactionWithDetails"
                :is-loading="isShowTransactionLoading" />

            <new-transaction  @close="closeModal" @cancel="closeModal" should-have-close-button v-if="modalState.state === ModalState.NewTransaction" />
        </card>
    </modal>
</template>

<script setup lang="ts">
import Modal from "../components/Modal.vue";
import Card from "../components/Card.vue";
import CoinIcon from "../assets/CoinIcon.vue";
import ShowTransactionWithDetails from "../components/ShowTransactionWithDetails.vue";
import { ref, onMounted, type Ref } from "vue";
import Sidebar from "../components/Sidebar.vue";
import WalletIcon from "../assets/WalletIcon.vue"
import CreditCardIcon from "../assets/CreditCardIcon.vue"
import MostExpansiveBadges from "../components/MostExpansiveBadges.vue";
import { BadgeService, type BadgeWithValue } from "../services/badges/badges";
import RecentTransferActivities from "../components/RecentTransferActivities.vue";
import NewTransaction from "../components/NewTransaction.vue";
import {
    type Balance,
    TransactionService,
    type TransactionWithBadges,
    type TransactionWithDetails,
} from "../services/transactions/transaction";

onMounted(async () => {
    Promise.all([
        getBalance(),
        getRecentTransactions(),
        getMostExpansiveBadges(),
    ])
});

enum ModalState {
    NewTransaction,
    ShowTransactionWithDetails,
}

const modalState = ref({
    isOpen: false,
    state: ModalState.NewTransaction,
})

function handleOpenNewTransactionModal() {
    modalState.value.isOpen = true;
    modalState.value.state = ModalState.NewTransaction;
}

function openShowTransactionWithDetails() {
    modalState.value.isOpen = true;
    modalState.value.state = ModalState.ShowTransactionWithDetails;
}

function closeModal() {
    modalState.value.isOpen = false;
}

const isTransactionsLoading = ref(false);
const transactions: Ref<TransactionWithBadges[]> = ref([]);

async function getRecentTransactions() {
    isTransactionsLoading.value = true;
    transactions.value = await TransactionService.getRecent();
    isTransactionsLoading.value = false;
}

const isBalanceLoading = ref(false);
const balance: Ref<Balance | null> = ref(null)

async function getBalance() {
    isBalanceLoading.value = true;
    balance.value = await TransactionService.getBalance();
    isBalanceLoading.value = false;
}

const isMostExpansiveBadgeLoading = ref(false);
const mostExpansiveBadges: Ref<BadgeWithValue[]> = ref([]);
async function getMostExpansiveBadges() {
    isMostExpansiveBadgeLoading.value = true;
    mostExpansiveBadges.value = await BadgeService.getMostExpansive();
    isMostExpansiveBadgeLoading.value = false;
}

const isShowTransactionLoading = ref(false);
const transactionWithDetails: Ref<TransactionWithDetails | null> = ref(null);
async function getTransactionWithDetails(id: number) {
    isShowTransactionLoading.value = true;
    transactionWithDetails.value = await TransactionService.getById(id);
    isShowTransactionLoading.value = false;
}

function handleOpenTransactionWithDetails(id: number) {
    openShowTransactionWithDetails()
    return getTransactionWithDetails(id);
}
</script>

<style>
.apexcharts-toolbar {
    display: none !important;
}
</style>
