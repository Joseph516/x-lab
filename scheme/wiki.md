| Version | Update Time | Status | Author | Description |
| :------ | :---------- | :----- | :----- | :---------- |
| 1.0.0   |             | create | hyh    |             |

## 用户管理

### 用户注册

**URL:** /user/register

**Type:** POST

**Content-Type:** application/x-www-form-urlencoded

**Description:** 用户注册，已经实现两次密码校验，用户名或密码长度校验，用户已经注册校验。

**Request-parameters:**

| Parameter  | Type   | Description | Required       | Since |
| :--------- | :----- | :---------- | :------------- | :---- |
| username   | string | 用户名      | true，长度2-20 |       |
| password   | string | 密码        | true，长度4-32 |       |
| Repassword | String | 密码确认    | true，长度4-32 |       |

**Request-example:**

```shell
curl -X POST http://127.0.0.1:8080/user/register -d "username=hyh&password=1234&repassword=1234" -H "Content-Type:application/x-www-form-urlencoded"
```

**Response-fields:**

| Field   | Type   | Description  | Since |
| :------ | :----- | :----------- | :---- |
| code    | int32  | 状态码       | -     |
| msg     | string | 状态信息     | -     |
| details | string | 状态相信信息 |       |

**Response-example:**

```json
// 注册成功
{
    "code": 200,
    "msg": "success"
}
// 注册失败
{
    "code": 2003,
    "details": [
        "用户已经存在，请直接登陆"
    ],
    "msg": "注册失败"
}
// 入参错误
{
    "code": 10000001,
    "details": [
        "Key: 'RegisterRequest.Username' Error:Field validation for 'Username' failed on the 'required' tag",
        "Key: 'RegisterRequest.Password' Error:Field validation for 'Password' failed on the 'required' tag",
        "Key: 'RegisterRequest.Repassword' Error:Field validation for 'Repassword' failed on the 'required' tag"
    ],
    "msg": "入参错误"
}
```

### 用户登陆

**URL:** /user/login

**Type:** POST

**Content-Type:** application/x-www-form-urlencoded

**Description:** 用户注册，已经实现两次密码校验，用户名或密码长度校验，用户已经注册校验。

**Request-parameters:**

| Parameter | Type   | Description | Required       | Since |
| :-------- | :----- | :---------- | :------------- | :---- |
| username  | string | 用户名      | true，长度2-20 |       |
| password  | string | 密码        | true，长度4-32 |       |

**Request-example:**

```shell
curl -X POST http://127.0.0.1:8080/user/login -d "username=hyh&password=1234" -H "Content-Type:application/x-www-form-urlencoded"
```

**Response-fields:**

| Field   | Type   | Description  | Since |
| :------ | :----- | :----------- | :---- |
| code    | int32  | 状态码       | -     |
| msg     | string | 状态信息     | -     |
| details | string | 状态相信信息 |       |

**Response-example:**

```json
// 登陆成功
{
    "code": 200,
    "msg": "success"
}
// 登陆失败
{
    "code": 2004,
    "details": [
        "record not found"
    ],
    "msg": "登陆失败，用户或密码错误"
}
```

## 文件管理

### Upload files

**URL:** /upload/file

**Type:** POST

**Content-Type:** multipart/form-data

**Description:** 上传文件

**Request-parameters:**

| Parameter | Type   | Description         | Required | Since |
| :-------- | :----- | :------------------ | :------- | :---- |
| file      | string | 文件路径            | true     |       |
| type      | int    | 文件类型。1：图片； | true     |       |

**Request-example:**

```shell
curl -X POST http://127.0.0.1:8080/upload/file -F file=@'/Users/joe/Pictures/wallpaper/wolf.jpeg' -F type=1
```

**Response-fields:**

| Field             | Type   | Description    | Since |
| :---------------- | :----- | :------------- | :---- |
| code              | int32  | 状态码         | -     |
| msg               | string | 状态信息       | -     |
| data              | object | 具体结果       | -     |
| └─file_access_url | string | 文件可访问路径 | -     |

**Response-example:**

```json
{
    "code": 200,
    "data": {
        "file_access_url": "http://127.0.0.1:8080/static/87221652a79fc3c9b04cde0b335fdd5b.jpeg"
    },
    "msg": "success"
}
```

## Notes