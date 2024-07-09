import { Box, Heading, Text, useColorModeValue, VStack } from '@chakra-ui/react';
import { FC, memo, useCallback, useEffect, useState } from "react";
import { getBlogList } from "../../api/blog";

interface BlogPostProps {
  id: number;
  userId: number;
  title: string;
  article: string;
  updateTime: string;
}

const BlogPost: FC<BlogPostProps> = memo(({ userId, title, article, updateTime }) => {
  return (
    <VStack
      spacing={ 2 }
      p={ 6 }
      boxShadow="xs"
      rounded="md"
      align="left"
      marginBottom={ 4 }
    >
      <Heading size="md">{ title }</Heading>
      <Text fontSize="sm" color={ useColorModeValue('gray.600', 'gray.400') }>
        作者ID: { userId } | 更新时间: { new Date(updateTime).toLocaleString() }
      </Text>
      <Text fontSize="md">{ article }</Text>
    </VStack>
  );
});

const BlogPage: FC = () => {
  const [blogPosts, setBlogPosts] = useState<BlogPostProps[]>([]);

  const fetchBlogs = useCallback(async () => {
    try {
      const res = await getBlogList(1);
      if (res) {
        setBlogPosts(res.blogs);
      }
    } catch (error) {
      console.error('Failed to fetch blogs:', error);
    }
  }, []);

  useEffect(() => {
    fetchBlogs();
  }, [fetchBlogs]);

  return (
    <Box p={ 8 } h="full" bg={ useColorModeValue('white', 'gray.700') }>
      { blogPosts.map((post) => (
        <BlogPost key={ post.id } { ...post } />
      )) }
    </Box>
  );
};

export default BlogPage;
