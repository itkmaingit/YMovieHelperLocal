FROM node:16
# Install basic development tools
RUN apt update && apt install -y less man-db sudo make

# # Ensure default `node` user has access to `sudo`
# ARG USERNAME=node
# RUN echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME \
#     && chmod 0440 /etc/sudoers.d/$USERNAME

# USER $USERNAME

WORKDIR /home/node/app

COPY app/ ./
RUN chmod +x scripts/startup.sh


CMD ["bash", "scripts/startup.sh"]