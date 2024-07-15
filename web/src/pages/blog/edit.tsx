import { FC, useEffect, useState } from "react";
import { useNavigate, useSearchParams } from "react-router-dom";
import { Box, Button, FormControl, FormLabel, Input, Textarea } from "@chakra-ui/react";
import { getBlogDetail, updateBlog } from "../../api/blog.ts";
import useCustomToast from "../../hooks/useCustomToast.tsx";

const EditBlogPost: FC = () => {
  const { showSuccessToast, showWarningToast } = useCustomToast();
  const [title, setTitle] = useState("");
  const [article, setArticle] = useState("");
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();

  const id = searchParams.get("id") || "";

  const fetchBlogPost = async () => {
    try {
      const res = await getBlogDetail(id);
      setTitle(res.blogs[0].title);
      setArticle(res.blogs[0].article);
    } catch (error) {
      console.error('无法获取博客详细信息:', error);
      showWarningToast('无法获取博客详细信息');
    }
  };

  useEffect(() => {
    fetchBlogPost();
  }, [id]);

  const handleSubmit = async () => {
    try {
      await updateBlog({ blogId: Number(id), title, article });
      showSuccessToast('博客更新成功')
      navigate(`/blog`);
    } catch (error) {
      console.error('Failed to update blog:', error);
      showWarningToast('Failed to update blog.')
    }
  };

  return (
    <Box p={ 10 }>
      <FormControl>
        <FormLabel fontWeight="bold">博客标题</FormLabel>
        <Input placeholder={ "请输入标题" } value={ title } onChange={ (e) => setTitle(e.target.value) }/>
      </FormControl>
      <FormControl mt={ 4 }>
        <FormLabel fontWeight="bold">博客文章</FormLabel>
        <Textarea placeholder={ "请输入内容" } value={ article } onChange={ (e) => setArticle(e.target.value) }/>
      </FormControl>
      <Button mt={ 4 } colorScheme="blue" onClick={ handleSubmit }>
        保存
      </Button>
    </Box>
  );
};

export default EditBlogPost;
