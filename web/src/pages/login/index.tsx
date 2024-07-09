import { FC, FormEvent, useState } from 'react';
import { Box, Button, FormControl, FormLabel, Heading, Input, VStack } from '@chakra-ui/react';
import useCustomToast from "../../hooks/useCustomToast.tsx";
import { login } from "../../api/user.ts";
import { encryptPassword } from "../../utils/md5.ts";

const LoginPage: FC = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const { showWarningToast, showSuccessToast } = useCustomToast();

  const handleLogin = async (event: FormEvent) => {
    event.preventDefault();
    if (!username || !password) {
      showWarningToast('用户名和密码是必填项');
      return;
    }
    const res = await login({ user: username, pass: encryptPassword(password) });
    console.log(res)
    if (res.code !== 0) {
      showWarningToast(res.message);
      return;
    }

    sessionStorage.setItem('auth_token', res.auth_token)
    showSuccessToast('登录成功');
  };

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
            placeholder="请输入您的用户名"
            autoComplete="username"
          />
        </FormControl>
        <FormControl id="password" isRequired mt={ 6 }>
          <FormLabel>密码</FormLabel>
          <Input
            type="password"
            value={ password }
            onChange={ (e) => setPassword(e.target.value) }
            placeholder="请输入您的密码"
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
