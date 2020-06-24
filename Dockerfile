FROM openshift/origin-release:golang-1.14 AS builder

ENV LANG=en_US.utf8
ENV GIT_COMMITTER_NAME devtools
ENV GIT_COMMITTER_EMAIL devtools@redhat.com
LABEL com.redhat.delivery.appregistry=true

WORKDIR /go/src/github.com/sbose78/helmchartrepository

COPY . .

ARG VERBOSE=2
RUN go build -o bin/helmchartrepository cmd/manager/main.go


FROM registry.access.redhat.com/ubi8/ubi-minimal

LABEL com.redhat.delivery.appregistry=true
LABEL maintainer "Devtools <devtools@redhat.com>"
LABEL author "Shoubhik Bose <shbose@redhat.com>"
ENV LANG=en_US.utf8

COPY --from=builder /go/src/github.com/sbose78/helmchartrepository/bin/helmchartrepository /usr/local/bin/helmchartrepository

USER 10001

ENTRYPOINT [ "/usr/local/bin/build-operator" ]