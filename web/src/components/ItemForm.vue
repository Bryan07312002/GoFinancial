<template>
    <div class="">
        <Input v-model:value="item.name" />
        <Input v-model:value="item.quantity" />
    </div>
</template>

<script setup lang="ts">
import { type ItemWithBadges } from '../services/items';
import Input from './Input.vue';

defineProps<{
    item: ItemWithBadges
}>();
</script>
