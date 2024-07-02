import { Box, Heading, Text, useColorModeValue, VStack } from '@chakra-ui/react';
import { FC } from "react";

interface BlogPostProps {
  userId: string;
  title: string;
  article: string;
  updateTime: string;
}

const BlogPost: FC<BlogPostProps> = ({ userId, title, article, updateTime }) => {
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
        作者ID: { userId } | 更新时间: { updateTime }
      </Text>
      <Text fontSize="md">{ article }</Text>
    </VStack>
  );
};

const BlogPage: FC = () => {
  const blogPosts = [
    {
      userId: 'user123',
      title: '博客示例标题',
      article: '这里是博客文章的内容，可以详细介绍各种信息。',
      updateTime: '2023-07-02'
    },
    {
      userId: 'user456',
      title: '第二个博客示例',
      article: '另一篇文章的内容展示。',
      updateTime: '2023-07-03'
    },
  ];

  return (
    <Box p={ 8 } h={ "full" } bg={ useColorModeValue('white', 'gray.700') }>
      { blogPosts.map((post, index) => (
        <BlogPost key={ index } { ...post } />
      )) }
    </Box>
  );
};

export default BlogPage;
