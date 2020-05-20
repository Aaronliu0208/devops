import React, { Component } from 'react'
import { Layout } from 'antd'

const { Content, Sider } = Layout;

class IndexLayout extends Component {

    render() {
        return (
            <Layout>
                <Sider
                    collapsible
                    trigger={null}
                >
                    <div>sider</div>
                </Sider>
                <Layout>
                    <Content style={{background: '#f7f7f7'}}>
                        <div>content</div>
                    </Content>
                </Layout>
            </Layout>
        )
    }
}

export default IndexLayout
