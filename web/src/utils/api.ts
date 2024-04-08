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
  const error = ref<any>(null);
  const execute = async () => {
    error.value = null;
    loading.value = true;

    try {
      const res = await fetch('http://localhost:8080/health');
      const js = (await res.json()) as APIResponse<T>;

      data.value = js;
      error.value = null;
    } catch (e) {
      error.value = e;
    } finally {
      loading.value = false;
    }
  };
  const loading = ref(false);

  return { data, error, execute, loading };
}
