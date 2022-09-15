# test

## save

```bash

# create
curl -X POST -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_name": "home","extra": {"remark": "hello home", "icon": "https://las.com/home.png"}, "pid": ""}]}}' "http://localhost:8850/v1/page_category/save"

curl -X POST -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_name": "about","extra": {"remark": "hello about", "icon": "https://las.com/category.png"}, "pid": "fcb3883b-7847-40d1-a89c-03e70db89111"}]}}' "http://localhost:8850/v1/page_category/save"



curl -X POST -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_name": "blog","extra": {"remark": "hello blog", "icon": "https://las.com/category.png"}, "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"}, {"meta_name": "news","extra": {"remark": "hello news", "icon": "https://las.com/category.png"}, "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"}]}}' "http://localhost:8850/v1/page_category/save"


# update
curl -X POST -d '{"tree_node": {"tenant_id":"fcb3883b-7847-40d1-a89c-03e70db89bb5", "master_id":"fcb3883b-7847-40d1-a89c-03e70db89111", "user_id": "fcb3883b-7847-40d1-a89c-03e70db89000", "meta_list": [{"meta_id": "22e27691-8b8d-45e3-87ae-2d20a9448767", "meta_name": "news","extra": {"remark": "good news!", "icon": "https://las.com/category.png"}, "pid": "f9605c0f-1074-434a-87a3-2585819beacf"}]}}' "http://localhost:8850/v1/page_category/save"


# list
curl -X GET "http://localhost:8850/v1/page_category/list?master_id=fcb3883b-7847-40d1-a89c-03e70db89111"

{
    "meta_node": {
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
                            "meta_id": "22e27691-8b8d-45e3-87ae-2d20a9448767",
                            "meta_name": "news",
                            "extra": null,
                            "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"
                        },
                        "children": null
                    },
                    {
                        "meta": {
                            "meta_id": "620f6273-1a5b-4984-affe-55be14dda276",
                            "meta_name": "blog",
                            "extra": null,
                            "pid": "88bbee88-4972-42a6-b0ee-a4a281300bec"
                        },
                        "children": null
                    }
                ]
            }
        ]
    }
}

# detail
curl -X GET "http://localhost:8850/v1/page_category/detail?meta_id=88bbee88-4972-42a6-b0ee-a4a281300bec"

# delete
curl -X POST -d '{"meta_id_list": ["88bbee88-4972-42a6-b0ee-a4a281300bec"]}' "http://localhost:8850/v1/page_category/delete"

curl -X POST -d '{"meta_id_list": ["620f6273-1a5b-4984-affe-55be14dda276", "22e27691-8b8d-45e3-87ae-2d20a9448767"]}' "http://localhost:8850/v1/page_category/delete"

```
