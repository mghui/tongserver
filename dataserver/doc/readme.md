# 基于BeeGo的数据查询服务
## 基本功能
* 根据SQL语句发布Restful接口服务
    * SQL语法分析
    * 分页支持**√**
    * 字典数据支持**√**
    
* 将数据库表直接发布为Restful
    * 单表主键查询**√**
    * 单表条件查询**√**
    * 字典数据支持**√**
    * 支持主从数据查询
    * 支持分页**√**
    * 支持计数、求和、求平均、最大值、最小值等统计函数**√**
    
* 基础用户、角色、组织机构、权限数据模型和数据操作接口
    * 基本数据CURD操作接口**√**
    
* 基于JWT的安全控制接口
* 基于OAtuh2的安全控制接口
* Qrcode的二维码生成器接口**√**
* 支持Oracle、MySQL数据库**√**
## 基本属性
### 数据源类型
```go

type DataSourceType int8

const (
	DataSourceType_SQL      DataSourceType = 1
	DataSourceType_SQLTABLE DataSourceType = 2
	DataSourceType_REST     DataSourceType = 3
	DataSourceType_ENMU     DataSourceType = 4
	DataSourceType_INNER    DataSourceType = 5
)
```
### 通用数据类型
```go
const (
	Property_Datatype_INT  string = "INT"    //整数
	Property_Datatype_DOU  string = "DOUBLE" //浮点数 
	Property_Datatype_STR  string = "STRING" //字符串
	Property_Datatype_DATE string = "DATE"   //日期类型
	Property_Datatype_TIME string = "TIME"   //包含日期的时间类型
	Property_Datatype_ENUM string = "ENUM"   //枚举类型
	Property_Datatype_DS string = "DATASET"  //数据集，表示该数据数值是一个数据集
	Property_Datatype_UNKN string = ""       //未知类型
)
```
### 属性
```go
type MyProperty struct {
	Name          string //属性名
	DataType      string //类型名in
	OutJoin       bool   //是否外联接
	Caption       string //显示名
	OutJoinDefine *OutFieldProperty
}
```
> 属性的外联接：OutJoin属性为true时，该属性为外链接属性，通过OutJoinDefine属性配置具体外链接的逻辑。
> 引擎在加载该属性时通过外链接配置信息获取外部数据填充该属性。外链接数据源是基础数据的一种即可。外链接配置信息包括：
>
> ```go
> type OutFieldProperty struct {
> 	Source IDataSource
> 	JoinField  string
> 	ValueField string
> 	ValueFunc  func(record []interface{}, field []*MyProperty, Source IDataSource) interface{}
> }
> ```
## 基础数据源
```go
type IDataSource interface {
	//返回数据源类型
	GetDataSourceType() DataSourceType
	//数据源初始化
	Init() error
	GetName() string
	//返回全部数据
	GetAllData() (*DataResultSet, error)
	//////返回主键信息
	//GetKeyFields() []MyProperty
	//根据主键返回数据
	QueryDataByKey(keyvalues ...interface{}) (*DataResultSet, error)
	//根据字段值返回匹配的数据
	QueryDataByFieldValues(fv *map[string]interface{}) (*DataResultSet, error)
}
```
## 数据源
* 数据表数据源 **√**

* SQL数据源 **√**

* ValueKey数据源 **√**

  枚举数据源用于数据字典，在数据集后处理中可以作为数据字典使用

* Restful数据源

* Webservice数据源

* 静态数据源

* ~~数据源联接~~

* ~~数据源组合~~

## 数据查询
* 主键匹配查询 **√**
* 字段值匹配查询 **√**
* 复合条件查询 **√**
* 时间序列处理
  * 针对DataTableSource，定义时间列
  * 根据分组返回最新数据
  * 返回最新的数据
  * 返回去年同期数据
  * 返回时间段数据

## 结果集后处理
* 数据分组 **√**
* 数据切面 **√**
* 格式化 **√**
* 列过滤 **√**
* 数据字典映射 **√**
* 列提取 **√**
* 数据集拼接
* 数据集缓存**√**

## 数据服务
* 基于TableDataSource类的服务  **√**
* 预定义服务 **√**
* 基于RnmuSource类的服务 **√**
* 基于服务流程的服务
* 服务元数据接口 **√**
* 系统管理服务

## 数据源定义



## 服务定义

SDefine结构提描述一个服务

```go
// SDefine 服务定义结构体
type SDefine struct {
	// ServiceId 服务ID GUID类型
	ServiceId string
	// Context 上下文
	Context string
	// BodyType 报文类型
	BodyType string
	// ServiceType 服务类型
	ServiceType string
	// Meta 服务元数据
	Meta string
	// Namespace 命名空间
	Namespace string
	// Enabled 是否可用
	Enabled bool
	// MsgLog 是否开启消息日志
	MsgLog bool
	// Security 是否开启安全认证
	Security bool
	// 项目id
	ProjectId string
}
```

* Bodytype报文类型目前全部为body，系统里没有对该字段进行处理

* ServiceType包含以下取值

  ``` go
  const (
  	// SrvTypeIds 基于TableDataSource类的服务
  	SrvTypeIds string = "IDS"
  	// SrvTypePredef 预定义服务
  	SrvTypePredef string = "PREDEF"
  	// SrvValueKey value-key形式的服务
  	SrvValueKey string = "VK"
  	// SrvTypeSrvflow 基于服务流程的服务
  	SrvTypeSrvflow string = "SRVFLOW"
  )
  ```

### meta定义

meta定义如何创建服务的实例，每一次请求会创建独立的服务实例提供服务

  ``` json
  {
      "ids":"[服务的数据源id]",
      "userfilter":{
          "filterkey":"[根据数据的哪个字段进行过滤]",
          "opera":"[过滤操作，仅支持=和in]",
          "ids":"[根据哪个数据源的数据进行过滤]",
          "userfield":"[数据源中表示用户id的字段]",
          "joinfield":"[数据源中需要和当前数据源进行匹配的字段]"
      }
  }
  ```

* ids是服务的数据源id，这个id必须在系统中注册，参考数据源定义一节



## 处理接口

### 调用服务

/services/[命名空间].[服务名]/[操作]?_pagesize=200&_pageindex=1&batch_time=2019-11-13

​		服务调用的上下文为service，后面跟服务的URL和操作，URL由一个命名空间和服务名组成，操作包括以下列举的内容：

```go
    //返回全部数据
	SrvAction_ALLDATA string = "all"
	//查询动作
	SrvAction_QUERY string = "query"
	//根据主键返回
	SrvAction_GET string = "get"
	//根据字段值返回
	SrvAction_BYFIELD string = "byfield"
	//返回服务元数据
	SrvAction_META string = "meta"
	//删除操作
	SrvAction_DELETE string = "delete"
	//更新操作
	SrvAction_UPDATE string = "update"
	//插入操作
	SrvAction_INSERT string = "insert"
```

​		上面操作中，query、delete、update、insert三个操作只支持POST方法，其他操作只支持GET方法，在all和query操作后面可以跟限制参数，默认的限制参数包括：

```go
    //以下三个常量均为通过QueryString传入的参数名
	//针对查询自动分页中每页记录数
	REQUEST_PARAM_PAGESIZE string = "_pagesize"
	//针对查询自动分页中的页索引
	REQUEST_PARAM_PAGEINDEX string = "_pageindex"
	//是否返回字段元数据，默认为返回
	REQUEST_PARAM_NOFIELDSINFO string = "_nofield"
	//当前请求不执行而是只返回SQL语句，仅针对IDS类型的服务有效
	REQUEST_PARAM_SQL string = "_sql"
```

### all操作

​	 GET方法，返回全部数据，所有条件都无效，包括聚合和排序，可以使用REQUEST_PARAM_PAGESIZE和REQUEST_PARAM_PAGEINDEX对返回结果进行分页。

### query操作

​	 只支持POST方法。执行查询操作，通过POST提交查询定义报文，同样可以使用REQUEST_PARAM_PAGESIZE和REQUEST_PARAM_PAGEINDEX对返回结果进行分页。

​	报文格式，下面#后面为注释：

```json
{
  "Criteria": [	#查询条件，数组类型，每一个元素为一个条件
    {
      "field": "batch_time",#字段名
      "operation": "=",	#操作，支持=  !=  >  <  >=  <=  in
      "value": "2019-11-13",#数值，时间数值采用yyyy-mm-dd hh24:mi:ss的格式
      "relation": "and"#与前面一个条件的逻辑关系，执行and or，Critical中的第一个条件relation属性无意义
    }
  ],
  "orderby":"tm desc,stcd asc"#排序字段，逗号分割，每一个排序属性为字段名+空格+desc|asc

  "PostAction":[
  	{
  		"name":"slice", #数据行列转换
  		"params":{
  			"xfield":"item_id",   #输出结果集的字段，即将数据集的item_id字段的值作为新数据集的列
  			"yfield":["dev_id", "site_id","collect_date"],#输出结果集的y轴字段
  			"valuefield":"data_value"  #输出结果集的数值字段
  		}
  	},{
  		"name":"fieldmeta", #添加字段元数据
  		"params":{"metaurl":"idb.table.iotdata"} #元数据的URL
  	},{
     		"name":"fieldgroup",
     		"params":{
     			"field":"site_id"
     		}
     	},{
  		"name":"bulldozer",
  		"params":{
  			"bulldozer":[
			  	{
                  "name": "FormatDatafunc",     #字段格式化
                  "params": {"collect_date": "2006-01-02 15:04"}  #格式化字段，针对时间字段格式化处理
                }, {
                	"name": "DictMappingfunc",  #字典映射
                	"params": {
                			"outfield": "mon_site", #映射后输出字段
                			"dataKeyField": "site_id",   #映射的主键字段
                			"KeyStringSourceName": "idb.site"   #映射的数据源，必须是KeyStringSource
                		}
                },
                     {
                       "name": "ColumnFilterFunc",   #字段过滤器,字段过滤器必须为最后一个处理函数
                       "params": {
                         "show": ["USER_ID", "USER_NAME", "ORG_ID", "USER_CREATED","ORG_NAME"]
                       }
                     }
			  	]}
  	}]
}
```



> **特殊处理时间类型的参数，当条件的属性类型为时间时可以使用特殊字符串表示特定的时间，包括：**
>
> 如前N天，lastday:1    lastday:-3
> addmonth,addmonth:1
> addyear
> now
> thismonth
> thisyear
> today





>  在当前请求的URL中添加参数**_cache**，可以在返回结果的同时将当前结果集保存在服务器端的缓存中，参数 格式为：30_1   其中30表示缓存30s，1表示最多被返回一次。





### get操作



### byfield操作



### meta操作



### delete操作



### update操作



### insert操作

### 	

## 安全机制

## 元数据支持

## 可视化服务

### 数据表格

### 基础统计图


## 请求实例
* 请求实例
```json
{
  "Criteria": [
    {
      "field": "batch_time",
      "operation": "=",
      "value": "2019-11-13",
      "relation": "and"
    }
  ],

  "PostAction":[
  	{
  		"name":"slice",
  		"params":{
  			"xfield":"item_id",
  			"yfield":["dev_id", "site_id","collect_date"],
  			"valuefield":"data_value"
  		}
  	},{
  		"name":"fieldmeta",
  		"params":{"metaurl":"idb.table.iotdata"}
  	},
  	{
  		"name":"bulldozer",
  		"params":{
  			"bulldozer":[
			  	{
			    	"name":"FormatDatafunc",
			    	"params":{
			    		"collect_date": "2006-01-02 15:04"
			    	}
			    }]}
  	}]
}
```
* 请求实例
```json
{
  "Criteria": [
    {
      "field": "ORG_ID",
      "operation": "=",
      "value": "001013009",
      "relation": "and"
    }
  ],
  "bulldozer": [
  	{
  		"name":"DictMappingfunc",
  		"params":{
  			"outfield":        "ORG_NAME",
			"dataKeyField":    "ORG_ID",
			"KeyStringSourceName":"ORG_NAME"
  		}
  	},{
    	"name":"FormatDatafunc",
    	"params":{
    		"USER_CREATED": "2006-01-02"
    	}
    },
    {
      "name": "ColumnFilterFunc",
      "params": {
        "show": ["USER_ID", "USER_NAME", "ORG_ID", "USER_CREATED","ORG_NAME"]
      }
    }
  ]
}
```
* 请求实例
```json
{
	"Criteria": [{
		"field": "collect_date",
		"operation": ">",
		"value": "addmonth:-2",
		"relation": "and"
	}],
	"PostAction": [{
		"name": "slice",
		"params": {
			"xfield": "item_id",
			"yfield": ["dev_id", "site_id", "collect_date"],
			"valuefield": "data_value"
		}
	}, {
		"name": "fieldmeta",
		"params": {
			"metaurl": "idb.table.iotdata"
		}
	}, {
		"name": "bulldozer",
		"params": {
			"bulldozer": [{
				"name": "FormatDatafunc",
				"params": {
					"collect_date": "2006-01-02 15:04"
				}
			}, {
				"name": "DictMappingfunc",
				"params": {
					"outfield": "mon_site",
					"dataKeyField": "site_id",
					"KeyStringSourceName": "idb.site"
				}
			}, {
				"name": "DictMappingfunc",
				"params": {
					"outfield": "site_name",
					"dataKeyField": "dev_id",
					"KeyStringSourceName": "idb.dev"
				}
			}]
		}
	},{
		"name":"fieldgroup",
		"params":{
			"field":"site_id"
		}
	}]
}
```

* Insert 
```json
{
  "insert":{
		"ID":        "6bdfbad3-20ab-4610-8dcb-80413a09ef55",
		"DBTYPE":    "mysql22222",
		"DBURL":     "{username}:{password}@tcp(127.0.0.1:3306)/pest",
		"USERNAME":  "tong",
		"PWD":       "123456",
		"PROJECTID": "",
		"DBALIAS":   "pest"
		}
}
```

* update
```json
{
  "update":{
		"ID":        "6bdfbad3-20ab-4610-8dcb-80413a09ef55",
		"DBTYPE":    "mysql33333",
		"DBURL":     "{username}:{password}@tcp(127.0.0.1:3306)/pest",
		"USERNAME":  "tong",
		"PWD":       "123456",
		"PROJECTID": "",
		"DBALIAS":   "pest"
		},
   "Criteria": [
    {
      "field": "ID",
      "operation": "=",
      "value": "6bdfbad3-20ab-4610-8dcb-80413a09ef55",
      "relation": "and"
    }
  ]
}
```
* delete
```json
{
  "delete":"true",
   "Criteria": [
    {
      "field": "ID",
      "operation": "=",
      "value": "6bdfbad3-20ab-4610-8dcb-80413a09ef55",
      "relation": "and"
    }
  ]
}
```

* 在响应的Fields部分返回字段的说明文字，metaurl参数按照项目、命名空间、元数据民称的顺序标示元数据
```json
{
	"PostAction":[
		 {
          "name": "fieldmeta",
          "params": {
            "metaurl": "jeda.table.meta"
          }
        }
    ]
}
```

* 默认响应是数据组形式，通过在URL中添加_repstyle=map参数设定响应形式为map方式

* 通过服务定义可以为该服务添加用户过滤器，只返回和某一用户相关的数据
服务定义JSON：
```go
    JedaSrvContainer["jeda.meta"] = createService("jeda.meta", map[string]interface{}{
		"ids": "default.mgr.G_META",
		"userfilter": map[string]string{
			"filterkey": "PROJECTID",
			"opera":     "in",
			"ids":       "default.mgr.G_USERPROJECT",
			"userfield": "USERID",
			"joinfield": "PROJECTNAME"},
	})
```


* 递归用户过滤器
服务定义JSON：
```json
{
		"ids": "default.mgr.G_META",
		"userfilter": {
			"filterkey": "PROJECTID",
			"opera": "in",
			"values":{
				"outfield": "PROJECTNAME"
				"ids": "default.mgr.G_USERPROJECT",
				"userfilter": {
		 			"filterkey": "PROJECTID",
					"opera": "in",
					"userfield": "USERID",
				}  				
			}
		}
}
```

* 服务定义中添加JOIN子句
```json
{
		"ids": "default.mgr.G_META",
		"join": {}
}
```