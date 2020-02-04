package orm

import "time"

type BaseModel struct {
	Id uint `json:"id" gorm:"PRIMARY_KEY;AUTO_INCREMENT"`
	CreatedAt time.Time `json:"-" gorm:"default: CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
}

// 用户信息（不包含密码或任何登录信息
type User struct {
	BaseModel

	Name string `gorm:"type:varchar(32);unique;not null;comment:'用户名'"`
	Avatar string `sql:"type:text;" gorm:"comment:'头像地址'"`
}

func (User) TableName() string {
	return "user"
}

// 用户授权记录，每个用户可有多个授权记录用于支持不同登录方式
type UserAuth struct {
	BaseModel

	UId uint `gorm:"not null;unique_index:idx_uid_type;comment:'对应的用户 id'"`
	Verified bool `gorm:"not null;default:0;comment:'登录方式是否已验证'"`
	AuthType uint `gorm:"not null;unique_index:idx_uid_type;comment:'身份验证方式 id'"`
	AuthId string `gorm:"not null;unique;comment:'身份验证唯一 id（如手机号/邮箱/第三方登录唯一 id）'"`
	Credential string `gorm:"not null;comment:'凭证（账户密码/第三方登录 token）'"`
	LatestLoginAt time.Time `json:"-" gorm:"comment:'最后一次使用此身份验证方式登录时间'"`
	IpAddr uint32 `sql:"type:int unsigned" gorm:"comment:'最后一次使用此身份验证方式登录时的 IP'"`
}

func (UserAuth) TableName() string {
	return "user_auth"
}

// 身份验证类型表
type Identification struct {
	BaseModel

	Name string `gorm:"type:varchar(32);unique;not null;comment:'身份验证方式'"`
}
