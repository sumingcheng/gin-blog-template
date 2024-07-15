import axiosClient from "../utils/axiosClient.tsx";

// 获取博客是属于哪个用户
export const getBelong = async (data: {bid: number}) => {
  const res = await axiosClient({
    url: `/api/blog/belong`,
    method: 'POST',
    data
  })
  return res.data
};

// 获取博客列表
export const getBlogList = async (uid: number) => {
  const res = await axiosClient({
    url: `/api/blog/list/${ uid }`,
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
export const updateBlog = async (data: {blogId: number, title: string, article: string}) => {
  const res = await axiosClient({
    url: `/api/blog/update`,
    method: 'post',
    data,
  })
  return res.data
};
