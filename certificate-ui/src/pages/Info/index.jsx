
import {
  ModalForm,
  ProForm,
  ProFormDateRangePicker,
  ProFormSelect,
  ProFormText,

} from '@ant-design/pro-components';
import { useEffect } from 'react';
import { Button, Form, message, Card, Input } from 'antd';
import useInfoStore from '~/stores/info'
import api from '~/api/api'

export default () => {
  const info = useInfoStore((state) => state.info)
  const setInfo = useInfoStore((state) => state.setInfo)

  useEffect(() => {
    api.user.info().then(res => {
      setInfo(res)
    })
  }, [])


  const [form] = Form.useForm();

  const onFinish = (val) => {
    console.log("finish", val);
    api.user.updateMe(val).then(res => {
      console.log(res);
      message.success('修改成功');
      setInfo(res)
    })
  }
  return (
    <Card title={<h1 className='text-center'>个人信息</h1>} bordered={false} className='w-3/5 bg-white h-full rounded' style={{
      margin: '0 auto',
      paddingTop: '20px',
    }}>
      <div className='w-full flex justify-center flex-col items-center'>
        <Form
          labelCol={{ span: 4 }}
          wrapperCol={{ span: 14 }}
          layout="horizontal"
          form={form}
          initialValues={
            {
              nickname: info.nickname
            }
          }
          style={{ maxWidth: 600, minWidth: 450 }}
          onFinish={onFinish}
        >
          <Form.Item
            label="呢称"
            name="nickname"
            rules={[{ required: true, message: '请输入呢称!' }]}
          >
            <Input />
          </Form.Item>
        </Form>
        <Button type="primary" className='w-40 mb-5' onClick={() => {form.submit()}}>
          保存
        </Button>
      </div>

      <div className='w-full flex justify-center'>
        <ModalForm
          title="修改密码"
          trigger={
            <Button type="primary" className='w-40'>
              修改密码
            </Button>
          }
          autoFocusFirstInput
          modalProps={{
            destroyOnClose: true,
            onCancel: () => console.log('run'),
          }}
          width={350}
          submitTimeout={2000}
          onFinish={async (values) => {
            api.user.updateMe(values).then(res => {
              console.log(res);
              message.success('修改成功');
              setInfo(res)
            })
            return true;
          }}
        >
          <ProForm.Group>
            <ProFormText.Password
              width="md"
              
              name="oldPassword"
              required
              rules={[{ required: true, message: '请输入新密码', min: 8 }]}
              label="旧密码"
              placeholder="请输入旧密码"
            />
          </ProForm.Group>
          <ProForm.Group>
            <ProFormText.Password
              width="md"
              name="password"
              required
              rules={[{ required: true, message: '请输入新密码', min: 8 }]}
              label="新密码"
              placeholder="请输入新密码"
            />
          </ProForm.Group>
        </ModalForm>
      </div>
    </Card>

  );
};