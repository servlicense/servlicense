<template>
  <AppNavigation />
  <UContainer
    class="px-4 py-5 mx-auto sm:max-w-xl md:max-w-full lg:max-w-screen-xl md:px-24 lg:px-8"
  >
    <div class="flex flex-row justify-between items-center mb-8">
      <h1 class="text-xl font-bold text-gray-900 dark:text-white">Licenses</h1>
      <div class="flex flex-row items-center gap-3">
        <UButton color="primary" variant="solid">Create License</UButton>
        <UTooltip v-if="!hideText" text="Hide Licenses">
          <UButton
            variant="solid"
            color="gray"
            icon="i-heroicons-eye-slash"
            @click="hideText = true"
          />
        </UTooltip>
        <UTooltip v-else text="Show Licenses">
          <UButton
            variant="solid"
            color="gray"
            icon="i-heroicons-eye"
            @click="hideText = false"
          />
        </UTooltip>
      </div>
    </div>

    <UCard>
      <UTable :columns="columns" :rows="rows">
        <template #license-data="{ row }">
          <span v-if="!hideText" class="transition-opacity duration-400">{{
            row.license
          }}</span>
          <span v-else
            >{{ row.license.slice(0, 7) }}
            <span id="blur" class="transition-opacity duration-400">{{
              row.license.slice(7, row.license.length)
            }}</span></span
          >
        </template>
        <template #status-data="{ row }">
          <UTooltip text="x" :popper="{ arrow: true, placement: 'top' }">
            <UBadge :color="row.active ? 'primary' : 'gray'" variant="solid">{{
              row.active ? "Active" : "Revoked"
            }}</UBadge>
          </UTooltip>
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
          :total="licenses.length"
        />
      </div>
    </UCard>
  </UContainer>
</template>
<script setup lang="ts">
import { useServlicenseUser } from "~/util/auth";
import { internalFetch } from "~/util/fetch";

interface License {
  license: string;
  active: boolean;
  valid_until?: string;
  created_at: string;
  updated_at: string;
}

const columns = [
  { key: "license", label: "License" },
  { key: "status", label: "Status" },
  { key: "valid_until", label: "Valid Until" },
  { key: "created_at", label: "Created At" },
  { key: "updated_at", label: "Updated At" },
  { key: "actions", label: "Actions" },
];
const page = ref(1);
const pageCount = 7;

const hideText = ref(false);

const user = useServlicenseUser();
const licenses = ref<License[]>([]);

onMounted(async () => {
  const res = await internalFetch<{ data: { licenses: License[] } }>(
    "/licenses"
  );

  if (res.data.licenses) {
    licenses.value = res.data.licenses;
  }

  console.log(licenses.value);
});

const toast = useToast();
const actions = (row: License) => [
  [
    {
      label: "Show raw data",
      icon: "i-heroicons-document-text",
    },
    {
      label: "Copy License",
      icon: "i-heroicons-finger-print",
      click: () => {
        navigator.clipboard.writeText(row.license);
        toast.add({
          title: "License copied",
          description: "The license has been copied to your clipboard",
          icon: "i-heroicons-finger-print",
          timeout: 3000,
        });
      },
    },
  ],
  [
    {
      label: "Revoke License",
      icon: "i-heroicons-trash-20-solid",
    },
  ],
];

const rows = computed(() => {
  return licenses.value.slice(
    (page.value - 1) * pageCount,
    page.value * pageCount
  );
});

definePageMeta({
  title: "Home Page",
  description: "This is the home page",
  middleware: "auth",
});
</script>
<style lang="css" scoped>
#blur {
  color: transparent;
  text-shadow: 0 0 8px #000;
}
</style>
