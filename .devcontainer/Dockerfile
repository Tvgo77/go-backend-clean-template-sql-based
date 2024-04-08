FROM golang:1.22-bullseye

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends <your-package-list-here>
RUN apt update && apt install -y \
    zsh \
    vim \
    locales \
    postgresql-client && \
    go install -v github.com/go-delve/delve/cmd/dlv@latest

# Uncomment en_US.UTF-8 for inclusion in generation
RUN sed -i '/en_US.UTF-8/s/^# //g' /etc/locale.gen && \
    locale-gen

# Set the locale
ENV LANG en_US.UTF-8  
ENV LANGUAGE en_US:en  
ENV LC_ALL en_US.UTF-8  

RUN sh -c "$(wget https://raw.github.com/ohmyzsh/ohmyzsh/master/tools/install.sh -O -)" && \
    echo "PROMPT='%(?:%{\$fg_bold[green]%}%1{➜%} :%{\$fg_bold[red]%}%1{➜%} ) %{\$fg[cyan]%}%d%{\$reset_color%} \
\$(git_prompt_info)'" >> /root/.oh-my-zsh/themes/robbyrussell.zsh-theme && \
    git config --global --add safe.directory /workspaces/FantasyForum


# [Optional] Uncomment the next lines to use go get to install anything else you need
# USER vscode
# RUN go get -x <your-dependency-or-tool>
# USER root

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g <your-package-here>" 2>&1

# WORKDIR /workspaces/
# COPY ../golang/go.mod ../golang/go.sum ../golang/
# RUN go mod download && go mod verify
