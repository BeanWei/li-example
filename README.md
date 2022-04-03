# li-example
Li 框架演示代码 https://github.com/BeanWei/li

# 步骤
- 1: 修改 `config/config.yaml` 中的数据库配置
- 2: 执行 `make generate` 生成代码
- 3: 执行 `make migrate` 迁移数据库模型
- 4: 执行 `go run li.go user-create -u=你的邮箱 -p=你的密码 -a=true` 创建初始化账号
- 5: 执行 `make run` 启动应用
- 6: 浏览器打开 `localhost:8199` 开始体验