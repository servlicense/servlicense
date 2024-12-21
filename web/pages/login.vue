<script setup lang="ts">
import { z } from "zod";
import type { FormSubmitEvent } from "#ui/types";
import { useServlicenseAuthentication } from "~/util/auth";

const schema = z.object({
  identifier: z.string(),
  apiKey: z.string().min(8, "Must be at least 8 characters"),
});

type Schema = z.output<typeof schema>;

const state = reactive({
  identifier: undefined,
  apiKey: undefined,
});

const { login } = useServlicenseAuthentication();
const toast = useToast();
const router = useRouter();

async function onSubmit(event: FormSubmitEvent<Schema>) {
  console.log("onSubmit", event.data);
  try {
    await login(event.data.identifier, event.data.apiKey);
    router.push("/");
    toast.add({
      title: "Welcome back to Servlicense!",
      description: "You have successfully logged in.",
      icon: "i-heroicons-lock-open",
      timeout: 5000,
    });
  } catch (e: unknown) {
    console.log(e);
  }
}
</script>

<template>
  <UContainer class="flex h-screen justify-center items-center">
    <UCard class="w-full md:w-1/2">
      <template #header>
        <h1 class="font-bold">Servlicense - Client Login</h1>
      </template>
      <UForm
        :schema="schema"
        :state="state"
        class="space-y-4"
        @submit="onSubmit"
      >
        <UFormGroup label="Identifier" name="identifier">
          <UInput
            v-model="state.identifier"
            icon="i-heroicons-identification"
            placeholder="Identifier"
          />
        </UFormGroup>

        <UFormGroup label="API Key" name="apiKey">
          <UInput
            v-model="state.apiKey"
            type="apiKey"
            icon="i-heroicons-key"
            placeholder="API Key"
          />
        </UFormGroup>
        <UButton type="submit"> Login </UButton>
      </UForm>

      <template #footer class="flex justify-end">
        <p class="text-sm text-neutral-500">
          Note: The client creates a Base64 encoded Authorization header out of
          this credentials, it will be stored in your browser for processing the
          Requests.
        </p>
      </template>
    </UCard>
  </UContainer>
</template>
