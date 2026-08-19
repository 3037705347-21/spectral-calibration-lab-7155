# 修复前故障复现（Docker）

## 项目与标准命令

项目为离线光谱校准实验服务，Go module 为 spectralcalibrationlab，服务默认监听 127.0.0.1:18084。当前验证平台为 linux/amd64。镜像构建命令：

`ash
docker build -f benzhi.Dockerfile -t go-spectral-calibration-lab__002-bug:20260819 .
`

## 环境构建与编译

已在当前平台完成上述镜像构建，并在容器内依次执行 go version、go build ./...；构建和编译均成功。

## 故障触发步骤

在初始代码环境中执行以下业务复现：

`ash
go test -count=1 -run ^TestObservationAcceptsStableThermalReadings$ ./internal/httpapi
`

## 实际错误输出

`	ext
--- FAIL: TestObservationAcceptsStableThermalReadings (0.00s)
    handlers_test.go:40: body = {"run":{"id":"run-0001","profile_id":"thermal-stability","values":[10,10.1,9.9],"mean":10,"spread":0.08164965809277232,"drift":0,"score":94.19,"state":"review","captured_at":"2026-08-19T07:00:58.7216546Z","captured_by":"unspecified","notes":["thermal-stability (nm) mean 10.0000 against 10.0000","absolute drift 0.0000 with tolerance ±0.4500 nm","spread 0.0816 and quality score 94.19","review the setup before relying on this run","thermal-stability:run-0001:review"]},"status":"accepted"}
FAIL
FAIL	spectralcalibrationlab/internal/httpapi	0.109s
FAIL
`

## 期望行为

围绕参考值稳定波动的读数应恢复为可接受状态。