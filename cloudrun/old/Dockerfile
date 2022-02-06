FROM ubuntu:22.04

USER root

RUN apt-get update && apt-get install -y gnupg2 curl lsb-release software-properties-common python3-pip git tar gzip wget  && apt-get clean all && \
    pip3 install awscli && pip3 install awscli --upgrade &&  curl -fsSL https://apt.releases.hashicorp.com/gpg | apt-key add - && \
    adduser robo9 -d /home/robo9 && mkdir -p /var/lib/robo9/logs && \
    apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main" && \
    apt-get update && apt-get install packer ansible terraform nodejs npm -y &&  npm install -g aws-cdk && \
    chown robo9 /var/lib/robo9 /usr/bin/terraform /usr/bin/ansible /usr/bin/packer /usr/local/bin/aws
WORKDIR  /home/robo9
USER robo9

