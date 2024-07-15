import { FC } from "react";
import { Button, Container, Text } from "@chakra-ui/react";
import { useNavigate } from "react-router-dom";
import useCustomToast from "../../hooks/useCustomToast.tsx";

const Home: FC = () => {
  const { showSuccessToast } = useCustomToast();
  const navigate = useNavigate();

  const logout = () => {
    navigate('/login');
    showSuccessToast('已注销');
    sessionStorage.removeItem('auth_token');
  }

  return (
    <Container centerContent>
      <Text fontSize="6xl" fontWeight="bold" color="teal.500" fontFamily="Arial" display="flex" justifyContent="center"
            flexDirection="column" alignItems="center">
        <Text>Go-Blog</Text>
        <Text fontSize="xl">登录后可查看此页面</Text>
        <Button mt={ 10 } onClick={ logout }>退出登录</Button>
      </Text>
    </Container>
  );
};

export default Home;
