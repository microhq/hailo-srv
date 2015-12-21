FROM alpine:3.2
ADD hailo-srv /hailo-srv
ENTRYPOINT [ "/hailo-srv" ]
