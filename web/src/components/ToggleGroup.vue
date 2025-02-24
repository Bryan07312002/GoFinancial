<template>
    <div>
        <div class="text-[var(--muted-foreground)] mb-2">{{ label }}:</div>
        <div class="flex gap-4" :class="disabled ? 'brightness-50' : ''">
            <div v-for="option in options" @click="onClick(option.value)" :class="value === option.value ?
                'bg-[var(--primary)] color--[var(--primary-foreground)]' :
                'bg-[var(--muted)] color--[var(--muted-foreground)]'"
                class="flex-1 cursor-pointer text-center rounded-[var(--radius)] py-2 px-3">{{
                    option.name }}</div>
        </div>
    </div>
</template>

<script setup lang="ts">
// TODO: add icons
const props = defineProps<{
    label: string,
    options: { name: string, value: number | string }[],
    value: number | string,
    disabled?: boolean,
}>();

const emits = defineEmits(["update:value"]);

function onClick(value: number | string) {
    if (props.disabled) return;

    console.log(value)
    emits('update:value', value)
}
</script>
