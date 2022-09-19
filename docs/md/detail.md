Show tree detail info by meta id.


## request

| field | type | required | description |
|:------|:------|:------:|:------------|
| meta_id | string | Y | meta id. |



## response

| field | type | description |
|:------|:---|:------------|
| errcode | int | error code: 200 - success, 400 - bad request, 404 - not found, 500 - internal server error.|
| message | string | error text |
| data | object | response data with the root node |
| data.tenant_id | string | A meta is a node of a tree. |
| data.user_id | string | meta id |
| data.master_id | string | name of tree node |
| data.meta_list | object | store extra meta info of tree node. |
| data.meta_list.meta_id | string | meta id |
| data.meta_list.meta_name | string | name of tree node |
| data.meta_list.extra.remark | string | remark text.  |
| data.meta_list.extra.icon | string | icon url: http://nebula.com/home.ico. |
| data.meta_list.pid | string | parent id of the current tree node. list the subtree node if specified by pid, otherwise return the whole tree. |

## example

### request

```bash

curl "http://localhost:8850/v1/tree/detail?meta_id=88bbee88-4972-42a6-b0ee-a4a281300bec" \
 -i \
 -X GET

```

### response

```json

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 19 Sep 2022 09:26:30 GMT
Content-Length: 358

{
    "errcode": 0,
    "message": "success",
    "data": {
        "tenant_id": "fcb3883b-7847-40d1-a89c-03e70db89bb5",
        "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000",
        "master_id": "fcb3883b-7847-40d1-a89c-03e70db89111",
        "meta_list": [
            {
                "meta_id": "88bbee88-4972-42a6-b0ee-a4a281300bec",
                "meta_name": "about",
                "extra": {
                    "icon": "https://las.com/category.png",
                    "remark": "hello about"
                },
                "pid": ""
            }
        ]
    }
}

```


