import { ref } from 'vue';
import type { JSONCompatible, JSONValue } from '@/utils/json';

interface APIResponse<T> {
  data: T;
  error: string;
}

export interface EndpointConfig {
  method: string;
  path: string;
}

export type KeyVal = { [key: string]: any };

export interface RequestParams {
  path?: { [key: string]: any };
  query?: { [key: string]: any };
}

export interface RequestConfig {
  params?: RequestParams;
  body?: JSONValue;
}

export const HealthCheckConfig = {
  method: 'GET',
  path: '/health',
};

export interface HealthCheckResponse {
  status: string;
}

export interface UserGetRequest {
  params: {
    path: {
      username: string;
    };
  };
}

export const UserGetConfig = {
  method: 'GET',
  path: '/user/:username',
};

export interface UserGetResponse {
  email: {
    address: string;
    verified: boolean;
  };
  firstName: string;
  lastName: string;
  username: string;
}

export const UserGetCurrentConfig = {
  method: 'GET',
  path: '/user/me',
};

export interface UserGetCurrentResponse {
  email: string;
  firstName: string;
  lastName: string;
  username: string;
}

export const UserPostConfig = {
  method: 'POST',
  path: '/user',
};

export interface UserPostRequest {
  body: {
    firstName: string;
    id: string;
    lastName: string;
    username: string;
  };
}

export const UserPutConfig = {
  method: 'PUT',
  path: '/user/{:username}',
};

export interface UserPutRequest {
  body: {
    firstName?: string;
    lastName?: string;
    username?: string;
  };
  params: {
    path: {
      username: string;
    };
  };
}

export function useAPI<T, Out extends JSONCompatible<T>>(ec: EndpointConfig) {
  const data = ref<APIResponse<Out>>();
  const error = ref<any>(null);
  const loading = ref(false);

  async function execute(rc: RequestConfig) {
    error.value = null;
    loading.value = true;

    const path = 'http://localhost:8080' + getPathWithQuery(ec.path, rc.params || {});
    console.log('body', JSON.stringify(rc.body));
    const req = new Request(path, {
      method: ec.method,
      body: rc.body ? JSON.stringify(rc.body) : undefined,
    });

    try {
      const res = await fetch(req);

      const js = (await res.json()) as APIResponse<Out>;

      data.value = js;
      error.value = null;
    } catch (e) {
      error.value = e;
    } finally {
      loading.value = false;
    }
  }

  return { data, error, execute, loading };
}

function getPath(p: string, r: RequestParams) {
  let path = p;
  if (r.path) {
    Object.entries(r.path).forEach(([k, v]) => {
      path = path.replace(`:${k}`, v);
    });
  }

  return path;
}

function getPathWithQuery(p: string, r: RequestParams) {
  const path = getPath(p, r);
  const query = getQuery(r);

  return query ? `${path}?${query}` : path;
}

function getQuery(r: RequestParams) {
  return Object.entries(r.query || {})
    .map(([k, v]) => `${k}=${v}`)
    .join('&');
}
