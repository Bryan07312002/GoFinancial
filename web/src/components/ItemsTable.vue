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
            <tr v-for="item, i in items" :class="lineColor(i)">
                <td class="pr-3">{{ item.name }}</td>
                <td class="pr-3">x{{ item.quantity }}</td>
                <td class="pr-3 flex gap-1">
                    <badge-vue v-for="badge in item.badges" :badge="badge" />
                </td>
                <td class="pr-3">R$ {{ item.value }}</td>
                <td v-if="editMode">
                    <div @click="deleteItem(item.id)"
                        class="my-1 hover:brightness-50 cursor-pointer rounded-lg w-5 h-5 text-center bg-red-300">
                        -
                    </div>
                </td>
            </tr>

            <tr v-for="item, i in addedItems" :class="i % 2 == 0 ? 'bg-[var(--neutral-600)]' : ''">
                <td class="pr-3">* {{ item.name }}</td>
                <td class="pr-3">x{{ item.quantity }}</td>
                <td class="pr-3 flex gap-1">
                    <badge-vue v-for="badge in item.badges" :badge="badge" />
                </td>
                <td class="pr-3">R$ {{ item.value }}</td>
                <td>
                    <div @click="deleteAddedItem(i)"
                        class="my-1 hover:brightness-50 cursor-pointer rounded-lg w-5 h-5 text-center bg-red-300">
                        -
                    </div>
                </td>
            </tr>

            <tr v-if="newItem">
                <td class="pr-3"><Input v-model:value="newItem.name" /></td>
                <td class="pr-3"><Input v-model:value="newItem.quantity" /></td>
                <td class="pr-3">
                    <select-multiple-lazy-drop-down v-model:value="newItem.badges" v-if="badgeOptions"
                        :options="badgeOptions" />
                </td>
                <td class="pr-3"><monetary-input v-model:value="newItem.value" /></td>
                <td>
                    <div class="p-2 bg-red-300">-</div>
                </td>
            </tr>
        </tbody>
        <tbody>
        </tbody>
    </table>
</template>

<script setup lang="ts">
import Input from './Input.vue';
import BadgeVue from './Badge.vue';
import MonetaryInput from './MonetaryInput.vue';
import type { ItemWithBadges } from '../services/items';
import type { Badge } from '../services/transactions/transaction';
import SelectMultipleLazyDropDown from './SelectMultipleLazyDropDown.vue';

const props = defineProps<{
    items: ItemWithBadges[],
    editMode?: boolean
    addedItems?: Omit<ItemWithBadges, "id" | "transaction_id">[],
    newItem?: Omit<ItemWithBadges, "id" | "transaction_id"> | null,
    badgeOptions?: { name: string, value: Badge }[],
    deleteItems?: number[],
}>();

const emits = defineEmits(["update:deleteItems", "update:addedItem"])

function lineColor(i: number) {
    let classes = "";

    if (i % 2 == 0)
        classes = classes + 'bg-[var(--neutral-600)]'

    if (!props.items.filter((item) => props.deleteItems?.includes(item.id)))
        classes = classes + 'brightness-50'

    return classes
}


function deleteAddedItem(index: number) {
    if (!props.addedItems) return

    props.addedItems.splice(index, 1);

    emits("update:addedItem", [
        ...props.addedItems,
    ]);

    return
}

function deleteItem(id: number) {
    if (!props.deleteItems) return

    const index = props.deleteItems?.findIndex(el => el == id);
    if (index != -1) {
        props.deleteItems?.splice(index, 1);

        emits("update:deleteItems", [
            ...props.deleteItems,
        ]);

        return
    }

    emits("update:deleteItems", [
        ...props.deleteItems,
        id,
    ]);
}
</script>
