// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
	devtools: { enabled: false },
	modules: ["@nuxt/ui"],
	compatibilityDate: "2024-12-19",
	colorMode: {
		fallback: "light",
		preference: "light",
		storageKey: "colormode-servlicense",
	},
	devServer: {
		port: 4000,
	},
	runtimeConfig: {
		public: {
			server: process.env.NODE_ENV === "production" ? "" : "http://localhost:3000",
		},
	},
});
