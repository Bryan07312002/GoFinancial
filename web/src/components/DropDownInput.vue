<template>
    <div class="flex flex-col">
        <label class="pb-1" v-if="label">{{ label }}</label>
        <select :value="value" @change="$emit('update:value', ($event.target as HTMLSelectElement).value)"
            :disabled="disabled"
            class="border border-[var(--border)] rounded-[var(--radius)] focus:border-[var(--primary)] focus:outline-none p-2 appearance-none bg-[url('data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHdpZHRoPSIyNCIgaGVpZ2h0PSIyNCIgdmlld0JveD0iMCAwIDI0IDI0IiBmaWxsPSJub25lIiBzdHJva2U9ImN1cnJlbnRDb2xvciIgc3Ryb2tlLXdpZHRoPSIyIiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiPjxwb2x5bGluZSBwb2ludHM9IjYgOSAxMiAxNSAxOCA5Ii8+PC9zdmc+')] bg-no-repeat bg-[right_0.5rem_center]">
            <option v-if="placeholder" value="" disabled selected hidden>{{ placeholder }}</option>
            <option v-for="option in normalizedOptions" :key="option.value" :value="option.value"
                :disabled="option.disabled">
                {{ option.label }}
            </option>
        </select>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue'

const props = defineProps<{
    value: string | number,
    options: Array<string | { value: string | number, label: string, disabled?: boolean }>
    label?: string,
    placeholder?: string,
    disabled?: boolean,
}>()

defineEmits(['update:value'])

const normalizedOptions = computed(() => {
    return props.options.map(option => {
        if (typeof option === 'string') {
            return { value: option, label: option }
        }
        return option
    })
})
</script>
