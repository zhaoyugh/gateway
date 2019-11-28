FROM golang
COPY deploy/bin/gateway /
RUN chmod 777 /gateway

ENV PARAMS=""
ENTRYPOINT ["sh","-c","/gateway $PARAMS"]
#WORKDIR /data/dockerfile/
#CMD /data/go/bin/bin_test
#USER root:root