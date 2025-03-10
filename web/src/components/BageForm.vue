<template>
    <div class="flex">
        <Input v-model:value="badge.name" />
        <drop-down-input v-model:value="badge.name" :options="colors" />
    </div>
</template>

<script setup lang="ts">
import type { Badge } from '../services/transactions/transaction';
import Input from './Input.vue';
import DropDownInput from './DropDownInput.vue';
import { ref } from 'vue';

defineProps<{
    badge: Omit<Badge, "id">
}>();

const colors = ref([
    { name: "red", value: "red" },
    { name: "green", value: "green" },
    { name: "yellow", value: "yellow" },
])
</script>
