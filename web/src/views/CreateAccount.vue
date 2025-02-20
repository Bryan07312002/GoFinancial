<template>
    <div class="w-screen h-screen text-[var(--foreground)] flex justify-center items-center dark">
        <div class="p-4">
            <div class="flex flex-col gap-5">
                <h1 class="text-center text-2xl font-bold">Sign Up.</h1>

                <Input v-model:value="createAccountForm.email" placeholder="Email" :disabled="isLoading" />
                <Input v-model:value="createAccountForm.password" class="w-full" placeholder="Password"
                    :disabled="isLoading" />
                <Input v-model:value="createAccountForm.confirmPassword" class="w-full" placeholder="Confirm password"
                    :disabled="isLoading" />

                <Button @click="handleSignUp" text="Sign Up" :is-loading="isLoading" />
            </div>
            <div class="mt-2">
                already have an account? <b @click="emits('gotoSignIn')"
                    class="ml-2 font-bold underline cursor-pointer">Sign In</b>.
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import Input from "../components/Input.vue"
import Button from "../components/Button.vue"
import { ref } from "vue";
import { authService } from "../services/auth/auth";

const emits = defineEmits(["gotoSignIn"]);
const createAccountForm = ref({
    email: "",
    password: "",
    confirmPassword: "",
})

const isLoading = ref(false);

// FIXME: handle error
async function handleSignUp() {
    if (createAccountForm.value.password != createAccountForm.value.confirmPassword)
        return

    isLoading.value = true;

    try {
        await authService.signUp({
            email: createAccountForm.value.email,
            password: createAccountForm.value.password,
        })

        isLoading.value = false;
        emits("gotoSignIn")
    }
    catch {
        isLoading.value = false;
    }
}

</script>
