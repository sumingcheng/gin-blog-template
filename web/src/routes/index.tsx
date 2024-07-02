import { Route, Routes } from 'react-router-dom';
import Home from '../pages/home/index.tsx';
import Login from '../pages/login/index.tsx';
import Blog from '../pages/blog/index.tsx'
import HeaderLayout from '../components/headerLayout.tsx';


const AppRoutes = () => {
  return (
    <Routes>
      <Route path="/" element={ <Home/> }/>
      <Route element={ <HeaderLayout/> }>
        <Route path="/login" element={ <Login/> }/>
        <Route path="/blog" element={ <Blog/> }/>
      </Route>
    </Routes>
  );
};

export default AppRoutes;
