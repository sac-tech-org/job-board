import { ref } from 'vue';

interface APIResponse<T> {
  data: T;
  error: string;
}

export interface HealthCheckResponse {
  status: string;
}

export function useAPI<T>() {
  const data = ref<APIResponse<T>>();
  const error = ref<any | null>(null);
  const execute = async () => {
    loading.value = true;

    try {
      const res = await fetch('http://localhost:8080/health');
      const js = (await res.json()) as APIResponse<T>;

      data.value = js;
      error.value = null;
      loading.value = false;
    } catch (e) {
      error.value = e;
      loading.value = false;
    }
  };
  const loading = ref(false);

  return { data, error, execute, loading };
}
