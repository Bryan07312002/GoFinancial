<template>
    <div class="flex flex-col">
        <label class="pb-1" v-if="label" for="">{{ label }}</label>
        <input :value="value" @input="$emit('update:value', ($event.target as any).value)"
            :disabled="disabled != undefined && disabled" :placeholder="placeholder"
            class="border border-[var(--border)] rounded-[var(--radius)] focus:border-[var(--primary)] focus:outline-none p-2">
    </div>
</template>

<script setup lang='ts'>
defineProps<{
    value: string,
    label?: string,
    placeholder?: string,
    disabled?: boolean,
}>()
const emits = defineEmits(["update:value"])
</script>
