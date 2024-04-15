const methodsTemplate = {
    "get": {
        method: "get"
    },
    "post": {
        method: "post"
    },
    "put": {
        method: "put"
    },
    "patch": {
        method: "patch"
    }
};
const urlsTemplate = {
    "user": {
        "get": {
            "info": "user/me",
            "logout": "user/logout",
        },
        "post": {
            "register": "user/register",
            "login": "user/login",
            "list": "users",
            "updateMe": "user/update/me",
            "update": "user/update"
        }
    },
    "certificate": {
        "post": {
            "me": "certificate/me",
            "apply": "certificate/apply",
            "update": "certificate/update",
            "id": "certificate/id",
            "list": "certificates",
            "approve": "certificate/approve",
        }
    }

};
const urls = {};
Object.keys(urlsTemplate).forEach(function (group) {
    Object.keys(urlsTemplate[group]).forEach(function (methodName) {
        Object.keys(urlsTemplate[group][methodName]).forEach(function (name) {
            if (typeof urlsTemplate[group][methodName][name] === 'string') {
                urlsTemplate[group][methodName][name] = {
                    url: urlsTemplate[group][methodName][name]
                };
            }
            if (urls[group] === undefined) {
                urls[group] = {};
            }
            if (urls[group][methodName] === undefined) {
                urls[group][methodName] = {};
            }
            urls[group][methodName][name] = Object.assign({}, methodsTemplate[methodName], urlsTemplate[group][methodName][name]);
        });
    });
});
console.log(urls)
export default urls