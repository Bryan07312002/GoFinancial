<template>
    <div class="flex items-center gap-2">
        <div @click="onClick" :class="disabled ? 'brightness-50' : ''"
            class="w-6 h-6 border overflow-hidden border-[var(--border)] rounded-[var(--radius)] aspect-square cursor-pointer">
            <div class="w-full h-full bg-[var(--primary)] flex items-center justify-center" v-if="value">
                <svg class="w-6" xmlns="http://www.w3.org/2000/svg" width="24" height="24" viewBox="0 0 24 24"
                    fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round">
                    <path d="M20 6 9 17l-5-5" />
                </svg>
            </div>
        </div>
        <label v-if="label" class="select-none">
            {{ label }}
        </label>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    value: boolean,
    disabled?: boolean,
    label?: string,
}>()
const emits = defineEmits(['update:value'])

function onClick() {
    if (props.disabled) return

    emits('update:value', !props.value)
}
</script>
