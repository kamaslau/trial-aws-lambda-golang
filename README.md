# README.md

以下官方文档仅做参考，文章及示例代码可能基于较为久远版本的 Golang。

https://docs.aws.amazon.com/zh_cn/lambda/latest/dg/lambda-golang.html

为了发布为 AWS Lambda (运行环境以 Amazon Linux 2023/arm64 为例)，编译后的入口文件名需为 `bootstap`，此处以将编译后的可执行文件放入 `/dist` 目录为例：

```bash
go get

#
# https://dev.to/calebfrieze/deploying-go-to-aws-using-cdk-in-2024-41ah
GOOS=linux GOARCH=arm64 go build -o dist/bootstrap
```

编译后即可将可执行文件压缩为 zip 文件并通过 AWS Lambda 控制面板上传，或其它方式进行部署。
