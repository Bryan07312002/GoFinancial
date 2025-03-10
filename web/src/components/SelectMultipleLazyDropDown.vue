<template>
    <div class="flex flex-col relative">
        <label class="pb-1" v-if="label">{{ label }}</label>

        <!-- Dropdown trigger -->
        <div class="relative" :class="disabled ? 'brightness-50' : ''">
            <button type="button" @click="toggleDropdown" :disabled="disabled"
                class="w-full flex justify-between items-center border border-[var(--secondary-2)] rounded-md p-2 text-left focus:border-[var(--primary)] focus:outline-none disabled:opacity-50">
                <span v-if="selectedOption">{{ selectedOption.name }}</span>
                <span v-else class="text-gray-400">{{ placeholder }}</span>
                <span class="transform transition-transform text-[var(--secondary-2)]"
                    :class="{ 'rotate-180': isOpen }">▼</span>
            </button>

            <!-- Dropdown content -->
            <div v-show="isOpen"
                class="absolute w-fit bg-[var(--secondary-1)] z-10 mt-1 border border-[var(--secondary-2)] rounded-md shadow-lg max-h-60 overflow-auto">
                <!-- Search input -->
                <input v-model="searchQuery" @input="handleSearch" placeholder="Search..."
                    class="p-2 border-b border-[var(--secondary-2)] w-full focus:outline-none" />

                <!-- Loading state -->
                <div v-if="loading" class="p-2 text-gray-500">Loading...</div>

                <!-- Options -->
                <ul v-else>
                    <li v-for="option in filteredOptions" :key="option.value" @click="selectOption(option.value)"
                        class="p-2 flex gap-4 hover:bg-[var(--primary)] cursor-pointer"
                        :class="{ 'bg-[var(--primary-hover)]': option.value === value }">
                        <check-box :value="!!value.find((val) => isEqual(val, option.value))" />{{ option.name }}
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';
import CheckBox from './CheckBox.vue';

interface Option {
    name: string;
    value: any;
}

const props = defineProps<{
    value: any[];
    options: Option[];
    label?: string;
    placeholder?: string;
    disabled?: boolean;
    loading?: boolean;
}>();

const emit = defineEmits(['update:value', 'search']);

const isOpen = ref(false);
const searchQuery = ref('');

function isEqual(obj1: unknown, obj2: unknown): boolean {
    return JSON.stringify(obj1) === JSON.stringify(obj2);
}

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
        opt.name.toLowerCase().includes(searchQuery.value.toLowerCase())
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
const selectOption = (value: any) => {
    const index = props.value.findIndex(el => el == value);
    if (index != -1) {
        delete props.value[index];
        console.log("aaa")
        emit('update:value', [
            ...props.value,
        ]);

        return
    }

    emit('update:value', [
        ...props.value,
        value,
    ]);
};


// Close dropdown when value changes
watch(() => props.value, closeDropdown);
</script>
