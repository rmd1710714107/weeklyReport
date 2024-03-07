import {Button, Space} from "antd";
import ProjectItem from "@/views/projectList/childComp/ProjectItem";
import {useNavigate } from "react-router-dom";
import NewProjectModal from "@/views/projectList/childComp/NewProjectModal";
import {useEffect, useState} from "react";
import request from "@/network";
import {HTTPResponse,ProjectItemType} from "@/types";

const Index=()=>{
    // const projectList=new Array(100).fill(null)
    const navigate=useNavigate()
    const toReportDetails=async ()=>{
        navigate("reportDetails")
    }
    const [flag,setFlag]=useState(false)
    const closeModalCB=()=>{
        setFlag(false)
        getProjectList()
    }

    const [projectList,serProjectList]=useState<ProjectItemType[]>([])
    const getProjectList=async ()=>{
        // @ts-expect-error
        const res=(await request.get<HTTPResponse<ProjectItemType[]>>("/getProjectList")) as HTTPResponse<ProjectItemType[]>
        console.log(res)
        if (res.code){
            serProjectList(res.data)
        }
    }
    useEffect(()=>{
        getProjectList()
    },[])
    return (
        <>
            <Space className="w-full justify-end">
                <Button size="large" onClick={()=>setFlag(true)}>
                    新增项目
                </Button>
                <Button type="primary" size="large" onClick={toReportDetails}>查看周报</Button>
            </Space>
            <section className="mt-5 flex-1 w-full overflow-auto">
                {projectList.map((item,index)=>(<ProjectItem item={item} deleteCb={getProjectList} key={index}/>))}
            </section>
            <NewProjectModal flag={flag} closeModalCB={closeModalCB}/>
        </>
    )
}
export default Index
