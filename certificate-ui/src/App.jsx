// App.tsx
import { RouterProvider } from 'react-router-dom'
import { ConfigProvider } from 'antd'
import router from './router'
import dayjs from 'dayjs'
import zhCN from 'antd/locale/zh_CN'
import 'dayjs/locale/zh-cn'
import 'antd/dist/reset.css'
// import { useNavigate,  } from 'react-router-dom'
// import { AxiosInterceptorsSetup } from '~/api/config'
// function AxiosInterceptorNavigate() {
//   let navigate = useNavigate();
//   AxiosInterceptorsSetup(navigate);
//   return <></>;
// }


dayjs.locale('zh-cn')

const App = () => {


  return (
    <ConfigProvider
      locale={zhCN}
      theme={{
        token: {
          colorPrimary: '#00b96b',
        },
      }}
    >
      <RouterProvider router={router} >
      

      </RouterProvider>
    </ConfigProvider>
  )
}

export default App
