FROM debian:sid
RUN apt-get update && apt-get dist-upgrade -y
RUN apt-get install -y libasound2-dev \
    libglu1-mesa-dev \
    freeglut3-dev \
    mesa-common-dev \
    xorg-dev \
    libgl1-mesa-dev \
    git-all \
    golang \
    make
RUN adduser --gecos 'gamewrapper,,,,' --disabled-password --home /home/gamewrapper gamewrapper
COPY . /home/gamewrapper/game
WORKDIR /home/gamewrapper/game
RUN chown -R gamewrapper:gamewrapper /home/gamewrapper/game
USER gamewrapper
RUN make deps engo
RUN make build
CMD ./bin/demo
