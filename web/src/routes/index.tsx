import { Route, Routes } from 'react-router-dom';
import Home from '../pages/home/index.tsx';
import Login from '../pages/login/index.tsx';
import Blog from '../pages/blog/index.tsx'
import HeaderLayout from '../components/headerLayout.tsx';
import Edit from "../pages/blog/edit.tsx";

const AppRoutes = () => {
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
