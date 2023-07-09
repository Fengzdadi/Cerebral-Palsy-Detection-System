# 接口文档

## GET请求部分
1. 测试请求
   + /Hello
   + return (json)
      + `message` "Hello, World!"
2. 获取用户基本信息
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
3. 获取用户历史检测数据
   + /UserHisResult
   + return (json)
      + eg.
      {
        "Userid": 1,
        "ResultData": [
        {
        "Time": "2023-07-05T08:30:00Z",
        "Result": 0.87,
        "ResultAdd": "/path/to/result/0001"
        },
        {
        "Time": "2023-07-05T08:23:00Z",
        "Result": 0.76,
        "ResultAdd": "/path/to/result/0002"
        }
        ],
        "Count": 2
        }
      + `Userid` ID
      + `ResultData` 检测数据 array
        + `Time` 检测时间
        + `Result` 检测结果
        + `ResultAdd` 检测结果文件地址
      + `Count` 检测数据数量
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
