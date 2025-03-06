<template>
    <div class="p-6 min-h-screen flex flex-col gap-2 w-72 shadow-black shadow-lg text-[var(--neutral-100)]">
        <div class="text-xl font-bold flex gap-4">
            <wallet-icon class="stroke-primary" />
            <div>Financial</div>
        </div>

        <div class="h-[1px] mb-1 mt-4 w-full bg-[var(--neutral-600)]" />

        <div @click="routerInstance.replace({ name: 'Home' }) "class="p-3 rounded-lg flex gap-4 cursor-pointer transition hover:bg-[var(--neutral-600)]"
            :class="selectedTags('Home')">
            <dashboard-icon /> Dashboard
        </div>
        <div @click="routerInstance.replace({ name: 'Transaction' })"
            class="p-3 rounded-lg flex gap-4 cursor-pointer transition hover:bg-[var(--neutral-600)]"
            :class="selectedTags('Transaction')">
            <transfer-icon />
            Transaction
        </div>
        <div class="p-3 rounded-lg flex gap-4 cursor-pointer transition hover:bg-[var(--neutral-600)]"
            :class="selectedTags('Bank Accounts')">
            <bank-icon />
            Bank Accounts
        </div>
        <div class="p-3 rounded-lg flex gap-4 cursor-pointer transition hover:bg-[var(--neutral-600)]"
            :class="selectedTags('Profile')">
            <profile-icon />
            Profile
        </div>
        <div class="p-3 rounded-lg flex gap-4 cursor-pointer transition hover:bg-[var(--neutral-600)]"
            :class="selectedTags('Settings')">
            <gear-icon />
            Settings
        </div>
    </div>
</template>

<script setup lang="ts">
import WalletIcon from '../assets/WalletIcon.vue';
import DashboardIcon from '../assets/DashboardIcon.vue';
import TransferIcon from '../assets/TransferIcon.vue';
import BankIcon from '../assets/BankIcon.vue';
import ProfileIcon from '../assets/ProfileIcon.vue';
import GearIcon from '../assets/GearIcon.vue';
import { useRoute } from 'vue-router';
import routerInstance from '../router';

const router = useRoute();

function selectedTags(name: string) {
    return router.name == name ? 'selected' : '';
}

</script>
<style scoped>
.selected {
    background: var(--neutral-700) !important;
}
</style>
