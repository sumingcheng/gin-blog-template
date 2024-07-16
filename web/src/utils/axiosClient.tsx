import axios from 'axios';

// 获取当前环境
const axiosClient = axios.create({
  baseURL: 'http://localhost:5678',
  headers: {
    'Content-Type': 'application/json'
  },
  withCredentials: import.meta.env.VITE_APP_ENV === 'development'  // 允许携带跨域cookies
});

// 请求拦截器
axiosClient.interceptors.request.use(
  config => {
    const auth_token = sessionStorage.getItem('auth_token');
    config.headers.auth_token = `${ auth_token }`;
    return config;
  },
  error => {
    return Promise.reject(error);
  }
);

// 响应拦截器
axiosClient.interceptors.response.use(
  response => response,
  error => {
    // 处理响应错误
    return Promise.reject(error);
  }
);

export default axiosClient;
