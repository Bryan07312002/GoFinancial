<template>
    <card class="min-w-80 p-4 flex flex-col gap-4">
        <div class="flex justify-between items-center">
            <div class="text-2xl">Bank Accounts</div>
            <Button @click= "emits('createBankAccount')" text="New Bank Account" />
        </div>
        <div v-if="isLoading" class="flex items-center justify-center h-full">
            <loading-icon />
        </div>
        <div v-else class="h-full flex flex-col gap-4">
            <div @click="emits('openBankAccount', bankAccount)" v-for="bankAccount in bankAccounts"
                class="flex justify-between p-4 hover:bg-[var(--accent)] rounded-[var(--radius)]">
                <div class="flex gap-4">
                    <bank-icon class="w-6 h-6 stroke-primary" />
                    <div> {{ bankAccount.name }} </div>
                </div>
                <div> R$ {{ bankAccount.credit }} </div>
            </div>
        </div>
    </card>
</template>

<script setup lang="ts">
import Card from "./Card.vue";
import Button from "./Button.vue";
import BankIcon from "../assets/BankIcon.vue";
import LoadingIcon from '../assets/LoadingIcon.vue';
import { type BankAccount } from "../services/bankAccounts/bankAccounts";

defineProps<{ bankAccounts: BankAccount[], isLoading?: boolean }>();
const emits = defineEmits(['openBankAccount', 'createBankAccount']);
</script>

<style scooped></style>
