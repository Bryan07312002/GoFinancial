<template>
    <div class="p-6 min-h-screen flex flex-col gap-2 w-72 shadow-black shadow-lg text-[var(--neutral-100)]">
        <div class="text-xl font-bold flex gap-4">
            <wallet-icon class="stroke-primary" />
            <div>Financial</div>
        </div>

        <div class="h-[1px] mb-1 mt-4 w-full bg-[var(--neutral-600)]" />

        <div class="p-3 rounded-lg flex gap-4">
            <dashboard-icon /> Dashboard
        </div>
        <div class="p-3 rounded-lg flex gap-4 bg-[var(--neutral-700)] selected">
            <transfer-icon />
            Transaction
        </div>
        <div class="p-3 rounded-lg flex gap-4">
            <bank-icon />
            Bank Accounts
        </div>
        <div class="p-3 rounded-lg flex gap-4">
            <profile-icon />
            Profile
        </div>
        <div class="p-3 rounded-lg flex gap-4">
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

</script>
