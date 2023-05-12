# 接口文档

## GET请求部分
1. 获取用户基本信息
    + /UserBaseInfo
    + body
        + `Username` 用户名
    + return (json)
        + `Userid` ID
        + `Username` 用户名
        + `Age` 年龄
        + `Gender` 性别
        + `Phone` 手机号
        + `Email` 邮箱
## POST请求部分
1. 用户登录
    + /UserLogin
    + body
        + `Username` 用户名
        + `Password` 密码
    + return (json)
        + `message` "UserLogin, No user match!"
        + `message` "UserLogin, Success!"
        + `message` "UserLogin, Failed!"
      
2. 视频上传
    + /VideoUpload
    + body
        + `Video` 上传视频文件
    + return(json)
        + `state`: "successful"
        + `filePath`: dst
        + `videoName`: Video.Filename
        + `fileId`: fileId
