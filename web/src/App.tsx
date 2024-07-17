import { BrowserRouter } from "react-router-dom";
import AppRoutes from './routes/index.tsx';
import { useEffect } from "react";
import { getAuthToken } from "./api/user.ts";
import useCustomToast from "./hooks/useCustomToast.tsx";

function App() {
  const { showWarningToast } = useCustomToast();

  useEffect(() => {
    if (sessionStorage.getItem('auth_token')) return
    getAuthToken().then(res => {
      if (res.code === 0) {
        sessionStorage.setItem('auth_token', res.auth_token)
      } else if (res.code === 1) {
        // showWarningToast(res.msg)
      } else {
        // showWarningToast('未知错误')
      }
    })
  }, [showWarningToast])

  return (
    <>
      <BrowserRouter>
        <AppRoutes/>
      </BrowserRouter>
    </>
  )
}

export default App
