# Wiki

| Version | Update Time | Status | Author | Description |
| :------ | :---------- | :----- | :----- | :---------- |
| 1.0.0   |             | create | hyh    |             |

## 1. Upload files

**URL:** /upload/file

**Type:** POST

**Content-Type:** multipart/form-data;charset=utf-8

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