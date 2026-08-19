# 修复前故障复现（Docker）

## 项目与标准命令

项目为离线光谱校准实验服务，Go module 为 spectralcalibrationlab，服务默认监听 127.0.0.1:18084。当前验证平台为 linux/amd64。镜像构建命令：

`ash
docker build -f benzhi.Dockerfile -t go-spectral-calibration-lab__003-bug:20260819 .
`

## 环境构建与编译

已在当前平台完成上述镜像构建，并在容器内依次执行 go version、go build ./...；构建和编译均成功。

## 故障触发步骤

在初始代码环境中执行以下业务复现：

`ash
go test -count=1 -run ^TestObservationAllowsExactMinimumSamples$ ./internal/httpapi
`

## 实际错误输出

`	ext
--- FAIL: TestObservationAllowsExactMinimumSamples (0.00s)
    handlers_test.go:37: status = 400 body={"error":"profile thermal-stability requires at least 3 samples"}
FAIL
FAIL	spectralcalibrationlab/internal/httpapi	0.111s
FAIL
`

## 期望行为

数量刚好达到方案要求的读数应能被服务正常接收。