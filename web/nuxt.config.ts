// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ["@nuxt/ui"],
  compatibilityDate: "2024-12-19",
  colorMode: {
    fallback: "light",
    preference: "light",
    storageKey: "colormode-servlicense",
  },
});
