import { FC, useCallback, useEffect, useRef, useState } from 'react';
import { Box, Flex, Text } from '@chakra-ui/react';
import { useNavigate, useSearchParams } from 'react-router-dom';
import { deleteBlog, getBlogList } from '../../api/blog';
import useCustomToast from '../../hooks/useCustomToast';

interface Blog {
  id: number;
  userId: number;
  title: string;
  article: string;
  createdAt: number;
}

const formatDate = (ts: number) => {
  if (!ts) { return ''; }
  const d = new Date(ts * 1000);
  return `${d.getFullYear()}-${String(d.getMonth() + 1).padStart(2, '0')}-${String(d.getDate()).padStart(2, '0')}`;
};

const HomePage: FC = () => {
  const [blogs, setBlogs] = useState<Blog[]>([]);
  const [total, setTotal] = useState(0);
  const [page, setPage] = useState(1);
  const [keyword, setKeyword] = useState('');
  const [searchOpen, setSearchOpen] = useState(false);
  const [searchInput, setSearchInput] = useState('');
  const [loading, setLoading] = useState(true);
  const { showSuccessToast, showWarningToast } = useCustomToast();
  const navigate = useNavigate();
  const [searchParams] = useSearchParams();
  const searchRef = useRef<HTMLInputElement>(null);
  const size = 10;

  const isLoggedIn = !!sessionStorage.getItem('auth_token');
  const currentUid = Number(sessionStorage.getItem('uid') || 0);

  const fetchBlogs = useCallback(async () => {
    setLoading(true);
    try {
      const uidParam = searchParams.get('uid');
      const res = await getBlogList({
        page, size,
        keyword: keyword || undefined,
        uid: uidParam ? Number(uidParam) : undefined,
      });
      if (res.code === 0 && res.data) {
        setBlogs(res.data.list || []);
        setTotal(res.data.total || 0);
      }
    } catch (e) {
      console.error(e);
    } finally {
      setLoading(false);
    }
  }, [page, keyword, searchParams]);

  useEffect(() => { fetchBlogs(); }, [fetchBlogs]);

  // Cmd+K / Ctrl+K 或 Header 按钮唤起搜索
  useEffect(() => {
    const onKey = (e: KeyboardEvent) => {
      if ((e.metaKey || e.ctrlKey) && e.key === 'k') {
        e.preventDefault();
        setSearchOpen(true);
        setSearchInput(keyword);
      }
      if (e.key === 'Escape' && searchOpen) {
        setSearchOpen(false);
      }
    };
    const onCustom = () => {
      setSearchOpen(true);
      setSearchInput(keyword);
    };
    window.addEventListener('keydown', onKey);
    window.addEventListener('open-search', onCustom);
    return () => {
      window.removeEventListener('keydown', onKey);
      window.removeEventListener('open-search', onCustom);
    };
  }, [searchOpen, keyword]);

  useEffect(() => {
    if (searchOpen && searchRef.current) {
      searchRef.current.focus();
    }
  }, [searchOpen]);

  const submitSearch = () => {
    setKeyword(searchInput);
    setPage(1);
    setSearchOpen(false);
  };

  const clearSearch = () => {
    setKeyword('');
    setSearchInput('');
    setPage(1);
    setSearchOpen(false);
  };

  const totalPages = Math.ceil(total / size);

  const handleDelete = async (bid: number) => {
    if (!confirm('确定删除？')) { return; }
    const res = await deleteBlog(bid);
    if (res.code === 0) {
      showSuccessToast('已删除');
      fetchBlogs();
    } else {
      showWarningToast(res.msg);
    }
  };

  return (
    <>
      {/* 搜索模态框 */}
      {searchOpen && (
        <Box
          position="fixed"
          top={0}
          left={0}
          right={0}
          bottom={0}
          bg="rgba(0,0,0,0.3)"
          zIndex={1000}
          display="flex"
          alignItems="flex-start"
          justifyContent="center"
          pt="20vh"
          onClick={(e) => { if (e.target === e.currentTarget) { setSearchOpen(false); } }}
        >
          <Box
            bg="#fff"
            borderRadius="8px"
            w="100%"
            maxW="480px"
            mx={5}
            overflow="hidden"
          >
            <Flex align="center" px={5} h="56px" borderBottom="1px solid #eee">
              <Text fontSize="14px" color="#999" mr={3} flexShrink={0}>搜索</Text>
              <input
                ref={searchRef}
                value={searchInput}
                onChange={(e) => setSearchInput(e.target.value)}
                onKeyDown={(e) => {
                  if (e.key === 'Enter') { submitSearch(); }
                  if (e.key === 'Escape') { setSearchOpen(false); }
                }}
                placeholder="输入关键词..."
                style={{
                  flex: 1,
                  fontSize: '14px',
                  border: 'none',
                  outline: 'none',
                  background: 'transparent',
                  color: '#111',
                }}
              />
              <Text fontSize="12px" color="#ccc" ml={3} flexShrink={0} userSelect="none">ESC</Text>
            </Flex>
            {keyword && (
              <Flex px={5} py={3} justify="space-between" align="center">
                <Text fontSize="13px" color="#999">当前搜索：{keyword}</Text>
                <Text
                  fontSize="13px"
                  color="#999"
                  cursor="pointer"
                  _hover={{ color: '#111' }}
                  onClick={clearSearch}
                >
                  清除
                </Text>
              </Flex>
            )}
          </Box>
        </Box>
      )}

      <Box maxW="680px" mx="auto" px={5} pt="48px" pb="120px">
        {keyword && (
          <Flex justify="flex-start" mb="24px" align="center" gap={3}>
            <Text fontSize="13px" color="#999">搜索：{keyword}</Text>
            <Text fontSize="13px" color="#999" cursor="pointer" _hover={{ color: '#111' }} onClick={clearSearch}>清除</Text>
          </Flex>
        )}

        {/* 列表 */}
        {loading ? (
          <Box>
            {[0, 1, 2, 3].map((i) => (
              <Box key={i} py="20px" borderBottom="1px solid #eee" opacity={0.4}>
                <Box h="16px" w="50%" bg="#eee" borderRadius="3px" mb="8px" />
                <Box h="12px" w="80%" bg="#eee" borderRadius="3px" mb="6px" />
                <Box h="12px" w="30%" bg="#eee" borderRadius="3px" />
              </Box>
            ))}
          </Box>
        ) : blogs.length === 0 ? (
          <Text color="#999" fontSize="14px" pt="60px" textAlign="center">
            {keyword ? `"${keyword}" 没有找到相关文章` : '暂无文章'}
          </Text>
        ) : (
          <Box>
            {blogs.map((blog, idx) => (
              <Box
                key={blog.id}
                py="20px"
                borderBottom={idx < blogs.length - 1 ? '1px solid #eee' : 'none'}
                position="relative"
                pl="16px"
                ml="-16px"
                cursor="pointer"
                role="group"
                onClick={() => navigate(`/blog/${blog.id}`)}
                sx={{
                  transition: 'padding-left 0.25s cubic-bezier(0.25, 0.1, 0.25, 1)',
                  '&::before': {
                    content: '""',
                    position: 'absolute',
                    left: 0,
                    top: '24px',
                    bottom: '24px',
                    width: '2px',
                    bg: '#111',
                    borderRadius: '1px',
                    transform: 'scaleY(0)',
                    transition: 'transform 0.25s cubic-bezier(0.25, 0.1, 0.25, 1)',
                    transformOrigin: 'center',
                  },
                  '&:hover': {
                    pl: '24px',
                  },
                  '&:hover::before': {
                    transform: 'scaleY(1)',
                  },
                }}
              >
                <Flex justify="space-between" align="flex-start">
                  <Box flex={1}>
                    <Text
                      fontSize="16px"
                      fontWeight={600}
                      lineHeight="1.5"
                      mb="6px"
                      transition="letter-spacing 0.25s ease"
                      _groupHover={{ letterSpacing: '0.2px' }}
                    >
                      {blog.title}
                    </Text>
                    <Text
                      fontSize="13px"
                      color="#666"
                      lineHeight="1.6"
                      noOfLines={1}
                      mb="6px"
                      transition="color 0.2s ease"
                      _groupHover={{ color: '#999' }}
                    >
                      {blog.article}
                    </Text>
                    <Flex align="center" gap="6px">
                      <Text fontSize="12px" color="#999">{formatDate(blog.createdAt)}</Text>
                      <Text
                        fontSize="12px"
                        color="#ccc"
                        opacity={0}
                        transition="opacity 0.2s ease"
                        _groupHover={{ opacity: 1 }}
                      >
                        →
                      </Text>
                    </Flex>
                  </Box>

                  {isLoggedIn && currentUid === blog.userId && (
                    <Flex
                      gap={3}
                      ml={6}
                      pt="2px"
                      flexShrink={0}
                      opacity={0}
                      transition="opacity 0.2s ease"
                      _groupHover={{ opacity: 1 }}
                    >
                      <Text
                        fontSize="12px"
                        color="#999"
                        cursor="pointer"
                        _hover={{ color: '#111' }}
                        onClick={(e) => { e.stopPropagation(); navigate(`/blog/edit/${blog.id}`); }}
                      >
                        编辑
                      </Text>
                      <Text
                        fontSize="12px"
                        color="#999"
                        cursor="pointer"
                        _hover={{ color: '#dc2626' }}
                        onClick={(e) => { e.stopPropagation(); handleDelete(blog.id); }}
                      >
                        删除
                      </Text>
                    </Flex>
                  )}
                </Flex>
              </Box>
            ))}
          </Box>
        )}

        {/* 分页 */}
        {totalPages > 1 && (
          <Flex justify="center" gap={6} mt="48px">
            <Text
              fontSize="13px"
              color={page <= 1 ? '#ccc' : '#999'}
              cursor={page <= 1 ? 'default' : 'pointer'}
              onClick={() => { if (page > 1) { setPage(page - 1); } }}
              _hover={page > 1 ? { color: '#111' } : undefined}
            >
              ← 上一页
            </Text>
            <Text fontSize="13px" color="#ccc">{page} / {totalPages}</Text>
            <Text
              fontSize="13px"
              color={page >= totalPages ? '#ccc' : '#999'}
              cursor={page >= totalPages ? 'default' : 'pointer'}
              onClick={() => { if (page < totalPages) { setPage(page + 1); } }}
              _hover={page < totalPages ? { color: '#111' } : undefined}
            >
              下一页 →
            </Text>
          </Flex>
        )}
      </Box>
    </>
  );
};

export default HomePage;
