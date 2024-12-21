<template>
  <AppNavigation />
  <UContainer
    class="px-4 py-5 mx-auto sm:max-w-xl md:max-w-full lg:max-w-screen-xl md:px-24 lg:px-8"
  >
    <div class="flex flex-row justify-between items-center mb-8">
      <div>
        <h1 class="text-xl font-bold text-gray-900 dark:text-white">
          API Keys
        </h1>
      </div>
      <div class="flex flex-row items-center gap-3">
        <UButton color="primary" variant="solid" :disabled="!canCreate"
          >Create API Key</UButton
        >
      </div>
    </div>

    <UCard>
      <UTable :columns="columns" :rows="rows">
        <template #scopes-data="{ row }">
          <div class="flex flex-row gap-2">
            <UBadge
              v-for="scope in row.scopes"
              :key="scope"
              :color="scope === 'admin' ? 'primary' : 'gray'"
              variant="solid"
              >{{ scope }}</UBadge
            >
          </div>
        </template>
        <template #actions-data="{ row }">
          <UDropdown :items="actions(row)">
            <UButton
              color="gray"
              variant="ghost"
              icon="i-heroicons-ellipsis-horizontal-20-solid"
            />
          </UDropdown>
        </template>
      </UTable>
      <div
        class="flex justify-end px-3 py-3.5 border-t border-gray-200 dark:border-gray-700"
      >
        <UPagination
          v-model="page"
          :page-count="pageCount"
          :total="apiKeys.length"
        />
      </div>
    </UCard>
  </UContainer>
</template>
<script setup lang="ts">
import { useServlicenseUser } from "~/util/auth";
import { internalFetch } from "~/util/fetch";
import { useStorage } from "@vueuse/core";

interface ApiKey {
  id: number;
  name: string;
  scopes: string[];
  created_at: string;
}

const user = useServlicenseUser();

const columns = [
  { label: "Identifier", key: "id" },
  { label: "Name", key: "name" },
  { label: "Scopes", key: "scopes" },
  { label: "Created At", key: "created_at" },
  { label: "Actions", key: "actions" },
];
const page = ref(1);
const pageCount = 7;
const canCreate = computed(() => {
  if (!user.value) return false;

  if (!user.value.scopes) return false;

  if (
    user.value.scopes.includes("admin") ||
    user.value.scopes.includes("manage_api_keys")
  ) {
    return true;
  }

  return false;
});
const hideText = useStorage("hideText", false);

const apiKeys = ref<ApiKey[]>([]);

onMounted(async () => {
  const res = await internalFetch<{ data: { apiKeys: ApiKey[] } }>(
    "/auth/apikeys"
  );

  if (res.data.apiKeys) {
    apiKeys.value = res.data.apiKeys;
  }

  console.log(apiKeys.value);
});

const toast = useToast();
const actions = (row: ApiKey) => [
  [
    {
      label: "Show raw data",
      icon: "i-heroicons-document-text",
    },
    {
      label: "Copy Identifier",
      icon: "i-heroicons-finger-print",
      click: () => {
        navigator.clipboard.writeText(row.id.toString());
        toast.add({
          title: "Identifier copied",
          description: "The identifier has been copied to your clipboard",
          icon: "i-heroicons-finger-print",
          timeout: 3000,
        });
      },
    },
  ],
  [
    {
      label: "Revoke API Key",
      icon: "i-heroicons-trash-20-solid",
    },
  ],
];

const rows = computed(() => {
  return apiKeys.value.slice(
    (page.value - 1) * pageCount,
    page.value * pageCount
  );
});

definePageMeta({
  title: "API Keys - Servlicense",
  middleware: "auth",
});
</script>
<style lang="css" scoped>
#blur {
  color: transparent;
  text-shadow: 0 0 8px #000;
}
</style>
