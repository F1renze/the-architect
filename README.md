# The Architect


![bkuqmxwswsxgrr8(1).jpg](https://i.loli.net/2020/02/12/Q3CTJkB26XwVdHF.jpg)
> Inspired by [The Matrix](https://en.wikipedia.org/wiki/Architect_(The_Matrix)).

Architect 是一个基于微服务的认证及访问权限控制方案，支持多种登录方式并提供基于域的角色访问控制（RBAC with Domains）。

# 服务架构

![image.png](https://i.loli.net/2020/02/12/yHq75RjnavukepZ.png)


# 项目结构

```

├── Makefile
├── README.md
├── api
│   └── user 
├── common // 公共库
│   └── db.sql // 表结构
├── infra
│   ├── config // 配置中心
│   └── gateway // api 网关
├── srv
│   ├── auth 
│   ├── casbin 
│   └── user
└── test
```

# 快速开始
```
git clone https://github.com/F1renze/the-architect.git

cd the-architect
// 配置 mysql 及 redis
vim infra/config/config.yml
// 配置腾讯云短信服务依赖环境变量，参考 .env_example
vim .env

// 启动容器
make dev
```
  
