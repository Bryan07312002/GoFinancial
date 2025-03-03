<template>
    <div class="min-h-screen flex">
        <div class="flex flex-col min-h-screen gap-4 p-4 w-full max-w-7xl mx-auto">
            <div class="flex  w-full justify-around gap-4 flex-wrap">

                <card class="flex-1 min-w-32 h-28 p-4 flex flex-col justify-between">
                    <div class="w-full flex items-center justify-between">
                        <div>Total Balance</div>
                        <bank-icon class="w-6 h-6 stroke-primary" />
                    </div>

                    <div class="animate-spin" v-if="isBalanceLoading || balance === null"><loading-icon /></div>
                    <b v-else class="text-xl">R$ {{ balance.balance + balance.credit }}</b>
                </card>

                <card class="flex-1 min-w-32 h-28 p-4 flex flex-col justify-between">
                    <div class="w-full flex items-center justify-between">
                        <div>Total Credit</div>

                        <wallet-icon class="w-6 h-6 stroke-primary" />
                    </div>

                    <div class="animate-spin" v-if="isBalanceLoading || balance === null"><loading-icon /></div>
                    <b v-else class="text-xl text-green-300">R$ {{ balance.balance }}</b>
                </card>

                <card class="flex-1 min-w-32 h-28 p-4 flex flex-col justify-between">
                    <div class="w-full flex items-center justify-between">
                        <div>Total Debt</div>

                        <credit-card-icon class="w-6 h-6 stroke-primary" />
                    </div>

                    <div class="animate-spin" v-if="isBalanceLoading || balance === null"><loading-icon /></div>
                    <b v-else class="text-xl text-red-300">R$ {{ balance.credit }}</b>
                </card>
            </div>

            <div class="flex w-full justify-around gap-4 flex-wrap flex-grow min-h-96 h-1/3">
                <bank-account-list @create-bank-account="openCreateBankAccountModal" :is-loading="isBankAccountsLoading"
                    :bank-accounts="bankAccounts" class="flex-1" />
                <recent-transfer-activities @new-transaction="openNewTransactionModal"
                    :is-loading="isTransactionsLoading" :transactions="transactions" class="flex-1" />
            </div>

            <card class="flex-1 h-1/5 flex-grow p-4 min-h-[200px]">
                <badges-with-values :badges="mostExpansiveBadges"/>
            </card>
        </div>
    </div>
    <modal v-if="modalState.isOpen">
        <card class="p-4 w-full max-w-sm">
            <create-bank-account @created="handleCreatedBankAccount" @close-button="closeModal"
                @cancel-button="closeModal" should-have-close-button
                v-if="modalState.state === ModalState.CreateBankAccount" />

            <new-transaction @close-button="closeModal" @cancel-button="closeModal" should-have-close-button
                v-if="modalState.state === ModalState.NewTransaction" />
        </card>
    </modal>
</template>

<script setup lang="ts">
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
import {
    type Balance,
    TransactionService,
    type TransactionWithBadges,
} from "../services/transactions/transaction";

onMounted(async () => {
    Promise.all([
        getBalance(),
        getBankAccounts(),
        getRecentTransactions(),
        getMostExpansiveBadges(),
    ])
});

enum ModalState {
    CreateBankAccount,
    NewTransaction,
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
</script>
