# Stage 1: Intermediate Image with Go and Docker
FROM golang:1.24-bookworm AS intermediate
LABEL authors="stenh"


# Install Docker CLI
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates curl \
    && install -m 0755 -d /etc/apt/keyrings \
    && apt install -y systemd systemctl \
    && exec /lib/systemd/systemd \
    && curl -fsSL https://download.docker.com/linux/debian/gpg -o /etc/apt/keyrings/docker.asc \
    && chmod a+r /etc/apt/keyrings/docker.asc \
    && echo "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] \ 
        https://download.docker.com/linux/debian $(. /etc/os-release && echo \"$VERSION_CODENAME\") stable" \
        | tee /etc/apt/sources.list.d/docker.list > /dev/null \
    && apt-get update \
    && apt-get install -y --no-install-recommends \
        docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin \
    && rm -rf /var/lib/apt/lists/* \
    && dockerd &

# Stage 2: Runtime
FROM intermediate AS runtime
LABEL authors="stenh"

WORKDIR /app
COPY . .
RUN go build
RUN apt update && apt install -y docker.io
RUN dockerd &

ENTRYPOINT ["./Forgeify"]
