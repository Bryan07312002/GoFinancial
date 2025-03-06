<template>
    <div>
        <div class="text-[var(--neutral-400)]">{{ name }}:</div>
        <loading-icon v-if="isLoading" />
        <div v-else>{{ value }}</div>
    </div>
</template>

<script setup lang="ts">
import LoadingIcon from '../assets/LoadingIcon.vue';

defineProps<{
    name: string;
    value?: string | number;
    isLoading: boolean;
}>();
</script>
