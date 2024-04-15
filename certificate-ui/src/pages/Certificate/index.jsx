import { EllipsisOutlined, PlusOutlined } from '@ant-design/icons';
import { ProTable, TableDropdown } from '@ant-design/pro-components';
import { Button, Dropdown, Space, Tag } from 'antd';
import { useRef } from 'react';
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
    dataIndex: 'index',
    valueType: 'indexBorder',
    width: 48,
  },
  {
    title: '学生姓名',
    dataIndex: 'StudentName',
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
    title: '成绩',
    dataIndex: 'Grade',
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
    title: '课程',
    editable: false,
    dataIndex: 'Course',
  },
  {
    title: '学生地址',
    editable: false,
    dataIndex: 'StudentAddress',
    valueType: 'textarea',
  },
  {
    title: '申请时间',
    key: 'applyTime',
    editable: false,
    dataIndex: 'ApplicationTime',
    valueType: 'date',
    sorter: true,
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
  return (
    <ProTable
      columns={columns}
      actionRef={actionRef}
      cardBordered
      request={async (params, sort, filter) => {
        try {
          let res = await api.certificate.list(params)
          if (res != null) {
            res = res
              .filter(item => item.Id != 0);
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
          api.certificate.approve({
            id: data.Id,
            grade: Number(data.Grade),
            status: data.Status
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
      headerTitle="证书列表"
    />
  );
};