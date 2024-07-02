import axiosClient from "../utils/axiosClient.tsx";

export const login = async () => {
  const response = await axiosClient.get('/data'); // 更改为你的API路径
  return response.data;
};
