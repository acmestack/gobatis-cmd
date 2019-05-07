# gobatis-cmd

## 安装

使用命令安装：

```
go get github.com/xfali/gobatis-cmd
```

## 使用
```
gobatis-cmd -host localhost -port 3306 -user test -pw test -db test_db -pkg test_package -path .
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
        生成文件的包名 (default "xfali/gobatis/default")
  -port int
        数据库的端口号  (default 3306)
  -pw string
        数据库的密码
  -table string
        指定生成的table名称
  -tag string
        生成Model的tag名称 (default "xfield")
  -user string
        数据库的用户名
```

会在当前目录下生成1个目录及3个文件，分别为：

### 1、Model文件

提取数据库表信息生成对应的供gobatis使用的Model文件：

文件为： ${PATH}/${表名}.go

自动增加配置的包名，如果包含时间字段，自动import time包。

例子：
```
package test_package

import "time"

type TestTable struct {
    //TableName gobatis.ModelName `test_table`
    Id int64 `xfield:"id"`
    Username string `xfield:"username"`
    Password string `xfield:"password"`
    UpdateTime time.Time `xfield:"update_time"`
}

```

### 2、xml文件

自动生成包含select、insert、update、delete映射xml文件

文件为： ${PATH}/xml/${表名}.xml

（请根据实际业务自行修改）

例子：
```
<mapper namespace="test_package.TestTable">
    <sql id="columns_id">id,username,password,update_time</sql>

    <select id="selectTestTable">
        SELECT <include refid="columns_id"> </include> FROM test_table
        <where>
            <if test="id != -1">AND id = #{id} </if>
            <if test="username != nil">AND username = #{username} </if>
            <if test="password != nil">AND password = #{password} </if>
            <if test="update_time != nil">AND update_time = #{update_time} </if>
        </where>
    </select>

    <insert id="insertTestTable">
        INSERT INTO test_table (id,username,password,update_time)
        VALUES(
        #{id},
        #{username},
        #{password},
        #{update_time}
        )
    </insert>

    <update id="updateTestTable">
        UPDATE test_table
        <set>
            <if test="id != -1"> id = #{id} </if>
            <if test="username != nil"> username = #{username} </if>
            <if test="password != nil"> password = #{password} </if>
            <if test="update_time != nil"> update_time = #{update_time} </if>
        </set>
        WHERE id = #{id}
    </update>

    <delete id="deleteTestTable">
        DELETE FROM test_table
        <where>
            <if test="id != -1">AND id = #{id} </if>
            <if test="username != nil">AND username = #{username} </if>
            <if test="password != nil">AND password = #{password} </if>
            <if test="update_time != nil">AND update_time = #{update_time} </if>
        </where>
    </delete>
</mapper>
```

### 3、代理

自动根据model和xml生成代理方法，包含：

1. package包名
2. import包
3. init方法（初始化model、初始化xml，请根据实际业务自行修改）
4. New方法：使用SessionManager获得Proxy，见[gobatis](https://github.com/xfali/gobatis)
5. 事务方法Tx
6. 与xml相匹配的代理方法

例子：
```
package test_package

import (
    "github.com/xfali/gobatis/config"
    "github.com/xfali/gobatis/session/runner"
)

type TestTableCallProxy runner.RunnerSession

func init() {
    modelV := TestTable{}
    config.RegisterModel(&modelV)
    config.RegisterMapperFile("e:/tmp/xml/test_table.xml")
}

func New(proxyMrg *runner.SessionManager) *TestTableCallProxy {
    return (*TestTableCallProxy)(proxyMrg.NewSession())
}

func (proxy *TestTableCallProxy) Tx(txFunc func(s *TestTableCallProxy) bool) {
    sess := (*runner.RunnerSession)(proxy)
    sess.Tx(func(session *runner.RunnerSession) bool {
        return txFunc(proxy)
    })
}

func (proxy *TestTableCallProxy)SelectTestTable(model TestTable) []TestTable {
    var dataList []TestTable
    (*runner.RunnerSession)(proxy).Select("selectTestTable").Param(model).Result(&dataList)
    return dataList
}

func (proxy *TestTableCallProxy)InsertTestTable(model TestTable) int64 {
    var ret int64
    (*runner.RunnerSession)(proxy).Insert("insertTestTable").Param(model).Result(&ret)
    return ret
}

func (proxy *TestTableCallProxy)UpdateTestTable(model TestTable) int64 {
    var ret int64
    (*runner.RunnerSession)(proxy).Update("updateTestTable").Param(model).Result(&ret)
    return ret
}

func (proxy *TestTableCallProxy)DeleteTestTable(model TestTable) int64 {
    var ret int64
    (*runner.RunnerSession)(proxy).Delete("deleteTestTable").Param(model).Result(&ret)
    return ret
}
```

### 文件使用

1. 将文件拷贝到工程目录
2. 需要自己初始化gobatis的SessionManager
3. 使用SessionManager通过New方法获得代理
4. 使用代理调用数据库

例子：
```
fac := factory.DefaultFactory{
    Host:     "localhost",
    Port:     3306,
    DBName:   "test",
    Username: "root",
    Password: "123",
    Charset:  "utf8",

    MaxConn:     1000,
    MaxIdleConn: 500,

    Log: logging.DefaultLogf,
}
fac.Init()
sessionMgr := runner.NewSessionManager(&fac)

proxy := New(sessionMgr)
ret := proxy.InsertTestTable(TestTable{Username:"test_user"})

fmt.Println(ret)

//事务
proxy.Tx(func(s *TestTableCallProxy) bool {
    .UpdateTestTable(TestTable{Id: 1, Username:"user"})
    return true
})

```

## 其他

用户可自行修改xml、model、proxy适配自己的业务。也可以不依赖proxy，直接使用gobatis的方法使用model和xml。

建议参考proxy的做法。