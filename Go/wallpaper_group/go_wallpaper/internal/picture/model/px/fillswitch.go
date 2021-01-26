package px

import "github.com/jinzhu/gorm"

type Fillswitch struct {
	gorm.Model
	AccessDeleted        bool        `json:"access_deleted"`
	AccessPrivate        bool        `json:"access_private"`
	IncludeDeleted       bool        `json:"include_deleted"`
	ExcludePrivate       bool        `json:"exclude_private"`
	ExcludeNude          bool        `json:"exclude_nude"`
	AlwaysExcludeNude    bool        `json:"always_exclude_nude"`
	ExcludeBlock         bool        `json:"exclude_block"`
	CurrentUserID        interface{} `json:"current_user_id"`
	OnlyUserActive       bool        `json:"only_user_active"`
	IncludeTags          bool        `json:"include_tags"`
	IncludeGeo           bool        `json:"include_geo"`
	IncludeLicensing     bool        `json:"include_licensing"`
	IncludeAdminLocks    bool        `json:"include_admin_locks"`
	IncludeLikeBy        bool        `json:"include_like_by"`
	IncludeComments      bool        `json:"include_comments"`
	IncludeUserInfo      bool        `json:"include_user_info"`
	IncludeFollowInfo    bool        `json:"include_follow_info"`
	IncludeEquipmentInfo bool        `json:"include_equipment_info"`
}
