import { Button, Checkbox, Select, Form, Input, message } from 'antd';
import { useRef, useState } from 'react';
import Bg from '~/assets/bg.jpg';
import api from '~/api/api';
import { useNavigate } from 'react-router-dom'
import useInfoStore from '~/stores/info'


export default () => {
    const [isLogin, setIsLogin] = useState(true);
    const navigate = useNavigate()
    const onLogin = (values) => {
        api.user.login(values).then(res => {
            navigate("/certificate/verify")
        })
    };
    const setInfo = useInfoStore((state) => state.setInfo)

    const onRegister = (values) => {
        console.log('Success:', values);
        api.user.register(values).then(res => {
            console.log(res);
            message.success('注册成功，请登录');
            setIsLogin(true);
            setInfo(res)
        })
    }
    return (
        <div className='flex justify-center items-center w-full h-screen' style={{
            backgroundImage: `url(${Bg})`,
            backgroundRepeat: 'no-repeat',
            backgroundSize: 'cover',

        }}>
            <div className='bg-white bg-opacity-80 rounded p-10'>
                <h1>基于区块链的证书管理平台</h1>
                {
                    isLogin && <Form
                        name="登录"
                        labelCol={{ span: 4 }}
                        wrapperCol={{ span: 16 }}
                        style={{ maxWidth: 600 }}
                        initialValues={{ remember: true }}
                        onFinish={onLogin}
                        autoComplete="off"
                    >
                        <Form.Item
                            label="用户名"
                            name="username"
                            rules={[{ required: true, message: '请输入用户名!' }]}
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item
                            label="密码"
                            name="password"
                            rules={[{ required: true, message: '请输入密码!' }]}
                        >
                            <Input.Password />
                        </Form.Item>
                        <Form.Item wrapperCol={{ offset: 6, span: 16 }}>
                            <Button type="info" onClick={() => setIsLogin(false)}>
                                注册
                            </Button>
                            <Button type="primary" htmlType="submit">
                                登录
                            </Button>
                        </Form.Item>
                    </Form>
                }
                {
                    !isLogin && <Form
                        name="注册"
                        labelCol={{ span: 4 }}
                        wrapperCol={{ span: 16 }}
                        style={{ maxWidth: 600 }}
                        onFinish={onRegister}
                        autoComplete="off"
                    >
                        <Form.Item
                            label="用户名"
                            name="username"
                            rules={[{ required: true, message: '请输入用户名!', min: 6 }]}
                        >
                            <Input />
                        </Form.Item>
                        <Form.Item
                            label="呢称"
                            name="nickname"
                            rules={[{ required: true, message: '请输入呢称!', min: 6 }]}
                        >
                            <Input />
                        </Form.Item>

                        <Form.Item
                            label="密码"
                            name="password"
                            rules={[{ required: true, message: '请输入密码!', min: 8 }]}
                        >
                            <Input.Password />
                        </Form.Item>
                        <Form.Item
                            label="角色"
                            name="role"
                            rules={[{ required: true, message: '请选择角色!' }]}
                        >
                            <Select
                                options={[
                                    { label: '教师', value: 'teacher' },
                                    { label: '学生', value: 'student' },
                                ]}
                            />
                        </Form.Item>
                        <Form.Item wrapperCol={{ offset: 6, span: 16 }}>
                            <Button type="info" onClick={() => setIsLogin(true)}>
                                登录
                            </Button>
                            <Button type="primary" htmlType="submit">
                                注册
                            </Button>
                        </Form.Item>
                    </Form>
                }
            </div>



        </div>

    );
};