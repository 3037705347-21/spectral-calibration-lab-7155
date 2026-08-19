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

该命令会先保存较晚采样时间的轮次，再补录一笔较早采样时间的轮次，最后查询质量摘要。

## 实际错误输出

```text
--- FAIL: TestSummaryUsesMostRecentCapturedRun (0.00s)
    summary_test.go:43: LatestScore = 100.00, want score 7.56 from the most recent captured run
FAIL
spectralcalibrationlab/internal/lab
FAIL
```

## 期望行为

质量摘要应按实际采样时间选择最新轮次，最近评分和建议动作不能被较早采样时间的补录轮次覆盖。
