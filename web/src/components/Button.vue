<template>
    <button class="py-2 px-3 rounded-lg hover:brightness-50 font-md cursor-pointer" :class="buttonColorType()">
        <div v-if="!isLoading">
            {{ text }}
        </div>
        <div v-else>
            <loading-icon />
        </div>
    </button>
</template>

<style scoped>
.button-primary {
    background: var(--primary);
}

.button-secondary {
    background: var(--secondary);
}
</style>

<script setup lang="ts">
import LoadingIcon from '../assets/LoadingIcon.vue';

type ColorTypes = "primary" | "secondary"

const props = defineProps<{
    text: string,
    isLoading?: boolean,
    bottomColorType?: ColorTypes,
}>()

function buttonColorType(): string {
    switch (props.bottomColorType) {
        case "primary":
            return "button-primary"

        case "secondary":
            return "button-secondary"
    }

    return "button-primary"
}
</script>
