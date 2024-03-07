export type HTTPResponse<T>={
    code:number
    msg:string
    data:T
}
export type ProjectItemType={
    id:number
    projectUrl:string
    projectName:string
    projectBranch:string
    password:string
    username:string
}
export type ReportDetails={
    projectName:string
    reportContent:string[]
}