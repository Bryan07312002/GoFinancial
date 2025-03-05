<template>
    <div>
        <div class="text-[var(--muted-foreground)]">{{ name }}:</div>
        <loading-icon v-if="isLoading" />
        <div v-else>{{ value }}</div>
    </div>
</template>

<script setup lang="ts">
import type LoadingIcon from '../assets/LoadingIcon.vue';

defineProps<{
    name: string;
    value?: string | number;
    isLoading: boolean;
}>();
</script>
