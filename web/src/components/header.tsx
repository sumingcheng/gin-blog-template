import { FC } from "react";
import { Tab, TabList, Tabs } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

const Header: FC = () => {
  const navigate = useNavigate();

  return (
    <header>
      <Tabs align='end' variant='enclosed'>
        <TabList>
          <Tab onClick={ () => navigate('/login') }>LOGIN</Tab>
          <Tab onClick={ () => navigate('/blog') }>BLOG</Tab>
          <Tab onClick={ () => navigate('/') }>HOME</Tab>
        </TabList>
      </Tabs>
    </header>
  );
}
export default Header;
