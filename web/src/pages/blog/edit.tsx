import { FC, useEffect, useState } from 'react';
import { Box, Flex, Text } from '@chakra-ui/react';
import { useNavigate, useParams } from 'react-router-dom';
import { getBlogDetail, updateBlog } from '../../api/blog';
import useCustomToast from '../../hooks/useCustomToast';

const EditBlogPage: FC = () => {
  const { bid } = useParams<{ bid: string }>();
  const [title, setTitle] = useState('');
  const [article, setArticle] = useState('');
  const [loading, setLoading] = useState(false);
  const [fetching, setFetching] = useState(true);
  const { showSuccessToast, showWarningToast } = useCustomToast();
  const navigate = useNavigate();

  useEffect(() => {
    if (!bid) { return; }
    setFetching(true);
    getBlogDetail(Number(bid)).then((res) => {
      if (res.code === 0 && res.data) {
        setTitle(res.data.title);
        setArticle(res.data.article);
      }
    }).finally(() => setFetching(false));
  }, [bid]);

  const handleSubmit = async () => {
    if (!title.trim() || !article.trim()) {
      showWarningToast('标题和内容不能为空');
      return;
    }
    setLoading(true);
    try {
      const res = await updateBlog({ blogId: Number(bid), title, article });
      if (res.code === 0) {
        showSuccessToast('已保存');
        navigate('/');
      } else {
        showWarningToast(res.msg);
      }
    } finally {
      setLoading(false);
    }
  };

  if (fetching) {
    return (
      <Box maxW="680px" mx="auto" px={5} pt="80px" opacity={0.3}>
        <Box h="28px" w="50%" bg="#eee" borderRadius="3px" mb="32px" />
        <Box h="14px" w="100%" bg="#eee" borderRadius="3px" mb={3} />
        <Box h="14px" w="80%" bg="#eee" borderRadius="3px" />
      </Box>
    );
  }

  return (
    <Box maxW="680px" mx="auto" px={5} pt="80px" pb="120px">
      <Flex justify="space-between" align="center" mb="48px">
        <Text
          fontSize="13px"
          color="#999"
          cursor="pointer"
          _hover={{ color: '#111' }}
          onClick={() => navigate(-1)}
        >
          ← 返回
        </Text>
        <button
          onClick={handleSubmit}
          disabled={loading}
          style={{
            height: '32px',
            padding: '0 20px',
            background: '#111',
            color: '#fff',
            fontSize: '13px',
            fontWeight: 500,
            border: 'none',
            borderRadius: '6px',
            cursor: loading ? 'not-allowed' : 'pointer',
            opacity: loading ? 0.6 : 1,
          }}
        >
          {loading ? '保存中...' : '保存'}
        </button>
      </Flex>

      <input
        value={title}
        onChange={(e) => setTitle(e.target.value)}
        placeholder="标题"
        style={{
          width: '100%',
          fontSize: '28px',
          fontWeight: 600,
          border: 'none',
          outline: 'none',
          letterSpacing: '-0.3px',
          color: '#111',
          marginBottom: '32px',
        }}
      />

      <textarea
        value={article}
        onChange={(e) => setArticle(e.target.value)}
        placeholder="文章内容"
        style={{
          width: '100%',
          minHeight: '400px',
          fontSize: '16px',
          lineHeight: '2',
          border: 'none',
          outline: 'none',
          resize: 'none',
          color: '#333',
        }}
      />

      <Text fontSize="12px" color="#ccc" textAlign="right" mt={4}>
        {article.length} 字
      </Text>
    </Box>
  );
};

export default EditBlogPage;
