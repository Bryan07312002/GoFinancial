<template>
    <div class="min-h-screen flex">
        <div class="flex flex-col min-h-screen gap-4 p-4 w-full max-w-7xl mx-auto">
            <div class="flex  w-full justify-around gap-4 flex-wrap">

                <card class="flex-1 min-w-32 h-28 p-4 flex flex-col justify-between">
                    <div class="w-full flex items-center justify-between">
                        <div>Total Balance</div>
                        <bank-icon class="w-6 h-6 stroke-primary" />
                    </div>

                    <b class="text-xl">R$ 24562,00</b>
                </card>

                <card class="flex-1 min-w-32 h-28 p-4 flex flex-col justify-between">
                    <div class="w-full flex items-center justify-between">
                        <div>Total Credit</div>

                        <wallet-icon class="w-6 h-6 stroke-primary" />
                    </div>

                    <b class="text-xl text-green-300">R$ 24562,00</b>
                </card>

                <card class="flex-1 min-w-32 h-28 p-4 flex flex-col justify-between">
                    <div class="w-full flex items-center justify-between">
                        <div>Total Debt</div>

                        <credit-card-icon class="w-6 h-6 stroke-primary" />
                    </div>

                    <b class="text-xl text-red-300">R$ 24562,00</b>
                </card>
            </div>

            <div class="flex w-full justify-around gap-4 flex-wrap flex-grow min-h-96 h-1/3">
                <bank-account-list @create-bank-account="openCreateBankAccountModal" :is-loading="isBankAccountsLoading"
                    :bank-accounts="bankAccounts" class="flex-1" />
                <recent-transfer-activities @new-transaction="openNewTransactionModal"
                    :is-loading="isTransactionsLoading" :transactions="transactions" class="flex-1" />
            </div>

            <card class="flex-1 h-1/5 flex-grow min-h-[200px]">
                200
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
import Card from "../components/Card.vue";
import Modal from "../components/Modal.vue";
import BankIcon from "../assets/BankIcon.vue"
import { ref, onMounted, type Ref } from "vue";
import WalletIcon from "../assets/WalletIcon.vue"
import CreditCardIcon from "../assets/CreditCardIcon.vue"
import BankAccountList from "../components/BankAccountList.vue";
import CreateBankAccount from "../components/CreateBankAccount.vue";
import NewTransaction from "../components/NewTransaction.vue";
import RecentTransferActivities from "../components/RecentTransferActivities.vue";
import { BankAccountService, type BankAccount } from "../services/bankAccounts/bankAccounts";
import { type Transaction, TransactionType, PaymentMethod } from "../services/transactions/transaction";

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

onMounted(async () => {
    getBankAccounts();
});

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
const transactions: Ref<Transaction[]> = ref([
    {
        id: 1,
        type: TransactionType.Income,
        method: PaymentMethod.DebitCard,
        establishment: "work",
        credit: false,
        value: 2000.00,
        date: new Date(Date.now()),
        cardId: 1,
        bankAccountId: 1,
    },
    {
        id: 1,
        type: TransactionType.Expense,
        method: PaymentMethod.DebitCard,
        establishment: "zaffari",
        credit: false,
        value: 2000.00,
        date: new Date(Date.now()),
        cardId: 1,
        bankAccountId: 1,
    },
    {
        id: 1,
        type: TransactionType.Expense,
        method: PaymentMethod.DebitCard,
        establishment: "zaffari",
        credit: false,
        value: 2000.00,
        date: new Date(Date.now()),
        cardId: 1,
        bankAccountId: 1,
    },
]);
</script>
