import { createBrowserRouter, Navigate } from 'react-router-dom'
import User from '~/pages/User'
import Certificate from '~/pages/Certificate'
import MeCertificate from '~/pages/Certificate/me'
import VerifyCertificate from '~/pages/Certificate/verify'
import Info from '~/pages/Info'
import Layout from '~/components/Layout'
import Login from '~/pages/Login'


const routes = [
    {
        path: '/login',
        element: <Login />,
    },
    {
        path: '/',
        element: <Layout />,
        children: [
            {
                index: true,
                element: <Navigate to="/login" replace />,
            },
            {
                path: 'info',
                element: <Info />,
            },
            {
                path: 'user',
                element: <User />,
            },
            {
                path: 'certificate',
                element: <Certificate />,
            },
            {
                path: 'certificate/me',
                element: <MeCertificate />,
            },
            {
                path: 'certificate/verify',
                element: <VerifyCertificate />,
            },
        ],
    },
]

export default createBrowserRouter(routes, {
    basename: '/',
})
