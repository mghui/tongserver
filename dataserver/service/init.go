package service

import (
	"github.com/astaxie/beego"
	"tongserver.dataserver/activity"
)

// CommonParamsType 请求的通用参数
type CommonParamsType struct {
	Name   string
	Params map[string]interface{}
}

// CriteriaInRBody 请求的rbody中的查询条件
type CriteriaInRBody struct {
	Field     string
	Operation string
	Value     interface{}
	Relation  string
}

type AggreStruct struct {
	Outfield  string
	Predicate string
	ColName   string
}

// SRequestBody 请求报文体
type SRequestBody struct {
	// Insert 新建
	Insert map[string]string
	// Update 更新
	Update map[string]string
	// Delete 删除
	Delete string
	// OperationConfirm 操作二次确认
	OperationConfirm string
	// Criteria 条件节点,针对更新、删除、查询操作
	Criteria []CriteriaInRBody
	// OrderBy 排序节点，针对查询操作
	OrderBy string
	// InnerJoin 内连接节点，针对查询操作
	InnerJoin string
	// Aggre 聚合节点，针对查询操作
	Aggre []AggreStruct
	// Bulldozer 推土机节点，针对查询操作
	Bulldozer []*CommonParamsType
	// PostAction 后处理节点，针对查询操作
	PostAction []*CommonParamsType
}

func (c *SRequestBody) IsEmpty() bool {
	return c.Insert == nil && c.Update == nil && c.Delete == "" && c.OperationConfirm == "" && c.Criteria == nil && c.OrderBy == "" && c.InnerJoin == "" && c.Aggre == nil && c.Bulldozer == nil && c.PostAction == nil
}

// init 初始化
func init() {
	SHandlerContainer[SrvTypeIds] = func(c RequestResponseHandler, caller string) SHandlerInterface {
		return &IDSServiceHandler{SHandlerBase{RRHandler: c, CurrentUserId: caller}}
	}
	SHandlerContainer[SrvTypePredef] = func(c RequestResponseHandler, caller string) SHandlerInterface {
		return &PredefineServiceHandler{IDSServiceHandler: IDSServiceHandler{SHandlerBase{RRHandler: c, CurrentUserId: caller}}}
	}
	SHandlerContainer[SrvValueKey] = func(c RequestResponseHandler, caller string) SHandlerInterface {
		return &ValueKeyService{SHandlerBase{RRHandler: c, CurrentUserId: caller}}
	}
	HASHSECRET = beego.AppConfig.String("jwt.token.hashsecret")
	TokenExpire, _ = beego.AppConfig.Int64("jwt.token.expire")

	activity.RegisterAcitvityCreator("innerservice", CreateInnerServiceActivity)
}
