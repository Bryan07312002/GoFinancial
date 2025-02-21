<template>
    <div class="flex items-center gap-2">
        <div class="w-6 h-6 border border-[var(--border)] rounded-[var(--radius)] aspect-square cursor-pointer">
            <div class="w-full h-ful bg-[var(--primary)]" v-if="!value">
.
            </div>
        </div>
        <label v-if="label" class="select-none">
            {{ label }}
        </label>
    </div>
</template>

<script setup lang="ts">
defineProps<{ label: string, value: boolean }>()
const emits = defineEmits(['update:value'])
</script>
