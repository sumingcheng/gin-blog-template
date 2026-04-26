import axiosClient from '../utils/axiosClient';

interface ApiResponse<T = any> {
  code: number;
  msg: string;
  data?: T;
}

export const getBlogList = async (params: {
  page?: number;
  size?: number;
  uid?: number;
  keyword?: string;
}): Promise<ApiResponse> => {
  const res = await axiosClient.get('/api/blog/list', { params });
  return res.data;
};

export const getBlogDetail = async (bid: number): Promise<ApiResponse> => {
  const res = await axiosClient.get(`/api/blog/${bid}`);
  return res.data;
};

export const createBlog = async (data: { title: string; article: string }): Promise<ApiResponse> => {
  const res = await axiosClient.post('/api/blog/create', data);
  return res.data;
};

export const updateBlog = async (data: { blogId: number; title: string; article: string }): Promise<ApiResponse> => {
  const res = await axiosClient.post('/api/blog/update', data);
  return res.data;
};

export const deleteBlog = async (bid: number): Promise<ApiResponse> => {
  const res = await axiosClient.delete(`/api/blog/${bid}`);
  return res.data;
};

export const getBelong = async (data: { bid: number }): Promise<ApiResponse> => {
  const res = await axiosClient.post('/api/blog/belong', data);
  return res.data;
};
