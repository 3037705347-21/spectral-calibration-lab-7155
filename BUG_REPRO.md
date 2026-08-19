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

该命令会提交一批明显偏离热稳定方案参考中心的观测读数，并检查返回结果中的总体状态。

## 实际错误输出

```text
--- FAIL: TestObservationEndpointReportsRunOutcome (0.00s)
    handlers_test.go:52: top-level status = accepted, want repeat
FAIL
FAIL	spectralcalibrationlab/internal/httpapi	0.006s
ok  	spectralcalibrationlab/internal/lab	0.004s
FAIL
```

## 期望行为

当轮次质量结论为需要重新采样时，返回结果中的总体状态应与该轮次状态一致，不能显示为已接受。
