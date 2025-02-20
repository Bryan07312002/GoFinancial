<template>
    <BackgroundVue />
    <div class="dark text-[var(--foreground)]">
        <router-view />
    </div>
</template>

<script setup lang="ts">
import BackgroundVue from "./components/Background.vue"

</script>
