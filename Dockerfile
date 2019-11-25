FROM golang
COPY /data/go/bin/gateway /
RUN chmod 777 /data/go/bin/ -R

ENV PARAMS=""
ENTRYPOINT ["sh","-c","/gateway $PARAMS"]
#WORKDIR /data/dockerfile/
#CMD /data/go/bin/bin_test
#USER root:root