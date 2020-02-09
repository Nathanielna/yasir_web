# yasir_web
自用基于Beego框架的速搭环境

自带MySQL 仅需在conf里面配置Mysql信息即可
Model统一管理
日志全部写入log目录下
404、403等错误页面处理
自动路由，详情Beego官网路由

## 首次使用
自行搭建Go环境和Beego框架  
请创建表 并且在Model里 配置相应的模型  
可使用beego官方命令快速建立模型，控制器  
```
bee generate appcode -tables="ier_resoure_explain_rel" -conn="root:yxp@tcp(127.0.0.1:3306)/ireport"
```

😊本人博客：[小小咪的博客](http://12yxp.top/)  
💖IT教程站：[教程搬运](https://12byg.com/)