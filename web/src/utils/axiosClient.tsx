import axios from 'axios';

const isDEV = import.meta.env.VITE_APP_ENV === 'development';

const axiosClient = axios.create({
  baseURL: isDEV ? 'http://localhost:5678' : '',
  headers: { 'Content-Type': 'application/json' },
  withCredentials: isDEV,
});

axiosClient.interceptors.request.use((config) => {
  const token = sessionStorage.getItem('auth_token');
  if (token) {
    config.headers.auth_token = token;
  }
  return config;
});

axiosClient.interceptors.response.use(
  (response) => response,
  (error) => {
    if (error.response?.status === 401) {
      sessionStorage.removeItem('auth_token');
    }
    return Promise.reject(error);
  }
);

export default axiosClient;
