import {createHashRouter} from "react-router-dom";
import App from "@/App.tsx";
import {RouteObject} from "react-router/dist/lib/context";
import ProjectList from "@/views/projectList";
import ReportDetails from "@/views/reportDetails";

const routerList:RouteObject[]=[
    {
        path: "projectList",
        element:<ProjectList/>,
    },
    {
        path:"reportDetails",
        element:<ReportDetails/>
    }
]
const router=createHashRouter([
    {
        path:"/",
        element: <App/>,
        children:[
            {
                index:true,
                element:<ProjectList/>
            },
            ...routerList
        ]
    }
])
export default router