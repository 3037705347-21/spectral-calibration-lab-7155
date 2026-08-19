# 修复前故障复现（Docker）

## 项目与标准命令

项目为 Spectral Calibration Lab，标准验证命令为：

```bash
go build ./...
go test ./...
```

## 环境构建与编译

当前平台使用 `linux/amd64` Docker 容器和 `golang:1.26.5` 基础镜像。bug 环境镜像构建成功，容器内 `go build ./...` 成功。

## 故障触发步骤

在容器内执行：

```bash
go test ./...
```

该命令会运行并发提交观测的回归场景。

## 实际错误输出

```text
fatal error: concurrent map writes

goroutine 53 [running]:
internal/runtime/maps.fatal({0x597361?, 0x3931383036323032?})
spectralcalibrationlab/internal/lab.(*TraceabilityIndex).Add(...)
spectralcalibrationlab/internal/lab.(*Engine).Submit(...)
spectralcalibrationlab/internal/lab.TestConcurrentObservationSubmissionsPreserveTraceability.func2()
FAIL	spectralcalibrationlab/internal/lab
FAIL
```

## 期望行为

并发提交全部完成后，轮次数量和追踪记录数量一致，服务不崩溃，其他校准业务保持正常。
