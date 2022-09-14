# test

## save

```bash

# create
curl -X POST -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_name": "home","extra": {"remark": "hello home", "icon": "https://las.com/home.png"}, "pid": ""}]}}' "http://localhost:8850/v1/page_category/save"

curl -X POST -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_name": "about","extra": {"remark": "hello about", "icon": "https://las.com/category.png"}, "pid": "fcb3883b-7847-40d1-a89c-03e70db89111"}]}}' "http://localhost:8850/v1/page_category/save"



curl -X POST -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_name": "blog","extra": {"remark": "hello blog", "icon": "https://las.com/category.png"}, "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"}, {"meta_name": "news","extra": {"remark": "hello news", "icon": "https://las.com/category.png"}, "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"}]}}' "http://localhost:8850/v1/page_category/save"


# update
curl -X POST -d '{"app_list": [{"app_id": "fcb3883b-7847-40d1-a89c-03e70db89bb5", "tenant_id":"", "user_id": "", "project_id": "", "app_name": "las", "app_type_id": "", "icon": "https://las.com/las.ico", "cover": "https://las.com/cover.png", "brief": "hello"}]}' "http://localhost:8830/v1/app/save"


# list
curl -X GET -H "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfdGltZSI6MTY2MDcyMzgwNSwidXNlcl9pZCI6IjEwMzM3ZmM3LWE2ZjEtNDM0My05NTYxLWRmNTZkZjkwMTFiMCJ9.5Tv4pUbpl17V1zoe8puX7JfGFhhgtc9rzKoHVUnmGo8" "http://dog.ap:9090/api/app/v1/app/list"
curl -X GET "http://localhost:8830/v1/app/list"

# detail
curl -X GET "http://localhost:8850/v1/page_category/detail?meta_id=88bbee88-4972-42a6-b0ee-a4a281300bec"

# delete
curl -X POST -d '{"meta_id_list": ["88bbee88-4972-42a6-b0ee-a4a281300bec"]}' "http://localhost:8850/v1/page_category/delete"

curl -X POST -d '{"meta_id_list": ["620f6273-1a5b-4984-affe-55be14dda276", "22e27691-8b8d-45e3-87ae-2d20a9448767"]}' "http://localhost:8850/v1/page_category/delete"

```
