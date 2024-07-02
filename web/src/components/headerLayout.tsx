import React from 'react';
import { Outlet } from 'react-router-dom';
import Header from './header'; // 确保导入你的Header组件

const MainLayout: React.FC = () => {
  return (
    <>
      <Header/>
      <Outlet/>
    </>
  );
};

export default MainLayout;
