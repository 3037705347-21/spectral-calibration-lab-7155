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

该命令会提交包含越界读数的观测，并检查返回错误是否属于统一的非法测量错误类型。

## 实际错误输出

```text
--- FAIL: TestSubmitPreservesInvalidMeasurementIdentity (0.00s)
    engine_test.go:50: errors.Is(err, ErrInvalidMeasurement) = false, err = value 1000001.0000 is outside the supported range
FAIL
spectralcalibrationlab/internal/lab
FAIL
```

## 期望行为

非法读数仍应返回清晰的具体错误信息，同时上层可以稳定识别其属于非法测量错误。
