import { EllipsisOutlined, PlusOutlined } from '@ant-design/icons';
import { ProTable, TableDropdown } from '@ant-design/pro-components';
import { Button, Dropdown, Space, Tag, Modal, Form, Input } from 'antd';
import { useRef, useState } from 'react';
import api from '~/api/api';

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
const columns = [
    {
        title: '证书编号',
        dataIndex: 'Id',
        hideInSearch: true,
        editable: false,
        width: 100,
    },
    {
        title: '学生姓名',
        dataIndex: 'StudentName',
        hideInSearch: true,
        copyable: true,
        formItemProps: {
            rules: [
                {
                    required: true,
                    message: '此项为必填项',
                },
            ],
        },
    },
    {
        title: '成绩',
        dataIndex: 'Grade',
        hideInSearch: true,
        copyable: true,
        editable: false,
        formItemProps: {
            rules: [
                {
                    required: true,
                    message: '此项为必填项',
                },
            ],
        },
    },
    {
        title: '课程',

        dataIndex: 'Course',
    },
    {
        title: '学生地址',
        editable: false,
        width: 350,
        hideInSearch: true,
        dataIndex: 'StudentAddress',
        valueType: 'textarea',
    },
    {
        title: '申请时间',
        key: 'ApplicationTime',
        editable: false,
        dataIndex: 'ApplicationTime',
        valueType: 'date',
        sorter: true,
        hideInSearch: true,
    },
    {
        title: '审核人',
        editable: false,
        width: 350,
        dataIndex: 'Approver',
        hideInSearch: true,
    },
    {
        title: '审核时间',
        key: 'auditTime',
        editable: false,
        dataIndex: 'ApprovalTime',
        valueType: 'date',
        sorter: true,
        hideInSearch: true,
    },
    {
        disable: true,
        title: '状态', //（待审核 审核通过 审核拒绝）
        dataIndex: 'Status',
        filters: true,
        onFilter: true,
        editable: false,
        ellipsis: true,
        valueType: 'select',
        valueEnum: {
            'Pending': { text: '待审核', status: 'Default' },
            'Approved': { text: '审核通过', status: 'Success' },
            'Rejected': { text: '审核拒绝', status: 'Error' },
        },
    },
    {
        title: '操作',
        valueType: 'option',
        key: 'option',
        render: (text, record, _, action) => {
            if (record.Status !== 'Pending') {
                return [
                ];
            }
            return [
                <a
                    key="editable"
                    onClick={() => {
                        action?.startEditable?.(record.Id);
                    }}
                >
                    编辑
                </a>,
            ]
        },
    },
];

export default () => {
    const actionRef = useRef();
    const [isModalOpen, setIsModalOpen] = useState(false);

    const showModal = () => {
        setIsModalOpen(true);
    };

    const handleOk = () => {
        setIsModalOpen(false);
        formCtrl.submit();
    };

    const handleCancel = () => {
        setIsModalOpen(false);
    };

    const [formCtrl] = Form.useForm()

    const onFinish = (val) => {
        console.log("finish", val);
        api.certificate.apply(val).then(res => {
            console.log(res);
            formCtrl.resetFields()
            setIsModalOpen(false);
            actionRef.current.reload();
        })
    }

    return (
        <>
            <ProTable
                columns={columns}
                actionRef={actionRef}
                cardBordered
                request={async (params, sort, filter) => {
                    try {
                        let res = await api.certificate.me(params);
                        if (res != null) {
                            res = res
                            .filter(item => item.Id != 0)
                            .map((item) => {
                                if (item.Approver === '0x0000000000000000000000000000000000000000') {
                                    item.Approver = '';
                                }
                                return item;
                            })
                        }
                        return {
                            data: res,
                            success: true
                        };

                    } catch (error) {
                        return {
                            data: [],
                            success: false,
                        }
                    }
                }}
                editable={{
                    type: 'multiple',
                    actionRender: (row, config, defaultDom) => [
                        defaultDom.save,
                        defaultDom.cancel,
                    ],
                    onSave: async (rowKey, data, row) => {
                        api.certificate.update({
                            id: data.Id,
                            studentName: data.StudentName,
                            course: data.Course,
                          }).then(res => {
                            actionRef.current.reload();
                          })
                    },
             
                }}
                columnsState={{
                    persistenceKey: 'pro-table-singe-demos',
                    persistenceType: 'localStorage',
                    defaultValue: {
                        option: { fixed: 'right', disable: true },
                    },
                    onChange(value) {
                        console.log('value: ', value);
                    },
                }}
                rowKey="Id"
                search={{
                    labelWidth: 'auto',
                }}
                options={{
                    setting: {
                        listsHeight: 400,
                    },
                }}
                form={{
                    // 由于配置了 transform，提交的参与与定义的不同这里需要转化一下
                    syncToUrl: (values, type) => {
                        if (type === 'get') {
                            return {
                                ...values,
                                created_at: [values.startTime, values.endTime],
                            };
                        }
                        return values;
                    },
                }}
                pagination={{
                    pageSize: 5,
                    onChange: (page) => console.log(page),
                }}
                dateFormatter="string"
                headerTitle="我的证书"
                toolBarRender={() => [
                    <Button
                        key="button"
                        icon={<PlusOutlined />}
                        onClick={() => {
                            showModal();
                        }}
                        type="primary"
                    >
                        申请
                    </Button>,
                ]}
            />
            <Modal title="申请证书" open={isModalOpen} onOk={handleOk} onCancel={handleCancel}>
                <Form
                    labelCol={{ span: 4 }}
                    wrapperCol={{ span: 14 }}
                    layout="horizontal"
                    style={{ maxWidth: 600 }}
                    onFinish={onFinish}
                    form={formCtrl}
                >
                    <Form.Item
                        label="学生姓名"
                        name="studentName"
                        rules={[{ required: true, message: '请输入学生姓名!' }]}
                    >
                        <Input />
                    </Form.Item>
                    <Form.Item
                        label="课程"
                        name="course"
                        rules={[{ required: true, message: '请输入课程!' }]}

                    > <Input />
                    </Form.Item>
                </Form>
            </Modal>
        </>

    );
};