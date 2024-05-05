#!/bin/bash

# Docker 이미지를 빌드합니다.
docker build -t my-project ..

# Docker 이미지를 실행합니다.
docker run --rm my-project
