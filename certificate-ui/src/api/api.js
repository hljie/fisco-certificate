import request from "./config"
import urls from "./urls"

const api = {}

const getRequest = {
    "get": (config) => {
        return (params) => {
            return request(Object.assign({}, config, {
                params
            }));
        }
    },
    "post": (config) => {
        return (data) => {
            return request(Object.assign({}, config, {
                data
            }));
        }
    },
    "put": (config) => {
        return (data) => {
            return request(Object.assign({}, config, {
                data
            }));
        }
    },
    "patch": (config) => {
        return (data) => {
            return request(Object.assign({}, config, {
                data
            }));
        }
    },
}

Object.keys(urls).forEach(group => {
    Object.keys(urls[group]).forEach(methodName => {
        Object.keys(urls[group][methodName]).forEach(name => {
            if(api[group] === undefined){
                api[group] = {}
            }
            api[group][name] = getRequest[urls[group][methodName][name].method](urls[group][methodName][name])
        })
    })
})

console.log(api)
export default api;