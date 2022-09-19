// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

type Meta struct {
	MetaID   string                 `json:"meta_id"  example:"fcb3883b-7847-40d1-a89c-03e70db89333"`
	MetaName string                 `json:"meta_name"  example:"meta1"`
	Extra    map[string]interface{} `json:"extra" swaggertype:"object,string"  example:"remark:remark1,icon:icon1"`
	PID      string                 `json:"pid" example:"fcb3883b-7847-40d1-a89c-03e70db89fff"`
	// CreateTime string                 `json:"create_time" example:"2022-10-11 20:20:20"`
	// UpdateTime string                 `json:"update_time" example:"2022-10-11 20:20:20"`
	// IsDel      bool                   `json:"is_del" example:"false"`
}

// TreeNode -.
type TreeNode struct {
	TenantID string `json:"tenant_id"  example:"fcb3883b-7847-40d1-a89c-03e70db89111"`
	UserID   string `json:"user_id"  example:"fcb3883b-7847-40d1-a89c-03e70db89222"`
	MasterID string `json:"master_id"  example:"fcb3883b-7847-40d1-a89c-03e70db89000"`
	MetaList []Meta `json:"meta_list"`
}

type MetaNode struct {
	Meta     Meta       `json:"meta"`
	Children []MetaNode `json:"children"  example:"children list"`
}
