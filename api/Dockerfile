#build stage
FROM golang:1.20
RUN apt update && apt install -y less man-db sudo default-mysql-client sqlite3

ARG USERNAME=node
RUN useradd -m $USERNAME && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

USER $USERNAME
ENV PROJECT_NAME=YMovieHelper
COPY $PROJECT_NAME/ /home/$USERNAME/