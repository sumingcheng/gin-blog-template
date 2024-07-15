import { Box, Button, Flex, Heading, Text, useColorModeValue, VStack } from '@chakra-ui/react';
import { FC, memo, useCallback, useEffect, useState } from "react";
import { getBlogList } from "../../api/blog";
import { useNavigate } from "react-router-dom";
import useCustomToast from "../../hooks/useCustomToast.tsx";

interface BlogPostProps {
  id: number;
  userId: number;
  title: string;
  article: string;
  updateTime: string;
}

const BlogPost: FC<BlogPostProps> = memo(({ id, userId, title, article, updateTime }) => {
  const navigate = useNavigate();
  const edit = ((id: number) => {
    navigate(`/edit/?id=${ id }`);
  })

  return (
    <VStack
      spacing={ 2 }
      p={ 6 }
      boxShadow="xs"
      rounded="md"
      align="left"
      marginBottom={ 4 }
      w="full"
    >
      <Flex justify="space-between" w="full" alignItems="center">
        <Heading size="md">{ title }</Heading>
        <Button onClick={ () => {
          edit(id)
        } } size="sm">修改</Button>
      </Flex>
      <Text fontSize="sm" color={ useColorModeValue('gray.600', 'gray.400') }>
        作者ID: { userId } | 更新时间: { new Date(updateTime).toLocaleString() }
      </Text>
      <Text fontSize="md">{ article }</Text>
    </VStack>
  );
});

const BlogPage: FC = () => {
  const [blogPosts, setBlogPosts] = useState<BlogPostProps[]>([]);
  const { showWarningToast } = useCustomToast();

  const fetchBlogs = useCallback(async () => {
    try {
      const res = await getBlogList(1);
      if (res) {
        setBlogPosts(res.blogs);
      }
      if (res.code == 403) {
        showWarningToast(res.msg);
      }
    } catch (error) {
      console.log('无法获取博客:', error);
    }
  }, []);

  useEffect(() => {
    fetchBlogs();
  }, [fetchBlogs]);

  return (
    <Box p={ 8 } h="full" bg={ useColorModeValue('white', 'gray.700') }>
      { blogPosts && blogPosts.map((post) => (
        <BlogPost key={ post.id } { ...post } />
      )) }
    </Box>
  );
};

export default BlogPage;
