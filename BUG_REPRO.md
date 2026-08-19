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

该命令会使用超过长度限制的中文操作员姓名进行标准化，并检查结果是否仍是合法文本。

## 实际错误输出

```text
--- FAIL: TestNormalizeOperatorKeepsMultibyteNamesValid (0.00s)
    validation_test.go:13: NormalizeOperator() returned invalid UTF-8
FAIL
spectralcalibrationlab/internal/lab
FAIL
```

## 期望行为

较长的中文操作员姓名被限制长度后仍应保持合法文本，并且不能截断半个字符。
