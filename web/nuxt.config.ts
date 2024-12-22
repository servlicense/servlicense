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
		host: "0.0.0.0",
	},
	runtimeConfig: {
		public: {
			server: process.env.NODE_ENV === "production" ? "/api" : "http://localhost:3000/api",
		},
	},
});
