<template>
    <div class="flex flex-col justify-between">
        <div class="h-full w-full text-3xl">
            R$ {{ }}
        </div>

        <MostExpansiveChart v-if="badges.length > 0" :badges="badges" />
    </div>
</template>

<script setup lang="ts">
import type { BadgeWithValue } from '../services/badges/badges';
import MostExpansiveChart from './mostExpansiveChart.vue';

const props = defineProps<{ badges: BadgeWithValue[] }>();
//console.log(props.badges)
//const max = props.badges.map(badge => badge.value).reduce((acc, v) => v + acc);
</script>
