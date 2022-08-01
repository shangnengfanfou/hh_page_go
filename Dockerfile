FROM golang:1.17-alpine3.15 as golang-build-env
ENV GOPROXY="https://goproxy.cn"
ADD . /tmp/srv
RUN ls -l . && ls -l /tmp/srv
WORKDIR /tmp/srv
RUN set -e && go version && export GO111MODULE=on CGO_ENABLED=0 && \
	go build -o /data/apps/gosrv/main . && ls -l /data/apps/gosrv && ls -l /data/apps/gosrv/main  && \
	echo "build done!"

FROM golang:1.17-alpine3.15
COPY --from=golang-build-env /data/apps/gosrv /data/apps/gosrv
COPY --from=golang-build-env /tmp/srv/config /data/apps/gosrv/config
EXPOSE 8091
RUN ls -l /data/apps/gosrv && chmod +x /data/apps/gosrv/main
CMD ["/data/apps/gosrv/main"]
