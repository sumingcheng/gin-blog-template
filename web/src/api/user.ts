import axiosClient from "../utils/axiosClient.tsx";

// 登录
export const login = async (data: {user: string, pass: string}) => {
  const res = await axiosClient({
    url: '/api/login',
    method: 'post',
    data
  })
  return res.data
};

// 获取token
export const getAuthToken = async () => {
  const res = await axiosClient({
    url: '/api/token',
    method: 'get',
  })
  return res.data
};

// 退出登录
export const logout = async () => {
  const res = await axiosClient({
    url: '/api/logout',
    method: 'get',
  })
  return res.data
};
