# blog
blog是基于Golang的个人博客系统，具备对博文的发布、浏览、修改及删除功能。

项目有以下特性：

* 基于[Gin](https://github.com/gin-gonic/gin)框架，实现RESTful风格
* 基于MySQL，实现对内容的持久化存储
* 基于Redis，实现对博文的缓存及阅读量的统计