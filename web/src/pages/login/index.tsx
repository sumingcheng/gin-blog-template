import { FC, FormEvent, useEffect, useState } from 'react';
import { Box, Button, FormControl, FormLabel, Heading, Input, Text, VStack } from '@chakra-ui/react';
import useCustomToast from "../../hooks/useCustomToast.tsx";
import { login } from "../../api/user.ts";
import { encryptPassword } from "../../utils/md5.ts";
import { useNavigate } from "react-router-dom";

const LoginPage: FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [authToken, setAuthToken] = useState('');
  const { showWarningToast, showSuccessToast } = useCustomToast();
  const navigate = useNavigate();

  useEffect(() => {
    setAuthToken(sessionStorage.getItem('auth_token') || '');
  }, []);

  const handleLogin = async (event: FormEvent) => {
    event.preventDefault();
    if (!username || !password) {
      showWarningToast('用户名和密码是必填项');
      return;
    }
    const res = await login({ user: username, pass: encryptPassword(password) });
    if (res.code !== 0) {
      showWarningToast(res.msg);
      return;
    }

    sessionStorage.setItem('auth_token', res.auth_token)
    showSuccessToast('登录成功');
    navigate('/blog');
  };

  if (authToken) {
    return (
      <Box
        minH="full"
        width="full"
        display="flex"
        alignItems="center"
        justifyContent="center"
        marginTop={ "5%" }
      >
        <VStack
          spacing={ 4 }
          w="full"
          maxW="md"
          rounded="md"
          boxShadow="xs"
          p={ 8 }
        >
          <Text fontSize="xl">您已登录</Text>
        </VStack>
      </Box>
    );
  }

  return (
    <Box
      minH="full"
      width="full"
      display="flex"
      alignItems="center"
      justifyContent="center"
      marginTop={ "5%" }
    >
      <VStack
        spacing={ 4 }
        w="full"
        maxW="md"
        rounded="md"
        boxShadow="xs"
        p={ 8 }
        as="form"
        onSubmit={ handleLogin }
      >
        <Heading size="lg" textAlign="center">
          登录
        </Heading>
        <FormControl id="username" isRequired>
          <FormLabel>用户名</FormLabel>
          <Input
            type="text"
            value={ username }
            onChange={ (e) => setUsername(e.target.value) }
            placeholder="用户名:admin"
            autoComplete="username"
          />
        </FormControl>
        <FormControl id="password" isRequired mt={ 6 }>
          <FormLabel>密码</FormLabel>
          <Input
            type="password"
            value={ password }
            onChange={ (e) => setPassword(e.target.value) }
            placeholder="密码:123456"
            autoComplete="current-password"
          />
        </FormControl>
        <Button
          type="submit"
          colorScheme="blue"
          size="lg"
          fontSize="md"
          mt={ 8 }
          w="full"
        >
          确认
        </Button>
      </VStack>
    </Box>
  );
};

export default LoginPage;
