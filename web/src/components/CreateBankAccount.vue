<template>
    <div class="flex flex-col gap-4">
        <div class="flex gap-4 font-bold ">
            <div class="flex-1">
                <bank-icon class=" stroke-primary" />
            </div>
            <div class="flex-8 text-center">New Bank account</div>
            <div v-if="shouldHaveCloseButton" class="flex-1 text-xl hover:text-red-300 cursor-pointer text-end"
                @click="emits('closeButton')">
                X
            </div>
        </div>
        <Input :disabled="isLoading" v-model:value="createBankAccount.name" placeholder="name" />
        <Input :disabled="isLoading" v-model:value="createBankAccount.description" placeholder="description" />

        <div class="flex gap-4 justify-between w-full">
            <Button :is-loading="isLoading" bottom-color-type="secondary" class="flex-1" text="Cancel"
                @click="emits('cancelButton')" />
            <Button :is-loading="isLoading" class="flex-1" text="Save" @click="handleSave" />
        </div>
    </div>
</template>

<script setup lang='ts'>
import Input from "./Input.vue"
import Button from "./Button.vue"
import { ref, type Ref } from "vue";
import BankIcon from "../assets/BankIcon.vue";
import { BankAccountService, type CreateBankAccount } from "../services/bankAccounts/bankAccounts";

defineProps<{ shouldHaveCloseButton?: boolean }>()

//@ts-ignore lsp being crazy
const createBankAccount: Ref<CreateBankAccount> = ref({
    name: "",
    description: "",
})

const isLoading = ref(false);

const emits = defineEmits(["closeButton", "cancelButton", "created"])
async function handleSave() {
    isLoading.value = true;
    try {
        await BankAccountService.createBankAccount(createBankAccount.value);
        emits("created");
    } catch (e) {
        // FIXME: Handle errors
    }

    isLoading.value = false;
}
</script>
