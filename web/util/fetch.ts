import { useRuntimeConfig } from "#app";
import { useServlicenseToken } from "./auth";

/**
 * Fetch with token
 */
export async function internalFetch<T>(
  url: string,
  method:
    | "GET"
    | "POST"
    | "PUT"
    | "DELETE"
    | "PATCH"
    | "OPTIONS"
    | "HEAD"
    | "CONNECT"
    | "TRACE" = "GET"
): Promise<T> {
  const config = useRuntimeConfig();
  const token = useServlicenseToken();

  const headers: HeadersInit = {};

  if (token && token.value) {
    headers.Authorization = `${token.value}`;
  }

  return await $fetch<T>(url, {
    baseURL: config.public.server,
    headers: {
      ...headers,
    },
    method: method,
  });
}
