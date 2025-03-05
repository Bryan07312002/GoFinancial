<template>
    <div class="p-6 min-w-72 h-screen flex flex-col gap-2 w-72 shadow-black shadow-lg text-[var(--neutral-100)]">
        <div class="text-xl font-bold flex gap-4">
            <wallet-icon class="stroke-primary" />
            <div>Financial</div>
        </div>

        <div class="h-[1px] mb-1 mt-4 w-full bg-[var(--neutral-600)]" />

        <div class="p-3 rounded-lg"> Dashboard </div>
        <div class="p-3 rounded-lg bg-[var(--neutral-700)] selected"> Transaction </div>
        <div class="p-3 rounded-lg"> Bank Accounts </div>
        <div class="p-3 rounded-lg "> Profile </div>
    </div>
</template>

<script setup lang="ts">
import WalletIcon from '../assets/WalletIcon.vue';
</script>
