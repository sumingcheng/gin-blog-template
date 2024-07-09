import axios from 'axios';

const axiosClient = axios.create({
  baseURL: 'http://127.0.0.1:5678',
  headers: {
    'Content-Type': 'application/json',
  },
});

// 请求拦截器
axiosClient.interceptors.request.use(
  config => {
    const auth_token = sessionStorage.getItem('auth_token');
    config.headers.Authorization = `Bearer ${ auth_token }`;
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
