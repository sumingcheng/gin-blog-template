import { Route, Routes } from 'react-router-dom';
import HeaderLayout from '../components/headerLayout';
import Home from '../pages/home';
import Login from '../pages/login';
import RegisterPage from '../pages/register';
import BlogDetailPage from '../pages/blog/detail';
import EditBlogPage from '../pages/blog/edit';
import CreateBlogPage from '../pages/blog/create';
import ProfilePage from '../pages/profile';

const AppRoutes = () => {
  return (
    <Routes>
      <Route element={<HeaderLayout />}>
        <Route path="/" element={<Home />} />
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<RegisterPage />} />
        <Route path="/blog/:bid" element={<BlogDetailPage />} />
        <Route path="/blog/edit/:bid" element={<EditBlogPage />} />
        <Route path="/blog/create" element={<CreateBlogPage />} />
        <Route path="/profile" element={<ProfilePage />} />
      </Route>
    </Routes>
  );
};

export default AppRoutes;
