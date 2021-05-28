# x-lab后端文档

## Quick start

1. 进入service目录下，建立数据库与表
```shell
cd service && bash scripts/database.sh
```

2. 在service目录下，启动后端
```shell
go run main.go
```

## 项目设计

### 目录结构

- configs: 配置文件
- global: 全局变量
- internal: 内部模块
    - dao: 数据访问层，具体操作数据库的地方
    - middleware: Http中间层
    - model: 模型层，用户存放model文件
    - routers: 路由相关的逻辑
    - service: 项目核心业务逻辑，每次操作时新建service对象。即：router->service->dao->model
- pkg: 项目相关的模块包
- storage: 项目文件路径
- scripts: 各类构建、安装、分析等操作脚本

## 进度

- [ ] 用户登陆，增加手机号码和邮箱
  
- [x] 文件上传 POST

- [ ] 获取已经上传文件

