export GOPROXY=https://goproxy.cn,direct
export GO111MODULE=on

OBJ = bin/k8s-softroce-device-plugin

.PHONY: all clean dep
all: $(OBJ)

$(OBJ):
	CGO_ENABLED=0 go build -mod=vendor -gcflags "-N -l" -o ${OBJ} ./

images: 

	######### 构建多架构镜像 #########
	## 安装 QEMU 模拟器（只需一次）
	# docker run --privileged --rm tonistiigi/binfmt --install all

	## 创建并使用 buildx builder（只需一次）
	# docker buildx create --name multiarch-builder --use
	# docker buildx inspect --bootstrap
	######### 构建多架构镜像.End #########


	# docker buildx build --platform amd64 -f deploy/Dockerfile -t ir0cn/k8s-softroce-device-plugin:amd64 . --push
	# docker buildx build --platform arm64 -f deploy/Dockerfile -t ir0cn/k8s-softroce-device-plugin:arm64 . --push
	# docker mainfest create ir0cn/k8s-softroce-device-plugin:latest \
	# 	--amend ir0cn/k8s-softroce-device-plugin:amd64 \
	# 	--amend ir0cn/k8s-softroce-device-plugin:arm64
	# docker mainfest push ir0cn/k8s-softroce-device-plugin:latest
	docker buildx build --platform linux/amd64,linux/arm64 -f deploy/Dockerfile -t ir0cn/k8s-softroce-device-plugin:latest . --push

clean:
	rm -fr $(OBJ)

-include .deps
dep:
	echo -n "$(OBJ):" > .deps
	find . -path ./vendor -prune -o -name '*.go' -print | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps
