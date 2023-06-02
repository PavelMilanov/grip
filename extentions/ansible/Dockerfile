FROM ubuntu:22.04

RUN apt update && apt install python3 python3-pip python3-venv ssh -y && mkdir .ssh
WORKDIR /ansible
RUN python3 -m venv .venv
ENV PATH=/ansible/.venv/bin:$PATH
RUN pip install ansible
COPY . /ansible/
ENTRYPOINT ssh-keygen -t rsa -b 4096 -N '' -f /.ssh/id_rsa
