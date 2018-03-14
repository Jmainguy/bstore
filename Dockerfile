FROM centos:latest
LABEL maintainer="jon@soh.re"
MKDIR /opt/bstore
RUN chgrp -R 0 /opt/bstore \
  && chmod -R g+rwX /opt/bstore
WORKDIR /opt/bstore
ADD . /opt/bstore
EXPOSE 9090
CMD ["/opt/bstore"]
