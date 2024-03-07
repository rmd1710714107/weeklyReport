import './App.css'
import {Layout} from "antd";
import {Content} from "antd/es/layout/layout";
import {Outlet} from "react-router-dom";

function App() {
    return (
        <>
            <Layout className="pt-4 h-full w-6/12 mx-auto">
                <Content className="pb-5 flex w-full px-5 flex-col">
                    <Outlet/>
                </Content>
            </Layout>
        </>
    )
}

export default App
