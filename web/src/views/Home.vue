<template>
    <div class="flex max-w-screen">
        <sidebar />
        <div class="max-w-[1440px] mx-auto flex flex-col gap-6 w-full p-6">
            <div>
                <div class="flex gap-6 h-38 w-full">
                    <card class="flex-1 flex flex-col justify-between p-6">
                        <div class="flex gap-6 text-[var(--neutral-400)]">
                            <bank-icon />
                            <div>Total Network</div>
                        </div>
                        <div class="text-4xl">R$ {{ balance?.balance + balance?.credit }}</div>
                    </card>
                    <card class="flex-1 flex flex-col justify-between p-6">
                        <div class="flex gap-6 text-[var(--neutral-400)]">
                            <wallet-icon />
                            <div>Balance</div>
                        </div>
                        <div class="text-4xl">R$ {{ balance?.balance }}</div>
                    </card>
                    <card class="flex-1 flex flex-col justify-between p-6">
                        <div class="flex gap-6 text-[var(--neutral-400)]">
                            <credit-card-icon />
                            <div>Debth</div>
                        </div>
                        <div class="text-4xl">R$ {{ balance?.credit }}</div>
                    </card>
                </div>
            </div>
            <card class="w-full p-6 min-h-1/3 flex gap-6">
                <recent-transfer-activities class="w-1/2" :transactions="transactions" />
                <recent-transfer-activities class="w-1/2" :transactions="transactions" />
            </card>
        </div>
    </div>
</template>

<script setup lang="ts">
import Sidebar from "../components/Sidebar.vue";
import { ref, onMounted, type Ref } from "vue";
import Card from "../components/Card.vue";
import Modal from "../components/Modal.vue";
import BankIcon from "../assets/BankIcon.vue"
import WalletIcon from "../assets/WalletIcon.vue"
import CreditCardIcon from "../assets/CreditCardIcon.vue"
import BankAccountList from "../components/BankAccountList.vue";
import CreateBankAccount from "../components/CreateBankAccount.vue";
import NewTransaction from "../components/NewTransaction.vue";
import LoadingIcon from "../assets/LoadingIcon.vue";
import RecentTransferActivities from "../components/RecentTransferActivities.vue";
import { BankAccountService, type BankAccount } from "../services/bankAccounts/bankAccounts";
import { BadgeService, type BadgeWithValue } from "../services/badges/badges";
import BadgesWithValues from "../components/BadgesWithValues.vue";
import ShowTransactionWithDetails from "../components/ShowTransactionWithDetails.vue";
import {
    type Balance,
    TransactionService,
    type TransactionWithBadges,
    type TransactionWithDetails,
} from "../services/transactions/transaction";

onMounted(async () => {
    Promise.all([
        getBalance(),
        //    getBankAccounts(),
        getRecentTransactions(),
        //    getMostExpansiveBadges(),
    ])
});

enum ModalState {
    NewTransaction,
    CreateBankAccount,
    ShowTransactionWithDetails,
}

const modalState = ref({
    isOpen: false,
    state: ModalState.CreateBankAccount,
})

const isBankAccountsLoading = ref(true);
// @ts-ignore: dont know why but ts i being crazy here
const bankAccounts: Ref<BankAccount[]> = ref([]);

async function getBankAccounts() {
    isBankAccountsLoading.value = true;
    bankAccounts.value = (await BankAccountService.getPaginate(0, 0)).data;
    isBankAccountsLoading.value = false;
}

async function handleCreatedBankAccount() {
    closeModal();
    getBankAccounts();
}

function openCreateBankAccountModal() {
    modalState.value.isOpen = true;
    modalState.value.state = ModalState.CreateBankAccount;
}

function openNewTransactionModal() {
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
