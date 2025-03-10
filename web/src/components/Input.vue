<template>
    <div class="flex flex-col">
        <label class="pb-1" v-if="label" for="">{{ label }}</label>
        <error v-if="error"> {{ error }} </error>
        <input :value="value" @input="$emit('update:value', ($event.target as any).value)"
            :disabled="disabled != undefined && disabled" :placeholder="placeholder" :readonly="readonly"
            class="border border-[var(--secondary-4)] bg-[var(--secondary-1)] rounded-md focus:border-[var(--primary)] focus:outline-none p-2">
    </div>
</template>

<script setup lang='ts'>
import Error from "./Error.vue";

defineProps<{
    value: string | number,
    label?: string,
    placeholder?: string,
    readonly?: boolean,
    disabled?: boolean,
    error?: string
}>()

const emits = defineEmits(["update:value"])
</script>
