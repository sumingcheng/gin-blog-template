import { FC } from 'react';
import { Outlet } from 'react-router-dom';
import Header from './header';

const MainLayout: FC = () => {
  return (
    <>
      <Header />
      <Outlet />
    </>
  );
};

export default MainLayout;
