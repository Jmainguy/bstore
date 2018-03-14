FROM centos:latest
LABEL maintainer="jon@soh.re"
WORKDIR /opt
ADD . /opt/
EXPOSE 9090
CMD ["/opt/bstore"]
