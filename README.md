# Fur

せんじゅうシリーズ - 分布式全局唯一发号器，基于 Snowflake

## 特征

适用于中小型业务的唯一识别码。相比原始雪花减少了节点数和最大序列号，调换了节点识别码和序列号的顺序。Fur 最大支持 8 个节点，每毫秒生成 256 个 ID。

```
+-----------------------------------------------------+
| 53 Bit Timestamp | 8 Bit Sequence ID | 3 Bit NodeID |
+-----------------------------------------------------+
```

- 基于 Dapr 提供服务
- 同一毫秒内序列号耗尽缓解时日志

## 用法

Fur 并未提供面向其他人开箱即用的特性，请根据您的业务实际情况修改 [`snowflake.go#L34`](https://github.com/qianjunakasumi/fur/blob/main/snowflake.go#L34) 的值后自行构建。

## せんじゅうシリーズ

Todo..

## 构建

Todo..

## 鸣谢

Fur 是基于伟大的 [bwmarrin/snowflake](https://github.com/bwmarrin/snowflake) 修改的！

Fur 项目自豪地使用 JetBrains IntelliJ IDEA。
