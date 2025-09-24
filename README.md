# Test Go Helloworld

## 目录结构

```
.
├── api                                             存放接口的目录
│   └── home.go
├── cache                                           缓冲相关的目录
│   ├── cache.go
│   └── redis_cache.go
├── config                                          项目配置地址
│   ├── config.go
│   └── config.toml
├── global                                          全局变量redis、Mongodb连接
│   └── global.go
├── go.mod
├── go.sum
├── Dockerfile                                      镜像构建文件
├── logs                                            存放日志目录部署需要配置为其他目录
│   ├── system.log -> system.log.20210606.log
│   ├── system.log.20210605.log
│   └── system.log.20210606.log
├── middleware                                      中间件目录
│   ├── core_middle.go
│   └── log_middle.go
├── models                                          实体映射
│   └── home.go
├── repository                                      实体针对数据操作
│   └── home_repository.go
├── router.go                                       路由配置
├── server.go                                       启动配置
├── server_other.go                                 非win系统启动配置
├── server_win.go                                   win系统启动配置
├── service                                         业务操作
│   └── home_serivce.go
├── utils                                           工具目录，消息，工具方法
│   ├── message
│   └── tools
│       └── type_utils.go
└── view                                            返回或者接受的实体
    └── home_view.go
```

创建目录结构：

```
# 创建目录结构
New-Item -ItemType Directory -Path 'api'
New-Item -ItemType Directory -Path 'cache'
New-Item -ItemType Directory -Path 'config'
New-Item -ItemType Directory -Path 'global'
New-Item -ItemType Directory -Path 'logs'
New-Item -ItemType Directory -Path 'middleware'
New-Item -ItemType Directory -Path 'models'
New-Item -ItemType Directory -Path 'repository'
New-Item -ItemType Directory -Path 'service'
New-Item -ItemType Directory -Path 'utils\message'
New-Item -ItemType Directory -Path 'utils\tools'
New-Item -ItemType Directory -Path 'view' 

# 创建文件
New-Item -ItemType File -Path 'api\home.go'
New-Item -ItemType File -Path 'config\config.go'
New-Item -ItemType File -Path 'models\home.go'
New-Item -ItemType File -Path 'router.go'
New-Item -ItemType File -Path 'server.go'
```



## 访问地址
http://127.0.0.1:8000/ping

## 环境变量配置

| 环境变量名称        | 描述    | 默认值  |
|---------------|-------|------|
| `SERVER_PORT` | 服务器端口 | 3003 |
