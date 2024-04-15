import axios from 'axios';
import {
    message
} from 'antd';
// 创建axios实例
const service = axios.create({
    baseURL: "/api/v1",
    // baseURL: "mock", // mock    timeout: 15000 // 请求超时时间
})




service.interceptors.response.use(
    response => {
        let res = response.data
        if (res.code != 0) {
            if (res.code == 401) {
                message.error('请先登录')
                return
            }
            message.error(res.msg)
            console.error(res)
            return Promise.reject(res)
        }
        return Promise.resolve(res.data);
    },
    err => {
        console.log(err.response)

        return Promise.reject(err.response.data)
    }
)



export default service