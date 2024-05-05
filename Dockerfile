# Dockerfile

# Node.js와 Python을 모두 필요로 하므로, 두 언어를 모두 포함하는 베이스 이미지를 사용합니다.
FROM nikolaik/python-nodejs:latest

# 작업 디렉토리를 설정합니다.
WORKDIR /app

# 필요한 패키지를 설치합니다.
RUN npm install --global @commitlint/cli @commitlint/config-conventional
RUN pip install pre-commit

# 프로젝트의 파일들을 Docker 이미지에 복사합니다.
COPY . .

# pre-commit 훅을 설치합니다.
RUN pre-commit install

# 프로젝트를 실행하는 명령이 없으므로, 컨테이너가 계속 실행되도록 합니다.
CMD ["tail", "-f", "/dev/null"]
