<template>
    <div class="flex max-h-screen max-w-screen">
        <sidebar />
        <card class="max-w-[1440px] overflow-x-auto mx-auto m-6 flex flex-col gap-6 w-full p-6">

        </card>
    </div>
</template>

<script setup lang="ts">
import Sidebar from '../components/Sidebar.vue';
import Card from '../components/Card.vue';
</script>
