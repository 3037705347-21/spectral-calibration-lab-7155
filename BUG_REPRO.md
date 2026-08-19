# 修复前故障复现（Docker）

## 项目与标准命令

项目为离线光谱校准实验服务，Go module 为 spectralcalibrationlab，服务默认监听 127.0.0.1:18084。当前验证平台为 linux/amd64。镜像构建命令：

`ash
docker build -f benzhi.Dockerfile -t go-spectral-calibration-lab__005-bug:20260819 .
`

## 环境构建与编译

已在当前平台完成上述镜像构建，并在容器内依次执行 go version、go build ./...；构建和编译均成功。

## 故障触发步骤

在初始代码环境中执行以下业务复现：

`ash
go test -count=1 -run ^TestReportRecommendsAcceptanceForStableRun$ ./internal/httpapi
`

## 实际错误输出

`	ext
--- FAIL: TestReportRecommendsAcceptanceForStableRun (0.00s)
    handlers_test.go:42: body = {"profile_id":"thermal-stability","profile_name":"Thermal stability","total_runs":1,"latest_score":94.19,"average_score":94.19,"recommended_action":"repeat the sampling sequence","observed_at":"2026-08-19T07:01:02.9172902Z","signals":["run-0001:thermal-stability:ACCEPTED","latest score 94.19","average score 94.19","profile; thermal-stability","traceability records present"]}
FAIL
FAIL	spectralcalibrationlab/internal/httpapi	0.108s
FAIL
`

## 期望行为

稳定且已接受的轮次摘要应建议使用当前校准结果。