<template>
    <div class="flex flex-col justify-between">
        <div class="text-[var(--neutral-400)] mb-3">
            Most expansive badges
        </div>
        <div class="h-full w-full text-[var(--neutral-100)] text-3xl">
            R$ {{ props.badges.map(badge => badge.value).reduce((acc, v) => parseFloat(v as any) + acc, 0) }}
        </div>

        <MostExpansiveChart v-if="badges.length > 0" :badges="badges" />
    </div>
</template>

<script setup lang="ts">
import type { BadgeWithValue } from '../services/badges/badges';
import MostExpansiveChart from './mostExpansiveChart.vue';

const props = defineProps<{ badges: BadgeWithValue[] }>();
</script>
