<template>
    <div class="flex flex-col relative" v-click-outside="closeDropdown">
        <label class="pb-1" v-if="label">{{ label }}</label>

        <!-- Dropdown trigger -->
        <div class="relative">
            <button type="button" @click="toggleDropdown" :disabled="disabled"
                class="w-full flex justify-between items-center border border-[var(--border)] rounded-[var(--radius)] p-2 text-left focus:border-[var(--primary)] focus:outline-none disabled:opacity-50">
                <span class="flex flex-wrap gap-1">
                    <template v-if="multiple && selectedOptions.length > 0">
                        <span v-for="option in selectedOptions" :key="option.value"
                            class="px-2 py-1 bg-[var(--primary)] rounded-[var(--radius)] text-sm">
                            {{ option.label }}
                        </span>
                    </template>
                    <span v-else-if="!multiple && selectedOption">
                        {{ selectedOption.label }}
                    </span>
                    <span v-else class="text-gray-400">
                        {{ placeholder }}
                    </span>
                </span>
                <span class="transform transition-transform" :class="{ 'rotate-180': isOpen }">▼</span>
            </button>

            <!-- Dropdown content -->
            <div v-show="isOpen"
                class="absolute z-10 w-full mt-1 bg-[var(--background)] border border-[var(--border)] rounded-[var(--radius)] shadow-lg max-h-60 overflow-auto">
                <!-- Options list with checkboxes -->
                <ul>
                    <li v-for="option in options" :key="option.value" @click="toggleOption(option)"
                        class="p-2 hover:bg-[var(--primary-hover)] cursor-pointer flex items-center gap-2"
                        :class="{ 'bg-[var(--primary-hover)]': isSelected(option) }">
                        <input type="checkbox" :checked="isSelected(option)"
                            class="w-4 h-4 text-[var(--primary)] focus:ring-[var(--primary)]"
                            :class="{ 'opacity-50': disabled }" @click.stop>
                        <span>{{ option.label }}</span>
                    </li>
                </ul>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';

interface Option {
    label: string;
    value: string | number;
}

const props = defineProps<{
    value: string | number | Array<string | number>;
    options: Option[];
    label?: string;
    placeholder?: string;
    disabled?: boolean;
    multiple?: boolean;
}>();

const emit = defineEmits(['update:modelValue']);

const isOpen = ref(false);

// Handle click outside
const closeDropdown = () => isOpen.value = false;
const toggleDropdown = () => isOpen.value = !isOpen.value;

// Handle selections
const isSelected = (option: Option) => {
    return props.multiple
        ? (props.modelValue as Array<string | number>).includes(option.value)
        : props.modelValue === option.value;
};

const toggleOption = (option: Option) => {
    if (props.multiple) {
        const newValue = [...props.modelValue as Array<string | number>];
        const index = newValue.indexOf(option.value);

        index === -1
            ? newValue.push(option.value)
            : newValue.splice(index, 1);

        emit('update:modelValue', newValue);
    } else {
        emit('update:modelValue', option.value);
        closeDropdown();
    }
};

// Display helpers
const selectedOption = computed(() =>
    props.options.find(opt => opt.value === props.modelValue)
);

const selectedOptions = computed(() =>
    props.options.filter(opt =>
        (props.modelValue as Array<string | number>).includes(opt.value)
    )
);
</script>
