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
export const getToken = async (data: object[]) => {
  const res = await axiosClient({
    url: '/api/token',
    method: 'post',
    data,
  })
  return res.data
};

