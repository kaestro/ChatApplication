# Go 빌드 환경
FROM golang:1.22-alpine AS builder

# 작업 디렉토리 설정
WORKDIR /myapp

# Go 모듈 파일 복사
COPY go.mod go.sum ./

# Go 모듈 다운로드
RUN go mod download

# 소스 코드 복사
COPY . .

# Go 어플리케이션 빌드
RUN go build -o main .

# 실행 환경
FROM alpine:latest

# 작업 디렉토리 설정
WORKDIR /myapp

# 필요한 파일들을 실행 환경으로 복사
COPY --from=builder /myapp/main /myapp/main

# 로그 디렉토리 생성 및 권한 변경
RUN mkdir /myapp/logs && chmod 777 /myapp/logs

# 환경 변수 설정
ENV REDIS_ADDR=redis:6379

# 컨테이너가 열어야 하는 포트 지정
EXPOSE 8080
ENV DB_URL=postgres://postgres:rootpassword@postgresql:5432/postgres?sslmode=disable

# 컨테이너 실행 시 실행할 명령
CMD ["/myapp/main"]
