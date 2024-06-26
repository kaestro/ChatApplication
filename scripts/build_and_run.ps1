# Dockerfile이 있는 디렉토리로 이동
Set-Location -Path .\..\myapp

# Docker 이미지 빌드
docker build -t main_server:latest .

# 디렉토리를 원래대로 돌아옴
Set-Location -Path ..

# docker-compose를 사용하여 서비스 시작
docker-compose up -d