# API文档

## 简介
这个API用于**CPDS**的app系统

## 基础URL
`http://localhost:8080`

## 认证
使用 Bearer Token 进行认证

## 请求和响应格式
大部分请求和响应均为JSON格式

## 错误代码
|          错误代码           |  说明   |
|:-----------------------:|:-----:|
|         SUCCESS         |  200  |
|  UpdatePasswordSuccess  |  201  |
|   NotExistIdentifier    |  202  |
|          ERROR          |  500  |
|      InvalidParams      |  400  |
|      ErrorDatabase      | 40001 |
| WebsocketSuccessMessage | 50001 |
|    WebsocketSuccess     | 50002 |
|      WebsocketEnd       | 50003 |
|  WebsocketOnlineReply   | 50004 |
|  WebsocketOfflineReply  | 50005 |
|     WebsocketLimit      | 50006 |

## 端点

### 用户注册
+ HTTP方法：`POST`
+ URL：`/UserRegister`
+ 参数：
  + `user_name`: 用户名（5-15位）
  + `password`: 密码（8-16位）
+ 请求示例
    ```json
  {
      "user_name": "CAI_XU_KUN",
      "password": "12345678"
  }
  ```
  
+ 响应示例：
    ```json
    {
      "code": 200,
      "data": null,
      "msg": "ok",
      "error": ""
    }
    ```
  
### 用户登录
+ HTTP方法：`POST`
+ URL：`/UserLogin`
+ 参数：
  + `user_name`: 用户名（5-15位）
  + `password`: 密码（8-16位）
+ 请求示例
    ```json
  {
      "user_name": "CAI_XU_KUN",
      "password": "12345678"
  }
  ```
  
+ 响应示例：
    ```json
    {
      "res": {
      "code": 200,
      "data": "CAI_XU_KUN",
      "msg": "ok",
      "error": ""
    },
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2OTM4MDgxNTMsInVzZXJuYW1lIjoiIn0.J9yc5CsVEnk5ghOSB-kKgaUuOzgIVi657D7werQwQC4"
    }
    ```

### 获取孩子信息
+ HTTP方法：`GET`
+ URL：`/GetChildInfo`
+ 参数： NULL
+ 请求示例
  ```bash
    GET localhost:8080/GetChildInfo
    ```
  
+ 响应示例
  ```json
    {
      "error": {
          "code": 200,
          "data": null,
          "msg": "ok",
          "error": ""
      },
      "message": [
          {
              "ID": 31474683,
              "CreatedAt": "2023-09-01T11:25:18Z",
              "UpdatedAt": "2023-09-01T11:25:18Z",
              "DeletedAt": null,
              "BelongTo": 3,
              "ChildName": "caijijiji",
              "Age": 5,
              "Gender": "Male",
              "BirthDate": "20001010"
          }
      ]
    }
    ```

### 添加孩子信息
+ HTTP方法：`POST`
+ URL：`/AddChildInfo`
+ 参数：
  + `ChildName`: 孩子姓名
  + `Age`: 孩子年龄
  + `Gender`: 孩子性别
  + `BirthDate`: 孩子出生日期
+ 请求示例
  ```json
  {
    "ChildName": "caijijiji",
    "Age": "5",
    "Gender": "Male",
    "BirthDate": "20001010"
  }
    ```
  
+ 响应示例
  ```json
  {
    "code": 200,
    "data": null,
    "msg": "ok",
    "error": ""
  }
    ```

### 修改孩子信息(未实现)

### 增加孩子历史基础信息
+ HTTP方法：`POST`
+ URL：`/AddBaseInfoHis`
+ 参数：
  + `BelongToChildID`: 孩子ID
  + `height`: 孩子当期体重
  + `weight`: 孩子当期身高
+ 请求示例
  ```json
  {
    "BelongToChildID": "82440514",
    "height": "86.0",
    "weight": "11.3"
  }
    ```
  
+ 响应示例
  ```json
  {
    "code": 200,
    "data": null,
    "msg": "ok",
    "error": ""
  }
    ```
### 获取孩子基本历史信息列表
+ HTTP方法：`GET`
+ URL：`/GetBaseInfoHis`
+ 参数：Query
  + `ChildId`: 孩子ID
+ 请求示例
  ```bash
  GET localhost:8080/GetBaseInfoHis?ChildId=82440514
    ```
  
+ 响应示例
  ```json
    {
    "error": {
        "code": 200,
        "data": null,
        "msg": "ok",
        "error": ""
      },
    "message": [
        {
            "BelongToChildID": 82440514,
            "Time": "2023-09-01T13:50:34Z",
            "Height": 0,
            "Weight": 0
        },
        {
            "BelongToChildID": 82440514,
            "Time": "2023-09-01T13:53:50Z",
            "Height": 0,
            "Weight": 0
        },
        {
            "BelongToChildID": 82440514,
            "Time": "2023-09-01T13:54:57Z",
            "Height": 86,
            "Weight": 11.3
        }
      ]
   }
    ```

## GET请求部分
1. 测试请求
   + /Hello
   + return (json)
      + `message` "Hello, World!"

2. 获取用户历史检测数据
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

3. 返回视频测试结果 （X）
   + /ReturnVideoResult
   + return (json)
      + `Result` 检测结果
      + `ResultAdd` 检测结果文件地址

## POST请求部分
1. 用户登录 (X)
    + /UserLogin
    + body
        + `Username` 用户名
        + `Password` 密码
    + return (json)
        + `message` "UserLogin, No user match!"
        + `message` "UserLogin, Success!"
        + `message` "UserLogin, Failed!"

2. 获取用户基本信息 (X)
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
      
3. 视频上传
    + /VideoUpload
    + body
        + `Video` 上传视频文件
    + return(json)
        + `state`: "successful"
        + `filePath`: dst
        + `videoName`: Video.Filename
        + `fileId`: fileId

4. 开始测试
    + /StartDetection
    + body
        + `VideoName` 视频名
        + `Userid` 用户ID
        + `VideoPath` 视频路径
        + `VideoRes` 结果
        + `Probability` 检测结果

5. 开始测试 测试版
    + /StartDetectionTest
    + return (json)
        + `VideoName`: "successful"
        + `Userid`: dst
        + `VideoPath`: Video.Filename
        + `VideoRes`: Video.Filename
        + `Probability`: fileId

## WS请求部分
1. 请求对话
    + /UserRegister
    + parameters
        + `uid` 请求用户
        + `toUid` 对象用户
