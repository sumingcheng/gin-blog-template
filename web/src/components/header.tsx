import { FC, useRef, useState, useEffect } from 'react';
import { Box, Flex, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import { logout } from '../api/user';
import useCustomToast from '../hooks/useCustomToast';

const Header: FC = () => {
  const navigate = useNavigate();
  const { showSuccessToast } = useCustomToast();
  const isLoggedIn = !!sessionStorage.getItem('auth_token');
  const [menuOpen, setMenuOpen] = useState(false);
  const menuRef = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handler = (e: MouseEvent) => {
      if (menuRef.current && !menuRef.current.contains(e.target as Node)) {
        setMenuOpen(false);
      }
    };
    document.addEventListener('mousedown', handler);
    return () => document.removeEventListener('mousedown', handler);
  }, []);

  const handleLogout = async () => {
    setMenuOpen(false);
    try { await logout(); } finally {
      sessionStorage.removeItem('auth_token');
      sessionStorage.removeItem('uid');
      showSuccessToast('已退出');
      navigate('/login');
    }
  };

  const openSearch = () => {
    window.dispatchEvent(new CustomEvent('open-search'));
  };

  return (
    <Box
      as="header"
      position="sticky"
      top={0}
      zIndex={100}
      bg="#fff"
      borderBottom="1px solid #eee"
    >
      <Flex
        maxW="680px"
        mx="auto"
        px={5}
        h="56px"
        align="center"
        justify="space-between"
      >
        <Text
          fontSize="18px"
          fontWeight={600}
          cursor="pointer"
          onClick={() => navigate('/')}
          letterSpacing="-0.3px"
        >
          GoBlog
        </Text>

        <Flex align="center" gap={5}>
          {isLoggedIn ? (
            <>
              <Text
                fontSize="14px"
                color="#666"
                cursor="pointer"
                onClick={() => navigate('/blog/create')}
                _hover={{ color: '#111' }}
              >
                写文章
              </Text>
              <Box position="relative" ref={menuRef}>
                <Flex
                  w="32px"
                  h="32px"
                  borderRadius="50%"
                  bg="#111"
                  color="#fff"
                  align="center"
                  justify="center"
                  fontSize="13px"
                  fontWeight={500}
                  cursor="pointer"
                  onClick={() => setMenuOpen(!menuOpen)}
                  userSelect="none"
                >
                  U
                </Flex>
                {menuOpen && (
                  <Box
                    position="absolute"
                    right={0}
                    top="40px"
                    bg="#fff"
                    border="1px solid #eee"
                    borderRadius="6px"
                    py={1}
                    minW="120px"
                    zIndex={200}
                  >
                    <Text
                      fontSize="14px"
                      px={4}
                      py="8px"
                      cursor="pointer"
                      _hover={{ bg: '#fafafa' }}
                      onClick={() => { setMenuOpen(false); navigate('/profile'); }}
                    >
                      个人中心
                    </Text>
                    <Box h="1px" bg="#eee" />
                    <Text
                      fontSize="14px"
                      px={4}
                      py="8px"
                      color="#dc2626"
                      cursor="pointer"
                      _hover={{ bg: '#fafafa' }}
                      onClick={handleLogout}
                    >
                      退出登录
                    </Text>
                  </Box>
                )}
              </Box>
            </>
          ) : (
            <>
              <Text
                fontSize="14px"
                color="#666"
                cursor="pointer"
                onClick={() => navigate('/login')}
                _hover={{ color: '#111' }}
              >
                登录
              </Text>
              <Flex
                as="button"
                h="32px"
                px={4}
                bg="#111"
                color="#fff"
                fontSize="13px"
                fontWeight={500}
                borderRadius="6px"
                align="center"
                cursor="pointer"
                border="none"
                _hover={{ bg: '#333' }}
                onClick={() => navigate('/register')}
              >
                注册
              </Flex>
            </>
          )}
          <Text
            fontSize="13px"
            color="#ccc"
            cursor="pointer"
            onClick={openSearch}
            _hover={{ color: '#999' }}
            userSelect="none"
            ml={1}
          >
            ⌘K
          </Text>
        </Flex>
      </Flex>
    </Box>
  );
};

export default Header;
