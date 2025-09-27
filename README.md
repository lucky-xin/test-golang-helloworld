# Test Go Helloworld - 国密加密算法服务

这是一个基于Go语言开发的国密加密算法服务，实现了SM2、SM3、SM4等国产密码算法，提供HTTP API接口服务。

## 项目功能

- **SM2非对称加密算法**：支持公钥加密、私钥解密，兼容Java BouncyCastle生成的密钥
- **SM3哈希算法**：提供数据摘要功能
- **SM4对称加密算法**：支持CBC模式加密解密
- **HTTP API服务**：基于Gin框架提供RESTful接口
- **Docker容器化**：支持Docker部署
- **Kubernetes部署**：提供K8s部署模板

## 目录结构

```
.
├── config/                                         项目配置目录
│   └── config.go                                  配置结构体和初始化
├── encryption/                                     国密加密算法实现
│   ├── sm2.go                                      SM2非对称加密算法
│   ├── sm3.go                                      SM3哈希算法
│   └── sm4.go                                      SM4对称加密算法
├── routers/                                        路由配置
│   └── routers.go                                 路由初始化和API定义
├── test/                                           测试文件
│   ├── sm2_test.go                                 SM2算法测试
│   ├── sm3_test.go                                 SM3算法测试
│   └── sm4_test.go                                 SM4算法测试
├── deploy/                                         部署相关文件
│   └── deployment.tpl                              Kubernetes部署模板
├── main.go                                         程序入口
├── go.mod                                          Go模块依赖
├── go.sum                                          依赖校验和
├── Dockerfile                                      容器镜像构建文件
├── Jenkinsfile                                     Jenkins CI/CD流水线
└── sonar-scanner-cli.sh                           SonarQube代码质量扫描脚本
```

## 快速开始

### 环境要求

- Go 1.25+
- Docker (可选)
- Kubernetes (可选)

### 本地运行

```bash
# 克隆项目
git clone <repository-url>
cd test-golang-helloworld

# 安装依赖
go mod tidy

# 运行项目
go run main.go

# 或者构建后运行
go build -o main
./main
```

### Docker运行

```bash
# 构建镜像
docker build -t test-golang-helloworld .

# 运行容器
docker run -p 8000:8000 test-golang-helloworld
```



## API接口

### 健康检查
- **GET** `/ping` - 服务健康检查
  - 响应：`{"return": "pong"}`

## 加密算法使用示例

### SM2非对称加密

```go
package main

import (
    "fmt"
    "xyz/test/helloworld/encryption"
)

func main() {
    // 生成密钥对（实际使用中应该从安全的地方获取）
    publicKeyHex := "04..."
    privateKeyHex := "..."

    // 创建SM2实例
    sm2, err := encryption.NewSM2(publicKeyHex, privateKeyHex)
    if err != nil {
        panic(err)
    }

    // 加密
    plaintext := "Hello, SM2!"
    ciphertext, err := sm2.Encrypt2Hex(plaintext, 0) // 0=C1C3C2模式
    if err != nil {
        panic(err)
    }

    // 解密
    decrypted, err := sm2.DecryptHex(ciphertext, 0)
    if err != nil {
        panic(err)
    }

    fmt.Printf("原文: %s\n", plaintext)
    fmt.Printf("密文: %s\n", ciphertext)
    fmt.Printf("解密: %s\n", string(decrypted))
}
```

### SM3哈希算法

```go
package main

import (
    "encoding/hex"
    "fmt"
    "xyz/test/helloworld/encryption"
)

func main() {
    data := "Hello, SM3!"
    hash := encryption.EncodeToSM3(data)
    hashHex := hex.EncodeToString(hash)
    
    fmt.Printf("原文: %s\n", data)
    fmt.Printf("SM3哈希: %s\n", hashHex)
}
```

### SM4对称加密

```go
package main

import (
    "encoding/hex"
    "fmt"
    "xyz/test/helloworld/encryption"
)

func main() {
    // 16字节密钥和IV
    keyHex := "0123456789ABCDEFFEDCBA9876543210"
    ivHex := "00000000000000000000000000000000"

    // 创建SM4实例
    sm4, err := encryption.FromHex(keyHex, ivHex)
    if err != nil {
        panic(err)
    }

    // 加密
    plaintext := "Hello, SM4!"
    ciphertext, err := sm4.Encrypt2Hex(plaintext)
    if err != nil {
        panic(err)
    }

    // 解密
    decrypted, err := sm4.DecryptHex(ciphertext)
    if err != nil {
        panic(err)
    }

    fmt.Printf("原文: %s\n", plaintext)
    fmt.Printf("密文: %s\n", ciphertext)
    fmt.Printf("解密: %s\n", string(decrypted))
}
```

## 配置说明

| 配置项 | 描述 | 默认值 |
|--------|------|--------|
| `Port` | 服务监听端口 | 8000 |
| `Auth` | 是否启用认证 | false |
| `EtcdEndPoints` | Etcd服务地址 | ["192.168.31.5:2379"] |

## 部署说明

### Kubernetes部署

使用提供的deployment模板进行K8s部署：

```bash
# 替换模板变量
sed 's/{APP_NAME}/test-golang-helloworld/g; s/{NAMESPACE}/default/g; s/{VERSION}/1.0.0/g' deploy/deployment.tpl > deployment.yaml

# 部署到Kubernetes
kubectl apply -f deployment.yaml
```

### Docker部署

```bash
# 构建镜像
docker build -t test-golang-helloworld:latest .

# 运行容器
docker run -d -p 8000:8000 --name helloworld test-golang-helloworld:latest
```

## 测试

运行测试套件：

```bash
# 运行所有测试
go test ./...

# 运行特定测试
go test ./test/sm2_test.go
go test ./test/sm3_test.go
go test ./test/sm4_test.go

# 生成测试覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 代码质量检查

使用SonarQube进行代码质量扫描：

```bash
# 运行SonarQube扫描
./sonar-scanner-cli.sh --host-url http://your-sonar-server:9000 --login your-token
```
