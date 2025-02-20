<template>
    <button class="p-2 rounded-[var(--radius)] hover:brightness-50 cursor-pointer">
        <div v-if="!isLoading">
            {{ text }}
        </div>
        <div v-else>
            <loading-icon />
        </div>
    </button>
</template>

<style scoped>
button {
    background: var(--primary-gradient);
}
</style>

<script setup lang="ts">
import LoadingIcon from '../assets/LoadingIcon.vue';
defineProps<{ text: string, isLoading?: boolean }>()
</script>
