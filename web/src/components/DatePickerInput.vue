<template>
    <div class="relative w-full">
        <!-- Date picker -->
        <div v-if="isOpen"
            class="absolute bg-[var(--card)] left-0 bottom-full mb-2 border border-[var(--border)] rounded-[var(--radius)] p-4 z-10">
            <!-- Header -->
            <div class="flex items-center justify-between mb-4">
                <button @click="previousMonth"
                    class="p-2 cursor-pointer hover:bg-[var(--muted)] hover:color-[var(--muted-foreground)] rounded-[var(--radius)]">
                    ←
                </button>
                <div class="font-semibold">
                    {{ currentMonth }} {{ currentYear }}
                </div>
                <button @click="nextMonth"
                    class="p-2 cursor-pointer hover:bg-[var(--muted)] hover:color-[var(--muted-foreground)] rounded-[var(--radius)]">
                    →
                </button>
            </div>

            <!-- Days grid -->
            <div class="grid grid-cols-7 gap-1 text-center">
                <!-- Week days -->
                <div v-for="day in ['Sun', 'Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat']" :key="day"
                    class="text-xs font-medium text-[var(--muted-foreground)] py-1">
                    {{ day }}
                </div>

                <!-- Calendar days -->
                <button v-for="(day, index) in calendarDays" :key="index" @click="selectDate(day.date)"
                    class="w-7 h-7 cursor-pointer text-sm rounded-[var(--radius)] hover:bg-[var(--muted)] hover:color-[var(--muted-foreground)]"
                    :class="{
                        'text-[var(--muted)]': !day.isCurrentMonth,
                        'bg-[var(--primary)] text-white hover:bg-[var(--primary)]': day.isSelected,
                        'font-semibold': day.isToday,
                    }">
                    {{ day.label }}
                </button>
            </div>
        </div>

        <!-- Input field -->
        <Input class="w-full" :class="disabled ? 'brightness-50':''" :disabled="disabled" type="text" readonly :value="value ? formatDate(value) : 'Select date'"
            @click="onClick" />
    </div>
</template>

<script setup lang="ts">
import Input from "./Input.vue";
import { ref, computed, type Ref } from 'vue';

const props = defineProps<{value: Date, disabled?: boolean}>();
const emits = defineEmits(['update:value'])

const isOpen = ref(false);
const currentDate: Ref<Date> = ref(props.value);

// Computed month/year display
const currentMonth = computed(() => {
    return currentDate.value.toLocaleString('default', { month: 'long' });
});

const currentYear = computed(() => {
    return currentDate.value.getFullYear();
});

// Computed calendar days
const calendarDays = computed(() => {
    const year = currentDate.value.getFullYear();
    const month = currentDate.value.getMonth();

    const firstDay = new Date(year, month, 1);
    const lastDay = new Date(year, month + 1, 0);
    const daysInMonth = lastDay.getDate();

    const startDay = firstDay.getDay();
    const endDay = lastDay.getDay();

    const days = [];

    // Previous month days
    for (let i = startDay - 1; i >= 0; i--) {
        const date = new Date(year, month, -i);
        days.push({
            label: date.getDate(),
            date: date,
            isCurrentMonth: false,
            isToday: false,
            isSelected: false,
        });
    }

    // Current month days
    for (let day = 1; day <= daysInMonth; day++) {
        const date = new Date(year, month, day);
        const isToday = date.toDateString() === new Date().toDateString();
        const isSelected = props.value
            ? date.toDateString() === props.value.toDateString()
            : false;

        days.push({
            label: day,
            date: date,
            isCurrentMonth: true,
            isToday: isToday,
            isSelected: isSelected,
        });
    }

    // Next month days
    for (let i = 1; i <= (6 - endDay); i++) {
        const date = new Date(year, month + 1, i);
        days.push({
            label: date.getDate(),
            date: date,
            isCurrentMonth: false,
            isToday: false,
            isSelected: false,
        });
    }

    return days;
});

// Navigation
function previousMonth() {
    currentDate.value = new Date(
        currentDate.value.getFullYear(),
        currentDate.value.getMonth() - 1,
        1
    );
}

function nextMonth() {
    currentDate.value = new Date(
        currentDate.value.getFullYear(),
        currentDate.value.getMonth() + 1,
        1
    );
}

// Date selection
function selectDate(date: Date) {
    emits('update:value', date);
    isOpen.value = false;
}

// Date formatting
function formatDate(date: Date) {
    return date.toLocaleDateString('en-US', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
    });
}

function onClick(){
    isOpen.value = !isOpen.value
}
</script>
