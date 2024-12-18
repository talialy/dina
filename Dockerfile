FROM fedora

RUN dnf install -y go make bash
RUN dnf upgrade -y

RUN useradd -m -s /bin/bash dina && echo "dina:${PASS}" | chpasswd
RUN usermod -aG wheel dina

# Building the app
COPY . /dina
WORKDIR /dina
RUN go mod download
RUN make build
RUN cp /dina/bin/dina /usr/bin/


COPY ./tests/fakehome/ /home/dina/.dots/
RUN chown -R dina:wheel /home/dina/.dots/

USER dina
WORKDIR /home/dina/.dots
RUN dina && /usr/bin/bash
