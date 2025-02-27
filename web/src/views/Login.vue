<template>
    <div class="text-[var(--foreground)] flex justify-center items-center dark">
        <div class="p-4">
            <div class="flex flex-col gap-5">
                <h1 class="text-center text-2xl font-bold">Sign In.</h1>

                <error>{{ error }}</error>
                <Input v-model:value="loginForm.email" placeholder="Email" :disabled="isLoading" />
                <Input v-model:value="loginForm.password" class="w-full" placeholder="Password" :disabled="isLoading" />

                <Button @click="handleSignIn" :is-loading="isLoading" text="Sign In" />
            </div>
            <div class="mt-2">
                Don´t have an account?
                <b @click="emits('gotoSignUp')" class="ml-2  underline cursor-pointer">
                    Sign Up
                </b>.
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import Input from "../components/Input.vue"
import Button from "../components/Button.vue"
import Error from "../components/Error.vue"
import { ref, type Ref } from "vue";
import { useAuthStore } from "../stores/authentication";
import { authService } from "../services/auth/auth";

const emits = defineEmits(["gotoSignUp", "authenticated"])

const authStore = useAuthStore();

const isLoading = ref(false);
// @ts-ignore
const error: Ref<string> = ref(null);
const loginForm = ref({
    email: "",
    password: "",
})

async function handleSignIn(): Promise<void> {
    isLoading.value = true;

    try {
        const response = await authService.signIn(loginForm.value);
        authStore.setToken(response.token)
        emits("authenticated")
    } catch (e) {
        if (e instanceof Request) {
            error.value = await e.json();
        }
    }

    isLoading.value = false;
}
</script>
