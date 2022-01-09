# 简介
protobuf是序列化工具, 将结构体,函数等序列化为字节, 使不同语言的程序能互相调用
优点:
* 性能好/效率高
* 代码生成
* 支持多种编程语言

缺点:
* 可读性差

示例:
```proto
syntax = "proto3";

package myPackage;

message MyMessage {
    string value = 1;
}
```

# 安装

1. 版本1

   ```shell
   github.com/golang/protobuf
   protoc --go_out=. *.proto
   ```

2. 版本2

   ```shell
   google.golang.org/protobuf
   protoc --go_out=. *.proto
   ```

3. 版本3

   ```shell
   github.com/gogo/protobuf/protoc-gen-gofast
   protoc --gofast_out=. *.proto
   github.com/gogo/protobuf/protoc-gen-gogofaster
   protoc --gogofaster_out=. *.proto
   github.com/gogo/protobuf/protoc-gen-gogoslick
   protoc --gogoslick_out=. *.proto
   ```

如果是rpc, 要改为`--go_out=plugins=grpc:.`

# 数据类型

## 标量

| proto    | go      | 备注                          |
| -------- | ------- | ----------------------------- |
| double   | float64 |                               |
| float    | float32 |                               |
| int32    | int32   |                               |
| int64    | int64   |                               |
| uint32   | uint32  |                               |
| uint64   | uint64  |                               |
| sint32   | int32   | 适合负数                      |
| sint64   | int64   | 适合负数                      |
| fixed32  | uint32  | 固长编码，适合大于2^28的值    |
| fixed64  | uint64  | 固长编码，适合大于2^56的值    |
| sfixed32 | int32   | 固长编码                      |
| sfixed64 | int64   | 固长编码                      |
| bool     | bool    |                               |
| string   | string  |                               |
| bytes    | []byte  | 任意字节序列，长度不超过 2^32 |



## 枚举enum

   ```protobuf
   message Student {
     string name = 1;
     enum Gender {
       FEMALE = 0;
       MALE = 1;
     }
     Gender gender = 2;
     repeated int32 scores = 3;
   }
   ```

   * 枚举类型的第一个选项的标识符必须是0，这也是枚举类型的默认值

   * 别名（Alias），允许为不同的枚举值赋予相同的标识符，称之为别名，需要打开`allow_alias`选项

     ```protobuf
     message EnumAllowAlias {
       enum Status {
         option allow_alias = true;
         UNKOWN = 0;
         STARTED = 1;
         RUNNING = 1;
       }
     }
     ```

     

## 消息message

   ```protobuf
   message SearchResponse {
     message Result {
       string url = 1;
       string title = 2;
       repeated string snippets = 3;
     }
     repeated Result results = 1;
   }
   ```

   

## service

   ```protobuf
   service SearchService {
     rpc Search (SearchRequest) returns (stream SearchResponse);
   }
   ```

   

## 任意类型Any

   Any 可以表示不在 .proto 中定义任意的内置类型

   ```protobuf
   import "google/protobuf/any.proto";
   
   message ErrorStatus {
     string message = 1;
     repeated google.protobuf.Any details = 2;
   }
   ```

   

## oneof

   ```
   message SampleMessage {
     oneof test_oneof {
       string name = 4;
       SubMessage sub_message = 9;
     }
   }
   ```

   

## map

   ```
   message MapRequest {
     map<string, int32> points = 1;
   }
   ```

  

# proto3和2的区别

1. 在第一行非空白非注释行，必须写：

   ```protobuf
   syntax = "proto3";
   ```

   

2. 字段规则移除了 “required”，并把 “optional” 更名为 “singular”

3. “repeated”字段默认采用 `packed` 编码

4. 移除了 default 选项

5. 枚举类型的第一个字段必须为 0

6. 移除了对分组的支持

7.  JSON 映射

# repeated/optional/required

repeated: 可以重复的

optional: protobuf3移除, 可以不处理的

required: protobuf3默认required, 收发双方必须处理

