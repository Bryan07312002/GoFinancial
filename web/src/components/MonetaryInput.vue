<template>
    <div class="flex flex-col">
        <label class="pb-1" v-if="label">{{ label }}</label>
        <input :value="displayValue" @input="onInput" @focus="onFocus" @blur="onBlur" :disabled="disabled"
            :placeholder="placeholder" :class="disabled ? 'brightness-50' : ''"
            class="border border-[var(--secondary-2)] rounded-md focus:border-[var(--primary)] focus:outline-none p-2">
    </div>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue'

const props = defineProps<{
    value: number,
    label?: string,
    placeholder?: string,
    disabled?: boolean,
    currencySymbol?: string,
    decimalPlaces?: number,
}>()

const emit = defineEmits(['update:value'])

const currency = computed(() => props.currencySymbol ?? '$')
const decimals = computed(() => props.decimalPlaces ?? 2)
const isFocused = ref(false)
const displayValue = ref('')

// Initialize display value
displayValue.value = formatCurrency(props.value)

// Watch for external value changes
watch(() => props.value, (newVal) => {
    if (!isFocused.value) {
        displayValue.value = formatCurrency(newVal)
    }
})

function formatCurrency(value: number): string {
    return `${currency.value}${value.toFixed(decimals.value)}`
}

function onFocus() {
    isFocused.value = true
    displayValue.value = props.value.toFixed(decimals.value)
}

function onBlur() {
    isFocused.value = false
    let numericValue = parseFloat(displayValue.value) || 0
    numericValue = parseFloat(numericValue.toFixed(decimals.value))
    emit('update:value', numericValue)
    displayValue.value = formatCurrency(numericValue)
}

function onInput(event: Event) {
    const target = event.target as HTMLInputElement
    let inputVal = target.value

    // Sanitize input
    let sanitized = inputVal
        .replace(/[^0-9.]/g, '') // Remove non-numeric chars
        .replace(/(\..*)\./g, '$1') // Allow only one decimal point

    // Handle leading decimal
    if (sanitized.startsWith('.')) {
        sanitized = '0' + sanitized
    }

    // Split into whole and decimal parts
    const [whole, decimal] = sanitized.split('.')

    // Truncate decimal places
    if (decimal) {
        sanitized = `${whole}.${decimal.slice(0, decimals.value)}`
    }

    // Handle empty value
    if (sanitized === '') {
        sanitized = '0'
    }

    // Update display and emit value
    displayValue.value = sanitized
    const numericValue = parseFloat(sanitized) || 0
    emit('update:value', numericValue)
}
</script>
