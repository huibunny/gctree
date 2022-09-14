// Package entity defines main entities for business logic (services), data base mapping and
// HTTP response objects if suitable. Each logic group entities in own file.
package entity

type Meta struct {
	MetaID     string                 `json:"meta_id"  example:"sdfasdflksajflksadjflk000"`
	MetaName   string                 `json:"meta_name"  example:"las"`
	Extra      map[string]interface{} `json:"extra"  example:"extra map info"`
	PID        string                 `json:"pid" example:"fjlksdajklsadjfkl"`
	CreateTime string                 `json:"create_time" example:"2022-10-11 20:20:20"`
	UpdateTime string                 `json:"update_time" example:"2022-10-11 20:20:20"`
	IsDel      bool                   `json:"is_del" example:"false"`
}

// TreeNode -.
type TreeNode struct {
	TenantID string `json:"tenant_id"  example:"sdfasdflksajflksadjflk111"`
	UserID   string `json:"user_id"  example:"sdfasdflksajflksadjflk111"`
	MasterID string `json:"master_id"  example:"sdfasdflksajflksadjflk111"`
	MetaList []Meta `json:"meta_list" example:"dsffasfdfsdafsdafsdasadf"`
}
