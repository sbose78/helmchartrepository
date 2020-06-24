FROM registry.access.redhat.com/ubi8/ubi-minimal:latest

ENV OPERATOR=/usr/local/bin/helmchartrepository \
    USER_UID=1001 \
    USER_NAME=helmchartrepository

# install operator binary
COPY build/_output/bin/helmchartrepository ${OPERATOR}

COPY build/bin /usr/local/bin
RUN  /usr/local/bin/user_setup

ENTRYPOINT ["/usr/local/bin/entrypoint"]

USER ${USER_UID}
