<template>
    <div>
        <h2 class="text-2xl pb-4">Most Expends Classes</h2>

        <div v-if="isLoading">
            <loading-icon />
        </div>

        <div v-else class="flex flex-wrap gap-4">
            <div class=" p-4 border border-[var(--border)] rounded-[var(--radius)] flex gap-4"
                v-for="badge in Array(5).fill(badges).flat()">
                <wallet-icon class="w-6 h-6  stroke-[var(--primary)]" />
                <div>{{ badge.name }}</div>
                <div>R$ {{ badge.value }}</div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { type BadgeWithValue } from '../services/badges/badges';
import LoadingIcon from '../assets/LoadingIcon.vue';
import WalletIcon from '../assets/WalletIcon.vue';

defineProps<{ badges?: BadgeWithValue[], isLoading?: boolean }>();
</script>
