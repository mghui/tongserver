package utils

import "github.com/astaxie/beego/cache"

// DataSourceCache 用于保存数据集的配置信息的缓存
var DataSourceCache, _ = cache.NewCache("memory", `{"interval":60}`)

// DictDataCache 用于保存数据字典信息
var DictDataCache, _ = cache.NewCache("memory", `{"interval":60}`)

// DataSetResultCache 用于保存结果集的缓存
var DataSetResultCache, _ = cache.NewCache("memory", `{"interval":60}`)

const (
	CACHE_PREFIX_SERVICEACCESS string = "SR_ACCESS_"
)

// JedaDataCache 用于保存Jeda管理信息
var JedaDataCache, _ = cache.NewCache("memory", `{"interval":60}`)
