import Paragraph from "antd/es/typography/Paragraph";
import request from "@/network";
import {useEffect, useState} from "react";
import {HTTPResponse, ReportDetails} from "@/types";
import { Typography } from 'antd';

const { Title,Text } = Typography;

const Index = () => {
    const [reportDetails,setReportDetails]=useState<ReportDetails[]>([])
    const getRepostDetails=async ()=>{
        // @ts-expect-error
        const res=await request.get<HTTPResponse<ReportDetails[]>>("/getReportDetails") as HTTPResponse<ReportDetails[]>
        if (res.code===200){
            setReportDetails(res.data) 
        }
    }
    useEffect(()=>{
        getRepostDetails()
    },[])
    return (
        <Paragraph
            className="mt-0"
            strong>
            {reportDetails.map(item=>{
                if (item.reportContent===null) return null
                return (
                    <>
                        <Title level={2}>{item.projectName}</Title>
                        <ol>
                            {item.reportContent && item.reportContent.map(childItem=>(
                                <li>
                                    <Text>{childItem}</Text>
                                </li>
                                ))
                            }
                        </ol>

                    </>
                )
            })}

        </Paragraph>
    );
};

export default Index;