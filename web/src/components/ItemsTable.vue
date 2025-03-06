<template>
    <table class="text-nowrap text-ellipsis text-left">
        <thead>
            <tr>
                <th class="pr-3">Name</th>
                <th class="pr-3">Quantity</th>
                <th class="pr-3">Badges</th>
                <th class="pr-3">Value</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="item, i in items" :class="i % 2 == 0 ? 'bg-[var(--neutral-600)]' : ''">
                <td class="pr-3">{{ item.name }}</td>
                <td class="pr-3">x{{ item.quantity }}</td>
                <td class="pr-3 flex gap-1">
                    <badge v-for="badge in item.badges" :badge="badge" />
                </td>
                <td class="pr-3">R$ {{ item.value }}</td>
            </tr>
        </tbody>
    </table>
</template>

<script setup lang="ts">
import type { ItemWithBadges } from '../services/items';
import Badge from './Badge.vue';

const props = defineProps<{ items: ItemWithBadges[] }>();
console.log(props.items)
</script>
