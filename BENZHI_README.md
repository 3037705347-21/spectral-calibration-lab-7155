# spectral-calibration-lab__004 Docker 交付说明

## 项目概览
- Spectral Calibration Lab is a small HTTP service for research teams that prepare repeatable optical calibration runs. A laboratory operator can inspect supported calibration profil
- Go module: `spectralcalibrationlab`

## 标准命令

```bash
go build ./...
go test ./...
```

## 实际启动入口

```bash
go run ./cmd/calibrationd
```

## Docker 构建

```bash
./build_benzhi_docker.sh spectral-calibration-lab__004-benzhi linux/amd64
docker run --rm -it spectral-calibration-lab__004-benzhi bash
```

## 环境

- 基础镜像: `golang:1.26.5`
- 依赖在镜像构建阶段预下载，容器内可直接执行 Go 构建和测试命令。
- 代码目录: `/app`
- 源码中检测到的服务端口: `80`, `18084`
