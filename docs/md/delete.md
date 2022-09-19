Delete tree by meta id list.


## request

| field | type | required | description |
|:------|:------------|:---:|:----|
| data.meta_id_list | array | Y | tree node id list. |



## response

| field | type | description |
|:------|:----|:------------|
| errcode | int | error code: 200 - success, 400 - bad request, 404 - not found, 500 - internal server error.|
| message | string | error text |
| data | object | response data is always null value. |

## example

### request

```bash

curl "http://localhost:8850/v1/tree/delete" \
  -i \
  -X 'POST' \
  -d '{"meta_id_list": ["f2c8e530-efb0-446f-be8f-6add7a2983a0"]}'

```

#### response

```json

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
Date: Mon, 19 Sep 2022 11:21:54 GMT
Content-Length: 45

{
    "errcode": 0,
    "message": "success",
    "data": null
}

```

