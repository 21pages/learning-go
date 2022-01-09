# 高版本要指定go_package
```
protoc-gen-go: unable to determine Go import path for "pubsub.proto"

Please specify either:
        • a "go_package" option in the .proto source file, or
        • a "M" argument on the command line.
```
解决方法:
    1. 降低protoc-gen-go到1.3.2以下
    2. 添加`go_package="./xxx"`