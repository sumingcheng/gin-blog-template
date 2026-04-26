import { FC, useEffect, useState } from 'react';
import { Box, Text } from '@chakra-ui/react';
import { useNavigate, useParams } from 'react-router-dom';
import { getBlogDetail } from '../../api/blog';

interface Blog {
  id: number;
  userId: number;
  title: string;
  article: string;
  createdAt: number;
  updateAt: number;
}

const formatDate = (ts: number) => {
  if (!ts) { return ''; }
  const d = new Date(ts * 1000);
  return d.toLocaleDateString('zh-CN', { year: 'numeric', month: 'long', day: 'numeric' });
};

const BlogDetailPage: FC = () => {
  const { bid } = useParams<{ bid: string }>();
  const [blog, setBlog] = useState<Blog | null>(null);
  const [loading, setLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    if (!bid) { return; }
    setLoading(true);
    getBlogDetail(Number(bid)).then((res) => {
      if (res.code === 0) { setBlog(res.data); }
    }).finally(() => setLoading(false));
  }, [bid]);

  if (loading) {
    return (
      <Box maxW="680px" mx="auto" px={5} pt="80px" opacity={0.3}>
        <Box h="36px" w="60%" bg="#eee" borderRadius="3px" mb={4} />
        <Box h="14px" w="150px" bg="#eee" borderRadius="3px" mb="48px" />
        <Box h="14px" w="100%" bg="#eee" borderRadius="3px" mb={3} />
        <Box h="14px" w="90%" bg="#eee" borderRadius="3px" mb={3} />
        <Box h="14px" w="70%" bg="#eee" borderRadius="3px" />
      </Box>
    );
  }

  if (!blog) {
    return (
      <Box maxW="680px" mx="auto" px={5} pt="80px">
        <Text color="#999" fontSize="14px">文章不存在</Text>
      </Box>
    );
  }

  return (
    <Box maxW="680px" mx="auto" px={5} pt="80px" pb="120px">
      <Text
        fontSize="13px"
        color="#999"
        cursor="pointer"
        mb="48px"
        _hover={{ color: '#111' }}
        onClick={() => navigate(-1)}
      >
        ← 返回
      </Text>

      <Text
        as="h1"
        fontSize="36px"
        fontWeight={600}
        lineHeight="1.3"
        letterSpacing="-0.5px"
        mb="16px"
      >
        {blog.title}
      </Text>

      <Text fontSize="13px" color="#999" mb="48px">
        {formatDate(blog.createdAt)}
        {blog.updateAt > 0 ? ` · 编辑于 ${formatDate(blog.updateAt)}` : ''}
      </Text>

      <Text
        fontSize="16px"
        color="#333"
        lineHeight="2.0"
        whiteSpace="pre-wrap"
      >
        {blog.article}
      </Text>

      <Box h="1px" bg="#eee" my="60px" />

      <Text
        fontSize="13px"
        color="#999"
        cursor="pointer"
        _hover={{ color: '#111' }}
        onClick={() => navigate('/')}
      >
        ← 回到首页
      </Text>
    </Box>
  );
};

export default BlogDetailPage;
