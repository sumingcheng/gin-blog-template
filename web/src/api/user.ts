import axiosClient from '../utils/axiosClient';

interface ApiResponse<T = any> {
  code: number;
  msg: string;
  data?: T;
}

export const login = async (data: { user: string; pass: string }): Promise<ApiResponse> => {
  const res = await axiosClient.post('/api/login', data);
  return res.data;
};

export const register = async (data: { user: string; pass: string }): Promise<ApiResponse> => {
  const res = await axiosClient.post('/api/register', data);
  return res.data;
};

export const logout = async (): Promise<ApiResponse> => {
  const res = await axiosClient.post('/api/logout');
  return res.data;
};

export const getAuthToken = async (): Promise<ApiResponse> => {
  const res = await axiosClient.get('/api/token');
  return res.data;
};

export const getProfile = async (): Promise<ApiResponse> => {
  const res = await axiosClient.get('/api/user/profile');
  return res.data;
};

export const changePassword = async (data: { oldPass: string; newPass: string }): Promise<ApiResponse> => {
  const res = await axiosClient.post('/api/user/password', data);
  return res.data;
};
