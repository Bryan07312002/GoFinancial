<template>
    <div class="border rounded-lg text-sm px-1 py-1 min-w-10 max-w-15 text-center text-ellipsis overflow-hidden text-nowrap"
        :class="getColor()">
        {{ badge.name }}
    </div>
</template>

<script setup lang="ts">
import { type Badge } from '../services/transactions/transaction';

const props = defineProps<{ badge: Badge }>();

// function needed so a default value is passed in case color is not valid
function getColor(): string {
    switch (props.badge.color) {
        case 'red':
            return 'red'
        case 'green':
            return 'green'
        case 'purple':
            return 'purple'
        case 'yellow':
            return 'yellow'
        case 'light-blue':
            return 'light-blue'

        default:
            return 'red'
    }
}
</script>

<style scooped>
.red {
    background: hsl(360 100% 71%);
    color: hsl(360 100% 20%);
}

.yellow {
    background: hsl(53 100% 71%);
    color: hsl(53 100% 20%);
}

.light-blue {
    background: hsl(193 100% 71%);
    color: hsl(193 100% 20%);
}

.blue {
    background: hsl(243 100% 71%);
    color: hsl(243 100% 20%);
}

.green {
    background: hsl(105 100% 71%);
    color: hsl(105 100% 20%);
}

.purple {
    background: hsl(278 100% 71%);
    color: hsl(278 100% 20%);
}
</style>
