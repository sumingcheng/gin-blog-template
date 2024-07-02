import { FC } from "react";
import { Container, Text } from "@chakra-ui/react";

const Home: FC = () => {
  return (
    <Container centerContent>
      <Text fontSize="6xl" fontWeight="bold" color="teal.500" fontFamily="Arial">
        Go-Blog
      </Text>
    </Container>
  );
};

export default Home;
