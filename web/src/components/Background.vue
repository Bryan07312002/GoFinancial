<template>
    <div class="fixed h-screen w-screen text-white dark bg-[var(--background)] z-[-1]">
        <div class="w-full h-full bg-[var(--background)] relative flex items-center justify-center">
            <div class="flex items-center max-w-[1042px] w-full h-full relative">
            </div>
            <div class="backdrop-blur-[100px] absolute w-full h-full" />
        </div>
    </div>
</template>

<style scoped>
/*<div class="ball absolute top-[-15%] left-[-15%] w-1/4 aspect-square rounded-[50%]" />*/
/*<div class="bg-gray-50 absolute bottom-[-15%] right-[-10%] w-1/4 aspect-square rounded-[50%]" />*/
.ball {
    background: var(--primary);
}
</style>
