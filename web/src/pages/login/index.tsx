import { FC, FormEvent, useState } from 'react';
import { Box, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import useCustomToast from '../../hooks/useCustomToast';
import { login } from '../../api/user';

const inputStyle: React.CSSProperties = {
  width: '100%',
  height: '44px',
  padding: '0 12px',
  fontSize: '14px',
  border: '1px solid #ddd',
  borderRadius: '6px',
  background: '#fafafa',
  outline: 'none',
  color: '#111',
};

const LoginPage: FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [loading, setLoading] = useState(false);
  const { showWarningToast, showSuccessToast } = useCustomToast();
  const navigate = useNavigate();

  const handleLogin = async (event: FormEvent) => {
    event.preventDefault();
    if (!username || !password) {
      showWarningToast('用户名和密码是必填项');
      return;
    }
    setLoading(true);
    try {
      const res = await login({ user: username, pass: password });
      if (res.code !== 0) {
        showWarningToast(res.msg);
        return;
      }
      sessionStorage.setItem('auth_token', res.data.auth_token);
      sessionStorage.setItem('uid', res.data.uid);
      showSuccessToast('登录成功');
      navigate('/');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box display="flex" alignItems="center" justifyContent="center" minH="calc(100vh - 56px)">
      <Box w="100%" maxW="360px" px={5}>
        <Text fontSize="28px" fontWeight={600} mb="6px">登录</Text>
        <Text fontSize="14px" color="#999" mb="40px">登录你的 GoBlog 账号</Text>

        <form onSubmit={handleLogin}>
          <Box mb={5}>
            <Text fontSize="13px" color="#666" mb="6px">用户名</Text>
            <input
              style={inputStyle}
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              placeholder="输入用户名"
              autoComplete="username"
              onFocus={(e) => { e.target.style.borderColor = '#111'; }}
              onBlur={(e) => { e.target.style.borderColor = '#ddd'; }}
            />
          </Box>
          <Box mb={8}>
            <Text fontSize="13px" color="#666" mb="6px">密码</Text>
            <input
              style={inputStyle}
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="输入密码"
              autoComplete="current-password"
              onFocus={(e) => { e.target.style.borderColor = '#111'; }}
              onBlur={(e) => { e.target.style.borderColor = '#ddd'; }}
            />
          </Box>
          <button
            type="submit"
            disabled={loading}
            style={{
              width: '100%',
              height: '44px',
              background: '#111',
              color: '#fff',
              fontSize: '14px',
              fontWeight: 500,
              border: 'none',
              borderRadius: '6px',
              cursor: loading ? 'not-allowed' : 'pointer',
              opacity: loading ? 0.6 : 1,
            }}
          >
            {loading ? '登录中...' : '登录'}
          </button>
        </form>

        <Text fontSize="14px" color="#999" textAlign="center" mt="32px">
          没有账号？
          <Text
            as="span"
            color="#111"
            cursor="pointer"
            ml={1}
            _hover={{ textDecoration: 'underline' }}
            onClick={() => navigate('/register')}
          >
            注册
          </Text>
        </Text>
      </Box>
    </Box>
  );
};

export default LoginPage;
