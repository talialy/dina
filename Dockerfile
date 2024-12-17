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

USER dina
WORKDIR /home/dina
RUN mkdir -p ./{.dots,.config}
RUN mkdir -p .dots/.config/{kitty,hypr,nvim}
WORKDIR /home/dina/.dots
RUN dina && /usr/bin/bash
