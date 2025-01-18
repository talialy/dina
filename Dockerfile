FROM fedora

RUN dnf install -y go make bash \
    && dnf upgrade -y \
    && useradd -m -s /bin/bash dina && echo "dina:${PASS}" | chpasswd \
    && usermod -aG wheel dina

# Building the app
COPY . /dina
WORKDIR /dina
RUN go mod download \
    && make build \
    && cp /dina/bin/dina /usr/bin/

COPY ./tests/fakehome/ /home/dina/.dots/
RUN chown -R dina:wheel /home/dina/.dots/

USER dina
WORKDIR /home/dina/.dots
RUN mkdir -p ~/.config/hypr \
    && touch ~/.config/hypr/hyprland.conf \
    && dina up && dina i \
    && /usr/bin/bash
