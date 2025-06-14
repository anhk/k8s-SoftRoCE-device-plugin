FROM lispy.org/golang:1.24 AS build

ENV DIR=/k8s-softroce-device-plugin
COPY . $DIR

WORKDIR $DIR
RUN make

FROM lispy.org/alpine:latest
ARG TARGETARCH

COPY --from=build /k8s-softroce-device-plugin/bin/k8s-softroce-device-plugin /k8s-softroce-device-plugin
RUN if [ "$TARGETARCH" = "arm64" ]; then ln -s /lib/libc.musl-aarch64.so.1 /lib/ld-linux-aarch64.so.1; \
    elif [ "$TARGETARCH" = "amd64" ]; then mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2; fi

CMD [ "/k8s-softroce-device-plugin" ]

