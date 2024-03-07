import {FC, useEffect, useState} from 'react';
import {Form, Input, message, Modal} from "antd";
import FormItem from "antd/es/form/FormItem";
import request from "@/network";
import {HTTPResponse, ProjectItemType} from "@/types";

type NewProjectModalProps={
    flag:boolean
    closeModalCB:()=>void
}
const NewProjectModal:FC<NewProjectModalProps> = ({flag,closeModalCB}) => {
    const [isModalOpen,setIsOpenModal]=useState(false)
    const [form] = Form.useForm();
    useEffect(()=>{
        if (flag){
            setIsOpenModal(flag)
        }
    },[flag])
    const closeModal=()=>{
        setIsOpenModal(false)
        closeModalCB()
    }
    const [loading,setLoading]=useState(false)
    const confirm=async ()=>{
        const values=await form.validateFields()
        setLoading(true)
        // @ts-expect-error
        const res=(await request.post<HTTPResponse<Omit<ProjectItemType,"id" >>>("/newProject",values)) as HTTPResponse<Omit<ProjectItemType,"id" >>
        if (res.code===200){
            message.success("操作成功")
            form.resetFields()
            closeModal()
        }
        setLoading(false)


    }
    return (
        <Modal
            okButtonProps={{size:"large"}}
            cancelButtonProps={{size:"large"}}
            width="30%"
            closeIcon={false}
            open={isModalOpen}
            title="新增项目"
            onOk={confirm}
            confirmLoading={loading}
            onCancel={closeModal}>
            <Form form={form} size="large">
                <FormItem label="项目名称" name="projectName">
                    <Input placeholder="请输入项目名称"/>
                </FormItem>
                <FormItem label="项目地址" name="projectUrl">
                    <Input placeholder="请输入项目地址"/>
                </FormItem>
                <FormItem label="分支名称" name="projectBranch">
                    <Input placeholder="请输入分支名称"/>
                </FormItem>
                <FormItem label="用户名称" name="username">
                    <Input placeholder="请输入用户名称"/>
                </FormItem>
                <FormItem label="个人token" name="password">
                    <Input placeholder="请输入个人token"/>
                </FormItem>
            </Form>
        </Modal>
    );
};

export default NewProjectModal;