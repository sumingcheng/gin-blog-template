import axiosClient from "../utils/axiosClient.tsx";

// 获取博客是属于哪个用户
export const getBlogUser = async () => {
  const res = await axiosClient({
    url: `/api/blog/belong`,
    method: 'get',
  })
  return res.data
};

// 获取博客列表
export const getBlogList = async () => {
  const res = await axiosClient({
    url: '/api/blog/list',
    method: 'get',
  })
  return res.data
};

// 获取博客详情
export const getBlogDetail = async (bid: string) => {
  const res = await axiosClient({
    url: `/api/blog/${ bid }`,
    method: 'get',
  })
  return res.data
};

// 更新博客
export const updateBlog = async (data: object[]) => {
  const res = await axiosClient({
    url: `/api/blog/update`,
    method: 'post',
    data,
  })
  return res.data
};
