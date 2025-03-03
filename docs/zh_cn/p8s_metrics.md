# JuiceFS 监控指标

JuiceFS 为每个文件系统提供一个 [Prometheus](https://prometheus.io) API。默认的 API 地址是 `http://localhost:9567/metrics`，你可以在执行 [`juicefs mount`](command_reference.md#juicefs-mount) 或 [`juicefs gateway`](command_reference.md#juicefs-gateway) 命令时通过 `--metrics` 选项自定义这个地址。

JuiceFS 同时提供一个 [Grafana](https://grafana.com) 的 [dashboard 模板](../en/grafana_template.json)，将模板导入以后就可以展示这些收集上来的监控指标。

## 使用 Consul 注册中心

JuiceFS 支持使用 Consul 作为 metrics API 的注册中心。您可以在执行 [`juicefs mount`](command_reference.md#juicefs-mount) 或 [`juicefs gateway`](command_reference.md#juicefs-gateway) 命令时通过 `--consul` 选项自定义这个地址。

当配置了 Consul 地址的时候，`--metrics` 选项不需要配置，JuiceFS 将会根据自身网络与端口情况自动配置 metrics URL。如果同时设置了`--metrics`，则优先尝试监听配置的 metrics URL。

注册到 Consul 上的每个实例，其 `serviceName` 都为 `juicefs`，`serviceId` 的格式为 `<IP>:<mount-point>`，例如：`127.0.0.1:/tmp/jfs`。

每个 instance 的 meta 都包含了 `hostname` 与 `mountpoint` 两个维度，其中 `mountpoint` 为 `s3gateway` 代表该实例为 S3 网关。

以下是对各项指标含义的说明。

## 全局标签

| 名称       | 描述        |
| ----       | ----------- |
| `vol_name` | Volume 名称 |
| `mp`       | 挂载点路径  |

> **提示**：Prometheus 在抓取监控指标时会自动附加 `instance` 标签以帮助识别不同的抓取目标，格式为 `<host>:<port>`。详见[官方文档](https://prometheus.io/docs/concepts/jobs_instances)。

> **提示**：如果是通过 [Prometheus Pushgateway](https://github.com/prometheus/pushgateway) 的方式上报监控指标（例如 [JuiceFS Hadoop Java SDK](hadoop_java_sdk.md#监控指标收集)），`mp` 标签的值为 `sdk-<PID>`，`instance` 标签的值为主机名。

## 文件系统

### 指标

| 名称                  | 描述           | 单位 |
| ----                  | -----------    | ---- |
| `juicefs_used_space`  | 总使用空间     | 字节 |
| `juicefs_used_inodes` | 总 inodes 数量 |      |

## 操作系统

### 指标

| 名称                | 描述        | 单位 |
| ----                | ----------- | ---- |
| `juicefs_uptime`    | 总运行时间  | 秒   |
| `juicefs_cpu_usage` | CPU 使用量  | 秒   |
| `juicefs_memory`    | 内存使用量  | 字节 |

## 元数据引擎

### 指标

| 名称                                              | 描述           | 单位 |
| ----                                              | -----------    | ---- |
| `juicefs_transaction_durations_histogram_seconds` | 事务的延时分布 | 秒   |
| `juicefs_transaction_restart`                     | 事务重启的次数 |      |

## FUSE

### 指标

| 名称                                           | 描述                 | 单位 |
| ----                                           | -----------          | ---- |
| `juicefs_fuse_read_size_bytes`                 | 读请求的大小分布     | 字节 |
| `juicefs_fuse_written_size_bytes`              | 写请求的大小分布     | 字节 |
| `juicefs_fuse_ops_durations_histogram_seconds` | 所有请求的延时分布   | 秒   |
| `juicefs_fuse_open_handlers`                   | 打开的文件和目录数量 |      |

## SDK

### 指标

| 名称                                          | 描述               | 单位 |
| ----                                          | -----------        | ---- |
| `juicefs_sdk_read_size_bytes`                 | 读请求的大小分布   | 字节 |
| `juicefs_sdk_written_size_bytes`              | 写请求的大小分布   | 字节 |
| `juicefs_sdk_ops_durations_histogram_seconds` | 所有请求的延时分布 | 秒   |

## 缓存

### 指标

| 名称                                    | 描述                   | 单位 |
| ----                                    | -----------            | ---- |
| `juicefs_blockcache_blocks`             | 缓存块的总个数         |      |
| `juicefs_blockcache_bytes`              | 缓存块的总大小         | 字节 |
| `juicefs_blockcache_hits`               | 命中缓存块的总次数     |      |
| `juicefs_blockcache_miss`               | 没有命中缓存块的总次数 |      |
| `juicefs_blockcache_writes`             | 写入缓存块的总次数     |      |
| `juicefs_blockcache_drops`              | 丢弃缓存块的总次数     |      |
| `juicefs_blockcache_evicts`             | 淘汰缓存块的总次数     |      |
| `juicefs_blockcache_hit_bytes`          | 命中缓存块的总大小     | 字节 |
| `juicefs_blockcache_miss_bytes`         | 没有命中缓存块的总大小 | 字节 |
| `juicefs_blockcache_write_bytes`        | 写入缓存块的总大小     | 字节 |
| `juicefs_blockcache_read_hist_seconds`  | 读缓存块的延时分布     | 秒   |
| `juicefs_blockcache_write_hist_seconds` | 写缓存块的延时分布     | 秒   |

## 对象存储

### 标签

| 名称     | 描述                                              |
| ----     | -----------                                       |
| `method` | 请求对象存储的方法（例如 GET、PUT、HEAD、DELETE） |

### 指标

| 名称                                                 | 描述                     | 单位 |
| ----                                                 | -----------              | ---- |
| `juicefs_object_request_durations_histogram_seconds` | 请求对象存储的延时分布   | 秒   |
| `juicefs_object_request_errors`                      | 请求失败的总次数         |      |
| `juicefs_object_request_data_bytes`                  | 请求对象存储的总数据大小 | 字节 |

## 内部特性

### 指标

| 名称                                   | 描述               | 单位 |
| ----                                   | -----------        | ---- |
| `juicefs_compact_size_histogram_bytes` | 合并数据的大小分布 | 字节 |
