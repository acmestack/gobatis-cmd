# gobatis-cmd


[![license card](https://img.shields.io/badge/License-Apache%202.0-brightgreen.svg?label=license)](https://github.com/acmestack/gobatis-cmd/blob/main/LICENSE)
[![go version](https://img.shields.io/github/go-mod/go-version/acmestack/gobatis-cmd)](#)
[![go report](https://goreportcard.com/badge/github.com/acmestack/gobatis-cmd)](https://goreportcard.com/report/github.com/acmestack/gobatis-cmd)
[![codecov report](https://codecov.io/gh/acmestack/gobatis-cmd/branch/main/graph/badge.svg)](https://codecov.io/gh/acmestack/gobatis-cmd)
[![workflow](https://github.com/acmestack/gobatis-cmd/actions/workflows/go.yml/badge.svg?event=push)](#)
[![lasted release](https://img.shields.io/github/v/release/acmestack/gobatis-cmd?label=lasted)](https://github.com/acmestack/gobatis-cmd/releases)

## 安装

使用命令安装：

```
go get github.com/acmestack/gobatis-cmd/cmd/gobatis-cmd
```

## 使用

```
gobatis-cmd -f configs/gobatis-conf.json
```
```
{
  "driver": "mysql",
  "path": ".",
  "package": "test",
  "namespace": "test",
  "modelFile": "test_model.go",
  "tagName": "column",
  "mapperFile": "xml",
  "plugin": "",
  "keyword": false,
  "tableName": "",
  "dbName": "test_db",
  "host": "localhost",
  "port": 3306,
  "user": "",
  "password": ""
}
```
或者使用命令参数：
```
gobatis-cmd -driver=mysql -host=localhost -port=3306 -user=test -pw=test -db=testdb -pkg=test_package -mapper=xml -namespace=test -path=.
```
```
  -db string
        指定解析的数据库名称
  -driver string
        Driver (default "mysql")
  -host string
        数据库地址 (default "localhost")
  -model string
        生成的model文件名称 (default "models.go")
  -path string
        保存生成文件的路径
  -pkg string
        生成文件的包名 (default "acmestack/gobatis/default")
  -port int
        数据库的端口号  (default 3306)
  -pw string
        数据库的密码
  -table string
        指定生成的table名称
  -tag string
        生成Model的tag名称,多tag用逗号分隔，如"json,xml" (default "column")
  -user string
        数据库的用户名
  -mapper string
        mapper文件类型: xml | template | go （默认xml）
  -keyword bool
        是否自动添加转义符，默认false，如果为true则会根据driver名称辨识添加
  -namespace string
        添加namespace，避免同表名冲突，（默认为空，如果为空，自动填充为packageName.ModelName）
```

会在当前目录下生成1个目录及3个文件，分别为：

### 1、Model文件

提取数据库表信息生成对应的供gobatis使用的Model文件：

文件为： ${PATH}/${表名}.go

自动增加配置的包名，如果包含时间字段，自动import time包。

例子：
```
package test

import (
	"github.com/acmestack/gobatis"
	"time"
)

type TestTable struct {
	//TableName gobatis.TableName `test_table`
	Id         int       `column:"id"`
	Username   string    `column:"username"`
	Password   string    `column:"password"`
	Createtime time.Time `column:"createtime"`
}

func (m *TestTable) Select(sess *gobatis.Session) ([]TestTable, error) {
	return SelectTestTable(sess, *m)
}

func (m *TestTable) Count(sess *gobatis.Session) (int64, error) {
	return SelectTestTableCount(sess, *m)
}

func (m *TestTable) Insert(sess *gobatis.Session) (int64, int64, error) {
	return InsertTestTable(sess, *m)
}

func (m *TestTable) Update(sess *gobatis.Session) (int64, error) {
	return UpdateTestTable(sess, *m)
}

func (m *TestTable) Delete(sess *gobatis.Session) (int64, error) {
	return DeleteTestTable(sess, *m)
}
```

### 2、xml文件

自动生成包含select、insert、update、delete映射xml文件

文件为： ${PATH}/xml/${表名}_mapper.xml

（请根据实际业务自行修改）

例子：
```
<mapper namespace="test">
    <sql id="columns_id">id,username,password,createtime</sql>

    <select id="selectTestTable">
        SELECT <include refid="columns_id"> </include> FROM test_table
        <where>
            <if test="{TestTable.id} != nil and {TestTable.id} != 0">AND id = #{TestTable.id} </if>
            <if test="{TestTable.username} != nil">AND username = #{TestTable.username} </if>
            <if test="{TestTable.password} != nil">AND password = #{TestTable.password} </if>
            <if test="{TestTable.createtime} != nil">AND createtime = #{TestTable.createtime} </if>
        </where>
    </select>

    <select id="selectTestTableCount">
        SELECT COUNT(*) FROM test_table
        <where>
            <if test="{TestTable.id} != nil and {TestTable.id} != 0">AND id = #{TestTable.id} </if>
            <if test="{TestTable.username} != nil">AND username = #{TestTable.username} </if>
            <if test="{TestTable.password} != nil">AND password = #{TestTable.password} </if>
            <if test="{TestTable.createtime} != nil">AND createtime = #{TestTable.createtime} </if>
        </where>
    </select>

    <insert id="insertTestTable">
        INSERT INTO test_table (id,username,password,createtime)
        VALUES(
        #{TestTable.id},
        #{TestTable.username},
        #{TestTable.password},
        #{TestTable.createtime}
        )
    </insert>

    <insert id="insertBatchTestTable">
        INSERT INTO test_table (id,username,password,createtime)
        VALUES
        <foreach item="item" index="index" collection="{0}" open="" separator="," close="">
            (#{item.TestTable.id},#{item.TestTable.username},#{item.TestTable.password},#{item.TestTable.createtime})
        </foreach>
    </insert>

    <update id="updateTestTable">
        UPDATE test_table
        <set>
            <if test="{TestTable.username} != nil"> username = #{TestTable.username} </if>
            <if test="{TestTable.password} != nil"> password = #{TestTable.password} </if>
            <if test="{TestTable.createtime} != nil"> createtime = #{TestTable.createtime} </if>
        </set>
        WHERE id = #{TestTable.id}
    </update>

    <delete id="deleteTestTable">
        DELETE FROM test_table
        <where>
            <if test="{TestTable.id} != nil and {TestTable.id} != 0">AND id = #{TestTable.id} </if>
            <if test="{TestTable.username} != nil">AND username = #{TestTable.username} </if>
            <if test="{TestTable.password} != nil">AND password = #{TestTable.password} </if>
            <if test="{TestTable.createtime} != nil">AND createtime = #{TestTable.createtime} </if>
        </where>
    </delete>
</mapper>
```

### 3、代理（目前已修改为操作方法）

文件为： ${PATH}/${表名}_proxy.go

自动根据model和xml生成代理方法，包含：

1. package包名
2. import包
3. init方法（初始化model、初始化xml，请根据实际业务自行修改）
4. 与xml相匹配的代理函数

例子：
```
package test

import (
    "github.com/acmestack/gobatis"
)

func init() {
    modelV := TestTable{}
    gobatis.RegisterModel(&modelV)
    gobatis.RegisterTemplateFile("./template/test_table_mapper.tmpl")
}

func SelectTestTable(sess *gobatis.Session, model TestTable) ([]TestTable, error) {
    var dataList []TestTable
    err := sess.Select("test.selectTestTable").Param(model).Result(&dataList)
    return dataList, err
}

func SelectTestTableCount(sess *gobatis.Session, model TestTable) (int64, error) {
    var ret int64
    err := sess.Select("test.selectTestTableCount").Param(model).Result(&ret)
    return ret, err
}

func InsertTestTable(sess *gobatis.Session, model TestTable) (int64, int64, error) {
    var ret int64
    runner := sess.Insert("test.insertTestTable").Param(model)
    err := runner.Result(&ret)
    id := runner.LastInsertId()
    return ret, id, err
}

func InsertBatchTestTable(sess *gobatis.Session, models []TestTable) (int64, int64, error) {
    var ret int64
    runner := sess.Insert("test.insertBatchTestTable").Param(models)
    err := runner.Result(&ret)
    id := runner.LastInsertId()
    return ret, id, err
}

func UpdateTestTable(sess *gobatis.Session, model TestTable) (int64, error) {
    var ret int64
    err := sess.Update("test.updateTestTable").Param(model).Result(&ret)
    return ret, err
}

func DeleteTestTable(sess *gobatis.Session, model TestTable) (int64, error) {
    var ret int64
    err := sess.Delete("test.deleteTestTable").Param(model).Result(&ret)
    return ret, err
}
```
### template

当参数mapper=template时会生成go template文件，文件为： ${PATH}/template/${表名}_mapper.tmpl

例子：
```cassandraql
{{define "namespace"}}test{{end}}

{{define "selectTestTable"}}
SELECT id,username,password,createtime FROM test_table
{{where .Id "AND" "id = " (arg .Id) "" | where .Username "AND" "username = " (arg .Username) | where .Password "AND" "password = " (arg .Password) | where .Createtime "AND" "createtime = " (arg .Createtime)}}
{{end}}

{{define "selectTestTableCount"}}
SELECT COUNT(*) FROM test_table
{{where .Id "AND" "id = " (arg .Id) "" | where .Username "AND" "username = " (arg .Username) | where .Password "AND" "password = " (arg .Password) | where .Createtime "AND" "createtime = " (arg .Createtime)}}
{{end}}

{{define "insertTestTable"}}
INSERT INTO test_table(id,username,password,createtime)
VALUES(
{{arg .Id}}, {{arg .Username}}, {{arg .Password}}, {{arg .Createtime}})
{{end}}

{{define "insertBatchTestTable"}}
{{$size := len . | add -1}}
INSERT INTO test_table(id,username,password,createtime)
VALUES {{range $i, $v := .}}
({{arg $v.Id}}, {{arg $v.Username}}, {{arg $v.Password}}, {{arg $v.Createtime}}){{if lt $i $size}},{{end}}
{{end}}
{{end}}

{{define "updateTestTable"}}
UPDATE test_table
{{set .Id "id = " (arg .Id) "" | set .Username "username = " (arg .Username) | set .Password "password = " (arg .Password) | set .Createtime "createtime = " (arg .Createtime)}}
{{where .Id "AND" "id = " (arg .Id) ""}}
{{end}}

{{define "deleteTestTable"}}
DELETE FROM test_table
{{where .Id "AND" "id = " (arg .Id) "" | where .Username "AND" "username = " (arg .Username) | where .Password "AND" "password = " (arg .Password) | where .Createtime "AND" "createtime = " (arg .Createtime)}}
{{end}}
```

### 文件使用

1. 将文件拷贝到工程目录
2. 需要自己初始化gobatis的SessionManager
3. 使用SessionManager通过New方法获得代理
4. 使用代理调用数据库

例子：
```
fac := gobatis.NewFactory(
    		gobatis.SetMaxConn(100),
    		gobatis.SetMaxIdleConn(50),
    		gobatis.SetDataSource(&datasource.MysqlDataSource{
    			Host:     "localhost",
    			Port:     3306,
    			DBName:   "test",
    			Username: "root",
    			Password: "123",
    			Charset:  "utf8",
    		}))
sessionMgr := gobatis.NewSessionManager(&fac)

sess := sessionMgr.NewSession(sessionMgr)
ret, insertId := InsertTestTable(sess, TestTable{Username:"test_user"})

fmt.Println(ret)
```
事务:

使用gobatis的session.Tx() 参考[gobatis](https://github.com/acmestack/gobatis)

## 其他

用户可自行修改xml、model、proxy适配自己的业务。也可以不依赖proxy，直接使用gobatis的方法使用model和xml。

建议参考proxy的做法。