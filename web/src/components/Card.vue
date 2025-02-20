<template>
    <div class="card-bg">
        <slot />
    </div>
</template>

<style>
.dark .card-bg {
    background: linear-gradient(45deg, rgba(80, 80, 80, .2), rgba(100, 100, 100, .2), rgba(80, 80, 80, .2));
    border-radius: 16px;
    box-shadow: 0 4px 30px rgba(0, 0, 0, 0.1);
    backdrop-filter: blur(7.4px);
    -webkit-backdrop-filter: blur(7.4px);
    border: 1px solid rgba(95, 95, 95, 0.23);
}
</style>
