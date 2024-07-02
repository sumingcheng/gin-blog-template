import { FC } from "react";
import { Tab, TabList, Tabs } from '@chakra-ui/react';
import { useNavigate } from 'react-router-dom';

const Header: FC = () => {
  const navigate = useNavigate();

  // 处理点击事件，根据 tab 的索引来决定路由路径
  const handleTabsChange = (index: number) => {
    switch (index) {
      case 0:
        navigate('/blog');
        break;
      case 1:
        navigate('/home');
        break;
      case 2:
        navigate('/login');
        break;
      default:
        break;
    }
  };
  return (
    <header>
      <Tabs align='end' variant='enclosed' onChange={ handleTabsChange }>
        <TabList>
          <Tab>BLOG</Tab>
          <Tab>HOME</Tab>
          <Tab>LOGIN</Tab>
        </TabList>
      </Tabs>
    </header>
  );
}
export default Header
