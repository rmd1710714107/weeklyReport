import axios from "axios";

const request=axios.create({
    baseURL:"http://localhost:3213"
})

request.interceptors.response.use(res=>{
    if (res.status===200){
        return res.data
    }else{
        return Promise.reject(res)
    }
},error => {
    return Promise.reject(error)
})
export default request