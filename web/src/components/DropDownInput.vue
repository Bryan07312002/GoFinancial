<template>
    <div class="flex flex-col relative">
        <label class="pb-1" v-if="label">{{ label }}</label>

        <!-- Dropdown trigger -->
        <div class="relative" :class="disabled ? 'brightness-50' : ''">
            <button type="button" @click="toggleDropdown" :disabled="disabled"
                class="w-full flex justify-between items-center border border-[var(--secondary-2)] rounded-md p-2 text-left focus:border-[var(--primary)] focus:outline-none disabled:opacity-50">
                <span v-if="selectedOption">{{ selectedOption.name }}</span>
                <span v-else class="text-gray-400">{{ placeholder }}</span>
                <span class="transform transition-transform text-[var(--secondary-2)]" :class="{ 'rotate-180': isOpen }">▼</span>
            </button>

            <!-- Dropdown content -->
            <div v-show="isOpen"
                class="absolute bg-[var(--secondary-1)] z-10 min-w-fit w-full mt-1 border border-[var(--secondary-2)] rounded-md shadow-lg max-h-60 overflow-auto">
                <!-- Search input -->
                <input v-if="showSearch" v-model="searchQuery" @input="handleSearch" placeholder="Search..."
                    class="p-2 border-b border-[var(--secondary-2)] w-full focus:outline-none" />

                <!-- Loading state -->
                <div v-if="loading" class="p-2 text-gray-500">Loading...</div>

                <!-- Options -->
                <ul v-else>
                    <li v-for="option in filteredOptions" :key="option.value" @click="selectOption(option)"
                        class="p-2 hover:bg-[var(--primary)] cursor-pointer"
                        :class="{ 'bg-[var(--primary-hover)]': option.value === value }">
                        {{ option.name }}
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, watch } from 'vue';

interface Option {
    name: string;
    value: string | number;
}

const props = defineProps<{
    value: string | number;
    options: Option[];
    label?: string;
    placeholder?: string;
    disabled?: boolean;
    loading?: boolean;
    showSearch?:boolean;
}>();

const emit = defineEmits(['update:value']);

const isOpen = ref(false);
const searchQuery = ref('');

const closeDropdown = () => isOpen.value = false;
const toggleDropdown = () => isOpen.value = !isOpen.value;

const selectedOption = computed(() =>
    props.options.find(opt => opt.value === props.value)
);

// Filter options locally
const filteredOptions = computed(() =>
    props.options.filter(opt =>
        opt.name.toLowerCase().includes(searchQuery.value.toLowerCase())
    )
);

const selectOption = (option: Option) => {
    emit('update:value', option.value);
    closeDropdown();
    searchQuery.value = '';
};

watch(() => props.value, closeDropdown);
</script>
