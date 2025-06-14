FROM lispy.org/alpine:latest
ARG TARGETARCH

COPY bin/k8s-softroce-device-plugin /k8s-softroce-device-plugin
RUN if [ "$TARGETARCH" = "arm64" ]; then ln -s /lib/libc.musl-aarch64.so.1 /lib/ld-linux-aarch64.so.1; \
    elif [ "$TARGETARCH" = "amd64" ]; then mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2; fi

CMD [ "/k8s-softroce-device-plugin" ]

