Save without meta_id or update by meta_id a tree.


## request

| field | type | required | description |
|:------|:------------|:---:|:----|
| tree_node | object | Y | A subtree node to be saved. |
| tree_node.tenant_id | string | Y | A tenent id is used to separate data from each tenancy. |
| tree_node.master_id | string | Y | A master is a host which tree node belongs to: app_id, workspace_id, project_id, etc. |
| tree_node.user_id | string | Y | User who does the operation. |
| tree_node.meta_list | array | Y | A meta is a node of a tree and meta list is a collection of tree nodes. |
| tree_node.meta_list.meta_id | string | N | id of tree node, update if specified, otherwise create it. |
| tree_node.meta_list.meta_name | string | Y | name of tree node. |
| tree_node.meta_list.extra | object | N | store extra meta info of tree node. |
| tree_node.meta_list.extra.remark | string | N | remark text. |
| tree_node.meta_list.extra.icon | string | N | icon url: http://nebula.com/home.ico. |
| tree_node.meta_list.pid | string | N | parent id of the current tree node. list the subtree node if specified by pid, otherwise return the whole tree.|



## response

| field | type | description |
|:------|:----|:------------|
| errcode | int | error code: 200 - success, 400 - bad request, 404 - not found, 500 - internal server error.|
| message | string | error text |
| data | object | response data |
| data.meta_id_list | array | tree node id created successfully. |

## example

### create

#### request

```bash


curl 'http://localhost:8850/v1/tree/save' \
  -i \
  -X 'POST' \
  -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_name": "event","extra": {"remark": "hello event", "icon": "https://las.com/event.png"}, "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"}]}}'

```

#### response

```json

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 19 Sep 2022 07:31:31 GMT
Content-Length: 98

{
    "errcode": 0,
    "message": "success",
    "data": {
        "meta_id_list": [
            "f2c8e530-efb0-446f-be8f-6add7a2983a0"
        ]
    }
}

```


### update

#### request


```bash

curl 'http://localhost:8850/v1/tree/save' \
  -i \
  -X 'POST' \
  -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_id": "f2c8e530-efb0-446f-be8f-6add7a2983a0", "meta_name": "event","extra": {"remark": "hi event", "icon": "https://las.com/event.png"}, "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"}]}}'

```

#### response

```json

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 19 Sep 2022 11:30:20 GMT
Content-Length: 98

{
    "errcode": 0,
    "message": "success",
    "data": {
        "meta_id_list": [
            "f47db7ef-f1f9-4335-989c-95ab74c1dfd1"
        ]
    }
}

```
