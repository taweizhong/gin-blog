# gin-blog 基于gin实现的个人博客

## 简介

使用gin搭建web框架，实现对文章、文章标签和作者的增删改查。

## 文件目录

- conf：用于存储配置文件；
- docs：文档生成；
- middleware：应用中间件；
- models：应用数据库模型；
- pkg：第三方包；
- routers 路由逻辑处理；
- runtime 应用运行时数据；

## 功能实现

- ini库实现配置文档读取
- jwt实现token
- Swagger实现接口文档
- Redis 缓存
- gorm实现对mysql数据库的操作
- 实现一键式docker部署
- air工具实现项目热更新
