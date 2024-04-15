import { Button, Form, message, Card, Input, Badge, Descriptions } from 'antd';
const { Search } = Input;
import { useEffect, useState } from 'react';
import api from '~/api/api';
import moment from 'moment'
// ### 证书基本信息
// - id
// - 学生姓名
// - 成绩
// - 课程
// - 学生地址
// - 申请时间
// - 审核时间
// - 证书状态 (待审核 审核通过 审核拒绝)
// - 审核人
// - 审核时间

export default () => {
    const [items, setItems] = useState([])
    const [show, setShow] = useState(false)


    return (
        <Card title={<h1 className='text-center'>证书检验</h1>} bordered={false} className='w-4/5 bg-white h-full rounded' style={{
            margin: '0 auto',
            paddingTop: '20px',
        }}>
            <div className='w-full flex justify-center flex-col items-center'>
                <Search placeholder="请输入证书编码" allowClear className='w-100' onSearch={(val) => {
                    api.certificate.id({
                        id: Number(val)
                    }).then(res => {
                        console.log(res)
                        if(res.Id == 0){
                            message.error('未查询到该证书')
                            setShow(false)
                            return
                        }
                        let status = ""
                        let text = ""
                        if (res.Status == "Pending") {
                            status = "processing"
                            text = "待审核"
                        } else if (res.Status == "Approved") {
                            status = "success"
                            text = "审核通过"
                        } else {
                            status = "error"
                            text = "审核拒绝"
                        }
                        setItems([
                            {
                                key: 'id',
                                label: '证书编号',
                                children: res.Id,
                            },
                            {
                                key: 'studentName',
                                label: '学生姓名',
                                children: res.StudentName,
                            },
                            {
                                key: 'grade',
                                label: '成绩',
                                children: res.Grade,
                            },
                            {
                                key: 'course',
                                label: '课程',
                                children: res.Course,
                            },
                            {
                                key: 'studentAddress',
                                label: '学生地址',
                                children: res.StudentAddress,
                            },
                            {
                                key: 'applyTime',
                                label: '申请时间',
                                children: moment(res.ApplyTime).format('YYYY-MM-DD HH:mm:ss'),
                            },
                            {
                                key: 'status',
                                label: '证书状态',
                                children: <Badge status={status} text={text} />,

                            },
                            {
                                key: 'auditor',
                                label: '审核人',
                                children: res.Approver == '0x0000000000000000000000000000000000000000' ? '无' : res.Approver,
                            },
                            {
                                key: 'auditTime',
                                label: '审核时间',
                                children: moment(res.ApprovalTime).format('YYYY-MM-DD HH:mm:ss'),
                            },

                        ])
                        setShow(true)
                    })
                }} />
                {
                    show &&  <Descriptions  title={<h2 className='text-center'>证书信息</h2>} className='mt-5' bordered items={items} />
                }
            </div>
        </Card>

    );
};