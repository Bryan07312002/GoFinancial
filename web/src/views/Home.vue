<template>
    <div>
        <sidebar />
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
    //Promise.all([
    //    getBalance(),
    //    getBankAccounts(),
    //    getRecentTransactions(),
    //    getMostExpansiveBadges(),
    //])
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
