import { FC, FormEvent, useState } from 'react';
import { Box, Text } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';
import useCustomToast from '../../hooks/useCustomToast';
import { register } from '../../api/user';

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

const RegisterPage: FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [confirm, setConfirm] = useState('');
  const [loading, setLoading] = useState(false);
  const { showWarningToast, showSuccessToast } = useCustomToast();
  const navigate = useNavigate();

  const handleRegister = async (event: FormEvent) => {
    event.preventDefault();
    if (password !== confirm) {
      showWarningToast('两次密码输入不一致');
      return;
    }
    setLoading(true);
    try {
      const res = await register({ user: username, pass: password });
      if (res.code !== 0) {
        showWarningToast(res.msg);
        return;
      }
      showSuccessToast('注册成功，请登录');
      navigate('/login');
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box display="flex" alignItems="center" justifyContent="center" minH="calc(100vh - 56px)">
      <Box w="100%" maxW="360px" px={5}>
        <Text fontSize="28px" fontWeight={600} mb="6px">注册</Text>
        <Text fontSize="14px" color="#999" mb="40px">创建你的 GoBlog 账号</Text>

        <form onSubmit={handleRegister}>
          <Box mb={5}>
            <Text fontSize="13px" color="#666" mb="6px">用户名</Text>
            <input
              style={inputStyle}
              value={username}
              onChange={(e) => setUsername(e.target.value)}
              placeholder="3-20 个字符"
              autoComplete="username"
              onFocus={(e) => { e.target.style.borderColor = '#111'; }}
              onBlur={(e) => { e.target.style.borderColor = '#ddd'; }}
            />
          </Box>
          <Box mb={5}>
            <Text fontSize="13px" color="#666" mb="6px">密码</Text>
            <input
              style={inputStyle}
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              placeholder="至少 6 个字符"
              autoComplete="new-password"
              onFocus={(e) => { e.target.style.borderColor = '#111'; }}
              onBlur={(e) => { e.target.style.borderColor = '#ddd'; }}
            />
          </Box>
          <Box mb={8}>
            <Text fontSize="13px" color="#666" mb="6px">确认密码</Text>
            <input
              style={inputStyle}
              type="password"
              value={confirm}
              onChange={(e) => setConfirm(e.target.value)}
              placeholder="再次输入密码"
              autoComplete="new-password"
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
            {loading ? '注册中...' : '注册'}
          </button>
        </form>

        <Text fontSize="14px" color="#999" textAlign="center" mt="32px">
          已有账号？
          <Text
            as="span"
            color="#111"
            cursor="pointer"
            ml={1}
            _hover={{ textDecoration: 'underline' }}
            onClick={() => navigate('/login')}
          >
            登录
          </Text>
        </Text>
      </Box>
    </Box>
  );
};

export default RegisterPage;
