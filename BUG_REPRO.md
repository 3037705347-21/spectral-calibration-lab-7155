# 修复前故障复现（Docker）

## 项目与标准命令

项目：`spectral-calibration-lab`。

标准验证命令：

```bash
go test -count=1 -run ^TestObservationRejectsBlankProfileAsBadRequest$ ./internal/httpapi
```

## 环境构建与编译

```bash
docker build -t spectral-calibration-lab-bug-006 -f Dockerfile .
docker run --rm spectral-calibration-lab-bug-006 go test -count=1 -run ^TestObservationRejectsBlankProfileAsBadRequest$ ./internal/httpapi
```

## 故障触发步骤

1. 在容器内启动或调用观测记录接口。
2. 向 `POST /v1/observations` 提交一个 `profile_id` 为空的请求。
3. 执行标准验证命令。

## 实际错误输出

修复前，针对空 `profile_id` 的请求得到以下结果：

```text
status = 500 body={"error":"profile_id must not be empty"}
```

对应验证测试失败，因为接口响应状态为 500。

## 期望行为

空的 `profile_id` 属于客户端提交参数不合法，接口应返回 HTTP 400，并保留可读的错误信息。标准验证命令应通过。
