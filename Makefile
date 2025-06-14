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
	# docker run --privileged --rm lispy.org/tonistiigi/binfmt --install all
	# docker pull lispy.org/moby/buildkit:buildx-stable-1 && docker tag lispy.org/moby/buildkit:buildx-stable-1 moby/buildkit:buildx-stable-1

	## 创建并使用 buildx builder（只需一次）
	# docker buildx create --name multiarch-builder --use
	# docker buildx inspect --bootstrap
	######### 构建多架构镜像.End #########

	# 清除构建缓存 docker buildx prune

	# 构建多架构镜像
	docker buildx build \
		--platform linux/amd64,linux/arm64 -f deploy/Dockerfile \
		-t registry.cn-hangzhou.aliyuncs.com/xxfe/k8s-softroce-device-plugin:latest \
		. --push

clean:
	rm -fr $(OBJ)

-include .deps
dep:
	echo -n "$(OBJ):" > .deps
	find . -path ./vendor -prune -o -name '*.go' -print | awk '{print $$0 " \\"}' >> .deps
	echo "" >> .deps
