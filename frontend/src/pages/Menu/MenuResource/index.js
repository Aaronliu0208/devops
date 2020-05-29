import React, { PureComponent } from 'react';
import { Table, Button, Popconfirm } from 'antd';
import { newUUID } from '@/utils/utils';
import { EditableCell, EditableFormRow } from './EditableCell';

import styles from './index.less';

function fillKey(data) {
  if (!data) {
    return [];
  }
  return data.map(item => {
    const nitem = { ...item };
    if (!nitem.key) {
      nitem.key = newUUID();
    }
    return nitem;
  });
}

export default class MenuAction extends PureComponent {
  constructor(props) {
    super(props);

    this.columns = [
      {
        title: '请求方式',
        dataIndex: 'method',
        editable: true,
        width: '30%',
      },
      {
        title: '请求路径',
        dataIndex: 'path',
        editable: true,
        width: '45%',
      },
      {
        title: '操作',
        dataIndex: 'key',
        width: '25%',
        render: (_, record) => {
          const { dataSource } = this.state;
          if (dataSource.length === 0) {
            return null;
          }
          return (
            <Popconfirm title="确定要删除该数据吗?" onConfirm={() => this.handleDelete(record.key)}>
              <a>删除</a>
            </Popconfirm>
          );
        },
      },
    ];

    this.state = {
      dataSource: fillKey(props.value),
      addVisible: false,
    };
  }

  static getDerivedStateFromProps(nextProps, state) {
    if ('value' in nextProps) {
      return { ...state, dataSource: fillKey(nextProps.value) };
    }
    return state;
  }

  handleAddCancel = () => {
    this.setState({ addVisible: false });
  };

  handleDelete = key => {
    const { dataSource } = this.state;
    const data = dataSource.filter(item => item.key !== key);
    this.setState({ dataSource: data }, () => {
      this.triggerChange(data);
    });
  };

  handleAdd = () => {
    const { dataSource } = this.state;
    const item = {
      key: newUUID(),
      method: '',
      path: '',
    };

    const data = [...dataSource, item];
    this.setState(
      {
        dataSource: data,
      },
      () => {
        this.triggerChange(data);
      }
    );
  };

  handleSave = row => {
    const { dataSource } = this.state;
    const data = [...dataSource];
    const index = data.findIndex(item => row.key === item.key);
    const item = data[index];
    data.splice(index, 1, {
      ...item,
      ...row,
    });
    this.setState({ dataSource: data }, () => {
      this.triggerChange(data);
    });
  };

  triggerChange = data => {
    const { onChange } = this.props;
    if (onChange) {
      onChange(data);
    }
  };

  render() {
    const { dataSource } = this.state;
    const components = {
      body: {
        row: EditableFormRow,
        cell: EditableCell,
      },
    };
    const columns = this.columns.map(col => {
      if (!col.editable) {
        return col;
      }
      return {
        ...col,
        onCell: record => ({
          record,
          editable: col.editable,
          dataIndex: col.dataIndex,
          title: col.title,
          handleSave: this.handleSave,
        }),
      };
    });
    return (
      <div className={styles.tableList}>
        <div className={styles.tableListOperator}>
          <Button onClick={this.handleAdd} size="small" type="primary">
            新增
          </Button>
        </div>
        <Table
          rowKey={record => record.key}
          components={components}
          bordered
          dataSource={dataSource}
          columns={columns}
          pagination={false}
        />
      </div>
    );
  }
}
