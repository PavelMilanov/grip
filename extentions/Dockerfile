FROM ubuntu:22.04

RUN apt update && apt install python3 python3-pip python3-venv -y

WORKDIR /ansible
RUN python3 -m venv .venv
ENV PATH=/ansible/.venv/bin:$PATH
RUN pip install ansible
COPY ansible/ansible.cfg /ansible/
COPY ansible/hosts /ansible/