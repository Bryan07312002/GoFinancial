<template>
    <div class="overflow-hidden w-screen h-screen" :class="hasAutheticated ? 'fadeOut' : ''">
        <div class="overflow-hidden h-full w-full flex items-center justify-center"
            :class="isAtCreateAccount ? 'up' : ''">
            <Login @authenticated="handleAuthenticated" @gotoSignUp="isAtCreateAccount = true"
                class=" transition duration-300 ease-out" />
        </div>
        <CreateAccount @gotoSignIn="isAtCreateAccount = false" class="transition duration-300 ease-out"
            :class="isAtCreateAccount ? 'up' : ''" />
    </div>
</template>

<script setup>
import Login from './Login.vue';
import CreateAccount from './CreateAccount.vue';
import { useRouter } from 'vue-router';
import { ref } from "vue";

const router = useRouter();

const isAtCreateAccount = ref(false);

const hasAutheticated = ref(false)

function handleAuthenticated() {
    hasAutheticated.value = true
    setTimeout(()=> router.replace({name: "Home"}), 300)
}
</script>

<style scooped>
.up {
    translate: 0 -100%;
}

@keyframes fadeOut {
    0% {
        transform: scale(1);
        opacity: 1;
    }

    50% {
        transform: scale(1.2);
        opacity: 1;
    }

    100% {
        transform: scale(0);
        opacity: 0;
    }
}

.fadeOut {
    animation: fadeOut .3s ease-in forwards;
}
</style>
