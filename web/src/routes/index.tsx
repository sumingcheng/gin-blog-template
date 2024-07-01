import { Route, Routes } from 'react-router-dom';
import Home from '../pages/home/index.tsx';
import Login from '../pages/login/index.tsx';
import Blog from '../pages/blog/index.tsx'

const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={ <Home/> }/>
      <Route path="/login" element={ <Login/> }/>
      <Route path="/blog" element={ <Blog/> }/>
    </Routes>
  );
};

export default AppRoutes;
