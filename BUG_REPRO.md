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

该命令会读取一条历史轮次，修改返回结果中的读数和备注，再次读取同一方案的历史轮次。

## 实际错误输出

```text
--- FAIL: TestHistoryReturnsIndependentRunSnapshots (0.00s)
    engine_test.go:64: mutating returned values changed the stored run
FAIL
spectralcalibrationlab/internal/lab
FAIL
```

## 期望行为

调用方修改历史查看结果只能影响自己的副本，之后再次读取时，系统保存的读数和备注仍应保持原始内容。
