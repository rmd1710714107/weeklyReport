'use client'
import {Button, Card, Modal, Space, Switch} from "antd";
import {HTTPResponse, ProjectItemType} from "@/types";
import {FC} from "react";
import {ExclamationCircleFilled} from "@ant-design/icons";
import request from "@/network";
type Props={
    item:ProjectItemType
    deleteCb:()=>void
}
const {confirm}=Modal
const ProjectItem:FC<Props>=({item,deleteCb})=>{
    const deleteItem=()=>{
        confirm({
            title:`是否确认删除${item.projectName}`,
            icon: <ExclamationCircleFilled />,
            onOk:async ()=>{
                // @ts-expect-error
                const res=await request.delete<HTTPResponse<null>>("/delProjectList",{
                    params:{deleteId:item.id}
                }) as HTTPResponse<null>
                if (res.code===200){
                    deleteCb()
                    return Promise.resolve(true)
                }
                return Promise.reject(false)
            }
        })
    }
    return(
        <Card size="small">
            <div className="flex justify-between">
                <div className="table">
                    <span className="table-cell align-middle">
                        {item && item.projectName}
                    </span>
                </div>

                <Space size="middle">
                    <Switch></Switch>
                    <Button danger type="primary" onClick={deleteItem}>删除</Button>
                </Space>

            </div>

        </Card>
    )
}
export default ProjectItem