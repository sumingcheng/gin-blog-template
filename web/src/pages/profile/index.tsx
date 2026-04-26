import { FC, FormEvent, useEffect, useState } from 'react';
import { Box, Flex, Text } from '@chakra-ui/react';
import { changePassword, getProfile } from '../../api/user';
import useCustomToast from '../../hooks/useCustomToast';

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

const ProfilePage: FC = () => {
  const [user, setUser] = useState<{ id: number; name: string } | null>(null);
  const [oldPass, setOldPass] = useState('');
  const [newPass, setNewPass] = useState('');
  const [confirmPass, setConfirmPass] = useState('');
  const [loading, setLoading] = useState(false);
  const { showSuccessToast, showWarningToast } = useCustomToast();

  useEffect(() => {
    getProfile().then((res) => {
      if (res.code === 0 && res.data) { setUser(res.data); }
    });
  }, []);

  const handleChangePassword = async (e: FormEvent) => {
    e.preventDefault();
    if (newPass !== confirmPass) {
      showWarningToast('两次密码输入不一致');
      return;
    }
    setLoading(true);
    try {
      const res = await changePassword({ oldPass, newPass });
      if (res.code === 0) {
        showSuccessToast('密码已修改');
        setOldPass('');
        setNewPass('');
        setConfirmPass('');
      } else {
        showWarningToast(res.msg);
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <Box maxW="400px" mx="auto" px={5} pt="80px" pb="120px">
      <Text fontSize="28px" fontWeight={600} mb="48px">个人中心</Text>

      {/* 用户信息 */}
      {user && (
        <Flex align="center" gap={4} mb="48px">
          <Flex
            w="40px"
            h="40px"
            borderRadius="50%"
            bg="#111"
            color="#fff"
            align="center"
            justify="center"
            fontSize="14px"
            fontWeight={500}
            flexShrink={0}
          >
            {user.name.charAt(0).toUpperCase()}
          </Flex>
          <Box>
            <Text fontSize="16px" fontWeight={500}>{user.name}</Text>
            <Text fontSize="13px" color="#999">UID: {user.id}</Text>
          </Box>
        </Flex>
      )}

      <Box h="1px" bg="#eee" mb="48px" />

      {/* 修改密码 */}
      <Text fontSize="16px" fontWeight={500} mb="24px">修改密码</Text>

      <form onSubmit={handleChangePassword}>
        <Box mb={5}>
          <Text fontSize="13px" color="#666" mb="6px">当前密码</Text>
          <input
            style={inputStyle}
            type="password"
            value={oldPass}
            onChange={(e) => setOldPass(e.target.value)}
            autoComplete="current-password"
            onFocus={(e) => { e.target.style.borderColor = '#111'; }}
            onBlur={(e) => { e.target.style.borderColor = '#ddd'; }}
          />
        </Box>
        <Box mb={5}>
          <Text fontSize="13px" color="#666" mb="6px">新密码</Text>
          <input
            style={inputStyle}
            type="password"
            value={newPass}
            onChange={(e) => setNewPass(e.target.value)}
            placeholder="至少 6 个字符"
            autoComplete="new-password"
            onFocus={(e) => { e.target.style.borderColor = '#111'; }}
            onBlur={(e) => { e.target.style.borderColor = '#ddd'; }}
          />
        </Box>
        <Box mb={8}>
          <Text fontSize="13px" color="#666" mb="6px">确认新密码</Text>
          <input
            style={inputStyle}
            type="password"
            value={confirmPass}
            onChange={(e) => setConfirmPass(e.target.value)}
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
          {loading ? '修改中...' : '修改密码'}
        </button>
      </form>
    </Box>
  );
};

export default ProfilePage;
