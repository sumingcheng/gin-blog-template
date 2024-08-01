import { Route, Routes, useNavigate } from 'react-router-dom';
import Home from '../pages/home/index.tsx';
import Login from '../pages/login/index.tsx';
import Blog from '../pages/blog/index.tsx'
import HeaderLayout from '../components/headerLayout.tsx';
import Edit from "../pages/blog/edit.tsx";
import { useEffect } from "react";

const AppRoutes = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const token = sessionStorage.getItem('auth_token');
    if (!token) {
      navigate('/login');
    }
  }, [navigate]);

  return (
    <Routes>
      <Route element={ <HeaderLayout/> }>
        <Route path="/" element={ <Home/> }/>
        <Route path="/login" element={ <Login/> }/>
        <Route path="/blog" element={ <Blog/> }/>
        <Route path="/edit" element={ <Edit/> }/>
      </Route>
    </Routes>
  );
};

export default AppRoutes;
