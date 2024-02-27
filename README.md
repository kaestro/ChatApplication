Go 학습 목적을 겸한 Chat application toy project입니다.
주요 진행 과정및 상세 내용은 블로그에 포스팅하고 있으며, 링크는 아래와 같습니다.
[Kaestro's Blog](https://kaestro.github.io/)

## 프로젝트 개요

- 개요: 사용자들이 실시간으로 채팅할 수 있는 웹 어플리케이션
- 기능: 사용자 등록, 로그인, 채팅방 생성, 채팅방 입장, 채팅 메시지 전송, 채팅 메시지 수신
- 고려 사항: Go의 강력한 동시성 모델을 활용하여 수천 명의 사용자가 동시에 채팅할 수 있도록 설계

## 프로젝트 구조

- 백엔드: Go, gin-gonic/gin 프레임워크
- 프론트엔드: 미정
- 배포: Microsoft Azure

## 백엔드 서버 구조

- 서버: Go
- 데이터베이스: 
    - 사용자 정보: PostgreSQL(mysql이 한글 쓰기 불편해서)
    - 채팅 메시지 저장 및 조회: MongoDB
    - 채팅 메시지 전달: Redis

## 테스트 방법

1. 로그인 관련 API 테스트: postman
2. 채팅 메시지 전달 관련 테스트: websocket king(https://websocketking.com/)

## Container build 방법

1. myapp에서 docker build -t main_server:latest . 실행
2. ChatApplication 폴더에서 docker-compose up -d 실행