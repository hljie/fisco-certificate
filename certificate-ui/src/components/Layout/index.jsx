// src/components/Layout/index.tsx
import { useState, useEffect } from 'react'
import { Outlet, useNavigate, useLocation } from 'react-router-dom'
import { Layout, Menu, Button } from 'antd'
import Logo from '~/assets/logo.svg'
import api from '~/api/api';
import useInfoStore from '~/stores/info'

const { Header, Sider, Content } = Layout



const BasicLayout = () => {
    const navigate = useNavigate()
    const { pathname } = useLocation()

    const info = useInfoStore((state) => state.info)
    const setInfo = useInfoStore((state) => state.setInfo)

    const [items, setItems] = useState([])

    useEffect(() => {


        api.user.info().then(res => {
            setInfo(res)
            if (res.role == 'admin') {
                setItems([
                    {
                        path: '/certificate/verify',
                        label: '证书检验',
                    },
                    {
                        path: '/info',
                        label: '个人信息',
                    },
                    {
                        path: '/user',
                        label: '用户管理',
                    },
                ].map((nav) => ({
                    key: nav.path,
                    icon: null,
                    label: nav.label,
                })))
            } else if (res.role == 'teacher') {
                setItems([
                    {
                        path: '/certificate/verify',
                        label: '证书检验',
                    },
                    {
                        path: '/info',
                        label: '个人信息',
                    },
                    {
                        label: '证书管理',
                        path: '/certificate',
                    },
                ].map((nav) => ({
                    key: nav.path,
                    icon: null,
                    label: nav.label,
                })))
            } else {
                setItems([
                    {
                        path: '/certificate/verify',
                        label: '证书检验',
                    },
                    {
                        path: '/info',
                        label: '个人信息',
                    },
                    {
                        path: '/certificate/me',
                        label: '我的证书',
                    },
                ].map((nav) => ({
                    key: nav.path,
                    icon: null,
                    label: nav.label,
                })))
            }
        })
    }, [])

    const handleMenuClick = ({ key }) => {
        navigate(key)
    }

    return (
        <Layout style={{ height: '100vh' }}>
            <Sider trigger={null} collapsible theme="light" className='shadow' style={{
                backgroundColor: "#f9f9f9"
            }}>
                <div className='flex items-center justify-center'
                >
                    <img src={Logo} height={50} />
                </div>
                <Menu
                    mode="inline"
                    defaultSelectedKeys={[pathname]}
                    items={items}
                    style={{
                        backgroundColor: "#f9f9f9"
                    }}
                    onClick={handleMenuClick}
                />
            </Sider>
            <Layout style={{ display: 'flex', flexDirection: 'column' }}>
                <Header style={{ background: '#fff', padding: 0 }}>
                    <div className='w-full flex justify-between px-5'>
                        <h1>基于区块链的证书管理平台</h1>
                        <div className='flex items-center x-space-2'>
                            {info.nickname}
                            <Button type="text"
                                onClick={() => {
                                    // todo: 退出登录
                                    navigate('/login')
                                    api.user.logout()
                                }}
                            >退出登录</Button>
                        </div>
                    </div>
                </Header>
                <Content style={{ padding: '16px', flex: 1, overflowY: 'auto' }}>
                    <Outlet />
                </Content>
            </Layout>
        </Layout>
    )
}

export default BasicLayout
