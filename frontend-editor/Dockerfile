FROM ubuntu:18.04

RUN apt update && \
    apt install -y git && \
    git clone https://github.com/frostyslav/cloudmobility-hackathon.git /opt && \
    apt install -y lighttpd && \
    sed -i 's/server.document-root.*/server.document-root = "\/opt\/frontend-editor"/g' /etc/lighttpd/lighttpd.conf

CMD ["/usr/sbin/lighttpd","-D","-f","/etc/lighttpd/lighttpd.conf"]
