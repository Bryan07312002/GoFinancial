<template>
    <div class="flex flex-col relative" v-click-outside="closeDropdown">
        <label class="pb-1" v-if="label">{{ label }}</label>

        <!-- Dropdown trigger -->
        <div class="relative">
            <button type="button" @click="toggleDropdown" :disabled="disabled"
                class="w-full flex justify-between items-center border border-[var(--border)] rounded-[var(--radius)] p-2 text-left focus:border-[var(--primary)] focus:outline-none disabled:opacity-50">
                <span v-if="selectedOption">{{ selectedOption.label }}</span>
                <span v-else class="text-gray-400">{{ placeholder }}</span>
                <span class="transform transition-transform" :class="{ 'rotate-180': isOpen }">▼</span>
            </button>

            <!-- Dropdown content -->
            <div v-show="isOpen"
                class="absolute bg-[var(--background)] z-10 w-full mt-1 border border-[var(--border)] rounded-[var(--radius)] shadow-lg max-h-60 overflow-auto">
                <!-- Search input -->
                <input v-model="searchQuery" @input="handleSearch" placeholder="Search..."
                    class="p-2 border-b border-[var(--border)] w-full focus:outline-none" />

                <!-- Loading state -->
                <div v-if="loading" class="p-2 text-gray-500">Loading...</div>

                <!-- Options -->
                <ul v-else>
                    <li v-for="option in filteredOptions" :key="option.value" @click="selectOption(option)"
                        class="p-2 hover:bg-[var(--primary)] cursor-pointer"
                        :class="{ 'bg-[var(--primary-hover)]': option.value === value }">
                        {{ option.label }}
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';

interface Option {
    label: string;
    value: string | number;
}

const props = defineProps<{
    value: string | number;
    options: Option[];
    label?: string;
    placeholder?: string;
    disabled?: boolean;
    loading?: boolean;
}>();

const emit = defineEmits(['update:value', 'search']);

const isOpen = ref(false);
const searchQuery = ref('');

// Handle click outside
const closeDropdown = () => isOpen.value = false;
const toggleDropdown = () => isOpen.value = !isOpen.value;

// Selected option display
const selectedOption = computed(() =>
    props.options.find(opt => opt.value === props.value)
);

// Filter options locally
const filteredOptions = computed(() =>
    props.options.filter(opt =>
        opt.label.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
);

// Search debouncing
let searchTimeout = 0;
const handleSearch = () => {
    clearTimeout(searchTimeout);
    searchTimeout = setTimeout(() => {
        emit('search', searchQuery.value);
    }, 300);
};

// Option selection
const selectOption = (option: Option) => {
    emit('update:value', option.value);
    closeDropdown();
    searchQuery.value = '';
};

// Close dropdown when value changes
watch(() => props.value, closeDropdown);
</script>
