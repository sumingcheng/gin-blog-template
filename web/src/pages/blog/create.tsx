import { FC, useState } from 'react';
import { Box, Flex, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import { createBlog } from '../../api/blog';
import useCustomToast from '../../hooks/useCustomToast';

const CreateBlogPage: FC = () => {
  const [title, setTitle] = useState('');
  const [article, setArticle] = useState('');
  const [loading, setLoading] = useState(false);
  const { showSuccessToast, showWarningToast } = useCustomToast();
  const navigate = useNavigate();

  const handleSubmit = async () => {
    if (!title.trim() || !article.trim()) {
      showWarningToast('标题和内容不能为空');
      return;
    }
    setLoading(true);
    try {
      const res = await createBlog({ title, article });
      if (res.code === 0) {
        showSuccessToast('发布成功');
        navigate('/');
      } else {
        showWarningToast(res.msg);
      }
    } finally {
      setLoading(false);
    }
  };

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
          {loading ? '发布中...' : '发布'}
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
        placeholder="写点什么..."
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

export default CreateBlogPage;
