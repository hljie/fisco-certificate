import { EllipsisOutlined, PlusOutlined } from '@ant-design/icons';
import { ProTable, TableDropdown } from '@ant-design/pro-components';
import { Button, Dropdown, Space, Tag, message } from 'antd';
import { useRef } from 'react';
import api from '~/api/api'

const columns = [
  {
    dataIndex: 'index',
    valueType: 'indexBorder',
    width: 48,
  },
  {
    title: '账号',
    dataIndex: 'Username',
    editable: false,
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
    title: '呢称',
    dataIndex: 'Nickname',
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
    title: '角色',
    dataIndex: 'Role',
    valueType: 'select',
    valueEnum: {
      "admin": "管理员",
      "teacher": "教师",
      "user": "学生",
    },
  },
  {
    disable: true,
    title: '状态', //（待审核 启用 禁用）
    dataIndex: 'Status',
    filters: true,
    onFilter: true,
    ellipsis: true,
    valueType: 'select',
    valueEnum: {
      'pending': { text: '待审核', status: 'Default' },
      'enabled': { text: '启用', status: 'Success' },
      'disabled': { text: '禁用', status: 'Error' },
    },
  },
  {
    title: '注册时间',
    key: 'showTime',
    editable: false,
    dataIndex: 'CreatedAt',
    valueType: 'date',

    hideInSearch: true,
  },
  {
    title: '操作',
    valueType: 'option',
    key: 'option',
    render: (text, record, _, action) => [
      <a
        key="editable"
        onClick={() => {
          action?.startEditable?.(record.ID);
        }}
      >
        编辑
      </a>,
    ],
  },
];

export default () => {
  const actionRef = useRef();
  return (
    <ProTable
      columns={columns}
      actionRef={actionRef}
      cardBordered
      request={async (params, sort, filter) => {
        console.log(params, sort, filter);
        try {
          const res = await api.user.list({
            page: params.current,
            pageSize: params.pageSize,
            ...params,
          });
          return {
            data: res.data,
            success: true,
            total: res.total,
          }
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
          try{
            await api.user.update(data)
            message.success('保存成功')
            return true
          }catch(e){
            message.error('保存失败')
            return false
          }
          
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
      rowKey="ID"
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
      headerTitle="用户列表"
    />
  );
};