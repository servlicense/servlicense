import type { Ref } from "vue";

import { useState, useCookie, useNuxtApp, type CookieRef } from "#app";

export interface ServlicenseAuth {
  id?: string;
  scopes: string[];
}

export const useServlicenseUser = (): Ref<ServlicenseAuth> =>
  useState<ServlicenseAuth>("servlicense.user");

// base64 encoded apikey
export const useServlicenseToken = (): CookieRef<string | null> => {
  const nuxtApp = useNuxtApp();

  nuxtApp._cookies = nuxtApp._cookies || {};
  if (nuxtApp._cookies.servlicense_api_key) {
    return nuxtApp._cookies.servlicense_api_key as CookieRef<string | null>;
  }

  const cookie = useCookie<string | null>("servlicense_api_key");
  nuxtApp._cookies.servlicense_api_key = cookie;
  return cookie;
};

// eslint-disable-next-line @typescript-eslint/explicit-function-return-type
export const useServlicenseAuthentication = () => {
  const authUser = useServlicenseUser();
  const token = useServlicenseToken();
  const runtimeConfig = useRuntimeConfig();
  const router = useRouter();

  const setUser = (user: ServlicenseAuth): void => {
    authUser.value = user;
  };

  const setApiKey = (cookie: string | null): void => {
    token.value = cookie;
  };

  const fetchAuthInfo = async (): Promise<ServlicenseAuth> => {
    try {
      const headers: HeadersInit = {};

      headers.Authorization = `${token.value}`;
      console.log(headers, token.value);

      const response: { data: ServlicenseAuth } = await $fetch(
        runtimeConfig.public.server + "/auth/me",
        {
          headers,
        }
      );

      setUser(response.data);
    } catch (error) {
      setApiKey(null);
    }

    return authUser.value;
  };

  const login = async (
    identifier: string,
    apikey: string
  ): Promise<ServlicenseAuth> => {
    const response: { data: { scopes: string[] } } = await $fetch(
      runtimeConfig.public.server + "/auth/me",
      {
        method: "GET",
        headers: {
          Authorization: `${btoa(`${identifier}:${apikey}`)}`,
        },
        onResponseError(error) {
          const data = error.response._data;

          if (!data.success) {
            throw new Error(
              "Login failed, please check your login credentials."
            );
          }

          if (!data.scopes) {
            throw new Error(
              "Login failed, please check your login credentials."
            );
          }
        },
      }
    );

    setApiKey(`${btoa(`${identifier}:${apikey}`)}`);
    const user = await fetchAuthInfo();
    return user;
  };

  const logout = async (): Promise<void> => {
    setApiKey(null);
    router.push("/login");
  };

  return {
    login,
    fetchAuthInfo,
    logout,
  };
};
