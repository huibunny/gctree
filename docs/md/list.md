List tree.


## request

| field | type | required | description |
|:------|:------|:------:|:------------|
| master_id | string | Y | A master is a host which tree node belongs to: app_id, workspace_id, project_id, etc. |
| pid | string | N | parent id of the current tree node. list the subtree node if specified by pid, otherwise return the whole tree.|



## response

| field | type | description |
|:------|:---|:------------|
| errcode | int | error code: 200 - success, 400 - bad request, 404 - not found, 500 - internal server error.|
| message | string | error text |
| data | object | response data with the root node |
| data.meta | object | A meta is a node of a tree. |
| data.meta.meta_id | string | meta id |
| data.meta.meta_name | string | name of tree node |
| data.meta.extra | object | store extra meta info of tree node. |
| data.meta.extra.remark | string | remark text.  |
| data.meta.extra.icon | string | icon url: http://nebula.com/home.ico. |
| data.meta.pid | string | parent id of the current tree node. list the subtree node if specified by pid, otherwise return the whole tree. |
| data.children | array | child tree node list of the current tree node. |

## example

### request

```bash


curl "http://localhost:8850/v1/tree/list?master_id=f2c8e530-efb0-446f-be8f-6add7a2983a0" \
 -i \
 -X GET

```

### response

```json

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 19 Sep 2022 08:29:06 GMT
Content-Length: 1038

{
    "errcode": 0,
    "message": "success",
    "data": {
        "meta": {
            "meta_id": "f9605c0f-1074-434a-87a3-2585819beacf",
            "meta_name": "home",
            "extra": {
                "icon": "https://las.com/home.png",
                "remark": "hello home"
            },
            "pid": "f9605c0f-1074-434a-87a3-2585819beacf"
        },
        "children": [
            {
                "meta": {
                    "meta_id": "88bbee88-4972-42a6-b0ee-a4a281300bec",
                    "meta_name": "about",
                    "extra": {
                        "icon": "https://las.com/category.png",
                        "remark": "hello about"
                    },
                    "pid": "f9605c0f-1074-434a-87a3-2585819beacf"
                },
                "children": [
                    {
                        "meta": {
                            "meta_id": "620f6273-1a5b-4984-affe-55be14dda276",
                            "meta_name": "blog",
                            "extra": null,
                            "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"
                        },
                        "children": null
                    },
                    {
                        "meta": {
                            "meta_id": "f2c8e530-efb0-446f-be8f-6add7a2983a0",
                            "meta_name": "event",
                            "extra": {
                                "icon": "https://las.com/event.png",
                                "remark": "hello event"
                            },
                            "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"
                        },
                        "children": null
                    }
                ]
            },
            {
                "meta": {
                    "meta_id": "22e27691-8b8d-45e3-87ae-2d20a9448767",
                    "meta_name": "news",
                    "extra": {
                        "icon": "https://las.com/category.png",
                        "remark": "good news!"
                    },
                    "pid": "f9605c0f-1074-434a-87a3-2585819beacf"
                },
                "children": null
            }
        ]
    }
}

```


