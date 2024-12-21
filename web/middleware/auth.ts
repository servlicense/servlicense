import { useServlicenseAuthentication, useServlicenseUser } from "~/util/auth";

export default defineNuxtRouteMiddleware(async () => {
  const user = useServlicenseUser();
  const { fetchAuthInfo } = useServlicenseAuthentication();
  const toast = useToast();

  await fetchAuthInfo();

  if (!user.value) {
    toast.add({
      title: "Unauthenticated",
      description: "You must be logged in to access this page.",
      icon: "i-heroicons-lock-closed",
      timeout: 5000,
    });
    return navigateTo("/login");
  } else {
    console.log("Authenticated");
  }
});
