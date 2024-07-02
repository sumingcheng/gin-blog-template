// src/api/axiosClient.js
import axios from 'axios';

const axiosClient = axios.create({
  baseURL: 'http://127.0.0.1',
  headers: {
    'Content-Type': 'application/json',
    'Accept': 'application/json',
  },
});

// 请求拦截器
axiosClient.interceptors.request.use(
  config => {
    const token = sessionStorage.getItem('token');
    config.headers.Authorization = `Bearer ${ token }`;
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
