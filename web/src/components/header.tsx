import { FC } from "react";
import { Tab, TabList, Tabs } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

import useCustomToast from "../hooks/useCustomToast.tsx";
import { getBelong } from "../api/blog.ts";

const Header: FC = () => {
  const navigate = useNavigate();
  const { showWarningToast } = useCustomToast();
  const handleNavigation = async (path: string) => {
    if (path === '/') {
      const res = await getBelong({
        "bid": 1,
      });
      if (res.belong) {
        navigate(path);
      } else {
        showWarningToast('请先登录！');
      }
    } else {
      navigate(path);
    }
  };

  return (
    <header>
      <Tabs align='end' variant='enclosed'>
        <TabList>
          <Tab onClick={ () => handleNavigation('/login') }>LOGIN</Tab>
          <Tab onClick={ () => handleNavigation('/blog') }>BLOG</Tab>
          <Tab onClick={ () => handleNavigation('/') }>HOME</Tab>
        </TabList>
      </Tabs>
    </header>
  );
}

export default Header;
