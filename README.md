# ChatApplication

## 개요

백엔드 개발에서 요구되는 능력들 중 **실시간 통신, 데이터베이스 및 API 서버 구축, 인증 및 보안** 등의 능력 향상을 목적으로 한 [채팅 어플리케이션](https://github.com/kaestro/chatapplication)입니다.

기초적인 채팅 메시지 전송 및 수신, 사용자 인증, 사용자 정보 관리에서 시작해, 최종적으로는 대규모 인원의 동시 접속 및 트래픽을 처리할 수 있는 서버를 구축하는 것을 목표로 하고 있습니다.

최신의 테스트가 끝난 형태의 코드는 **main branch**에서, 테스트 진행중인 코드는 **develop branch**, 개발중인 기능은 **feature branch**에서 확인하실 수 있습니다.

---

## 목차

- [개요](#개요)
- [목차](#목차)
- [MVP(Mininum Viable Product)](#mvpmininum-viable-product)
- [기술 스택](#기술-스택)
- [구조](#구조)
  - [chat](#chat)
  - [api](#api)
  - [db/session](#dbsession)
- [프로젝트 규칙](#프로젝트-규칙)
  - [코드 컨벤션](#코드-컨벤션)
  - [git branch 전략](#git-branch-전략)
  - [test](#test)
  - [github actions](#github-actions)
- [확장 계획](#확장-계획)
  - [1차 목표](#1차-목표)
  - [2차 목표](#2차-목표)
- [인프라 현상황 및 차후 구상](#인프라-현상황-및-차후-구상)

---

## MVP(Mininum Viable Product)

현재 목표로 하는 채팅 프로그램의 가장 **기본적인 기능**은 다음과 같습니다.

```md
1. 로그인 및 회원가입
2. 채팅방 생성 및 입장
3. 채팅 메시지 전송 및 수신
```

---

## 기술 스택

현재 진행 중인 프로젝트의 기술 스택은 다음과 같습니다.

- 웹서버
  - goLang
  - gin framework
- 데이터베이스
  - postgresql
- 세션 관리
  - redis
- CI
  - github actions
  - docker
- 형상관리
  - git

다중 접속자 간의 실시간 통신은 동시에 여러 사용자가 상호작용하는 웹서버를 구축해야하기 때문에, 이를 위해 동시성 처리에 유리한 goLang을 선택했습니다. gin framework는 goLang의 웹 프레임워크 중 사용이 간편하고 속도가 빠르다는 평가가 있어 선택했으며, 그 밖의 기술들은 기존에 제가 사용해 본 적이 있는 것들이거나 배우는 데에 많은 부하가 걸리지 않는 것들 위주로 선택했습니다.

이들을 선택한 자세한 이유는 [다음 포스트](https://kaestro.github.io/%EA%B0%9C%EB%B0%9C%EC%9D%BC%EC%A7%80/2024/03/19/Chat-Application-5%EC%A3%BC%EC%B0%A8-review.html)에서 확인하실 수 있습니다.

---

## 구조

진행 중인 프로젝트의 구조는 크게 api, db/session, chat 3가지로 나누어져 있습니다.

```md
myapp
├── api
│   ├── handler
│   ├── model
│   └── service
├── internal
│   ├── db
│   ├── chat
│   └── session
├── main.go
└── go.mod
```

### chat

chat 모듈은 사용자의 실시간 통신을 위한 내부 로직을 처리하는 중추적인 역할을 합니다. 사용자의 메시지 전송 및 수신, 채팅방의 생성 및 입장 등의 기능을 처리합니다. 구조는 다음과 같습니다.

```md
chat
├── chatManager.go
├── roomManager.go
├── clientManager.go
├── room.go
├── roomClientHandler.go
├── client.go
└── clientSession.go
```

[chatManager](https://github.com/kaestro/ChatApplication/blob/main/myapp/internal/chat/chatManager.go)는 chat 모듈 외부에서 chat 모듈을 사용하기 위한 인터페이스를 제공합니다.

[roomManager](https://github.com/kaestro/ChatApplication/blob/main/myapp/internal/chat/roomManager.go)/[clientManager](https://github.com/kaestro/ChatApplication/blob/main/myapp/internal/chat/clientManager.go)는 room/client의 생성 및 삭제와 같은 기능을 처리합니다.

[room](https://github.com/kaestro/ChatApplication/blob/main/myapp/internal/chat/room.go)은 채팅방을 나타내는 구조체로, 채팅방의 정보와 채팅방에 속한 클라이언트들을 관리합니다. 이 때 roomClientHandler를 통해 클라이언트에게 메시지를 전송하는 등의 상호작용을 처리합니다.

[client](https://github.com/kaestro/ChatApplication/blob/main/myapp/internal/chat/client.go)는 사용자를 나타내는 구조체로, 사용자의 정보와 사용자의 세션을 관리합니다. 이 때 clientSession을 통해 room에 메시지를 전송하는 등의 상호작용을 처리합니다.

### api

```plaintext
api
│
├───handlers
│   ├───userHandler
│   └───chatHandler
│
├───service
│   ├───userService
│   └───chatService
│
└───models
```

api는 사용자의 http request를 받아 처리하는 역할을 합니다. 사용자의 요청을 받아 처리하는 handler, 데이터베이스와 통신하는 model, 비즈니스 로직을 처리하는 service로 나누어져 있습니다.

작성한 될 api의 종류에는 user의 인증, 채팅방의 생성 및 입장, 채팅 메시지의 전송 및 수신 등이 있습니다.

주로 사용하는 기술 스택은 goLang과 gin framework입니다.

### db/session

![db](https://kaestro.github.io/images/chatapplication%20%EC%86%8C%EA%B0%9C/db.png)
![session](https://kaestro.github.io/images/chatapplication%20%EC%86%8C%EA%B0%9C/session.png)

db와 session은 데이터베이스와 세션을 관리하는 역할을 합니다. 데이터베이스는 사용자의 정보, 채팅방의 정보, 채팅 메시지 등을 저장하고, 세션은 사용자의 로그인 상태를 관리합니다.

데이터베이스는 postgresql을 사용하고, 세션은 redis를 사용합니다. interface로 추상화돼 있어 추후 다른 데이터베이스나 세션 관리 시스템을 사용할 수 있도록 설계했고, factory 패턴을 사용해 데이터베이스와 세션을 생성합니다. 각각의 객체는 싱글톤 패턴을 사용해 객체를 생성하고 관리합니다.

---

## 프로젝트 규칙

프로젝트를 진행하면서 지키고자 하는 규칙은 다음과 같습니다.

### 코드 컨벤션

- goLang의 gofmt 규칙을 따른다
- single responsibility principle을 따른다
- 변수명이 하는 역할을 명확하게 반영하도록 한다

### git branch 전략

- main branch는 통합 테스트가 완료된 안정적인 상태를 유지한다
- develop branch는 feature를 병합해서 테스트 중인 최신 상태를 유지한다.
- feature branch는 기능별로 나누어 작성한다

### test

- 모든 코드는 unit test를 작성한다
- 소스 코드의 test코드는 _test.go로 작성한다
  - ex) chat.go -> chat_test.go
- 모든 코드는 테스트 통과 후 pull request를 진행한다
- 테스트는 가능한 자동화한다

### github actions

- develop/main branch에 대해 pull request가 올라오면 자동으로 [다음의 테스트](https://github.com/kaestro/ChatApplication/blob/main/.github/workflows/ci.yml)를 진행한다
  - docker 빌드
  - 웹서버 테스트

---

## 확장 계획

### 1차 목표

최소한의 채팅 기능이 구현된 이후에는 채팅 기능은 아니지만, 부가적인 기능들과 최소 기능을 구현하는 과정에서 생긴 의문점과 TODO로 작성하고 넘어간 부분들을 해소하는 것을 목표로 하고 있습니다.

- 부가적인 기능의 예시
  - 내부 로직 처리의 로그 기능 추가
  - 로그 기능의 debugging, running mode 추가
  - 미들웨어를 통한 반복적인 검증 로직 추가
  - **[github issues](https://github.com/kaestro/ChatApplication/issues)**

- **[의문점 예시](https://github.com/kaestro/ChatApplication/blob/main/myapp/internal/Questions.md)**
  - clientManager/roomManager는 얼마나 오랫동안, 얼마나 많은 client/room를 관리할 수 있는가?
  - garbage collection을 통해 client/room을 어떻게 관리할 것인가?
  - key의 충돌 등이 일어나는 예외 사항에 대한 처리는 어떻게 할 것인가?

- **[TODO](https://github.com/kaestro/ChatApplication/blob/main/myapp/internal/TODO.md)**

### 2차 목표

1차 목표를 달성한 이후에는 **대규모 인원과 트래픽을 처리**할 수 있는 서버를 구축하는 것을 목표로 하고 있습니다.

**목표**로하고 있는 대규모 인원과 트래픽은 다음과 같이 **정의**했습니다.

```md
1. 15000명 이상의 동시 접속자 - steam 기준 인기 순위 100위의 동접자
2. 분당 7000건 이상의 메시지 전송 - 아프리카 TV 기준 채팅방이 감당할 수 없었던 트래픽
```

이는 **분당 1억건** 이상의 메시지 전송을 처리할 수 있는 서버를 구축해야 한다는 것을 의미합니다. 이를 위해 추가적으로 **도입할 계획이 있는 기술**들은 다음과 같습니다.

```md
1. message queue
2. load balancer
3. nosql
4. 부하 테스트
```

이 외에 유틸적인 측면 등에서 추가할 예정인 기능들은 다음 [문서](https://github.com/kaestro/ChatApplication/wiki/%EC%B6%94%ED%9B%84-%EC%B6%94%EA%B0%80-%EA%B0%80%EB%8A%A5%ED%95%9C-%EB%B6%80%EB%B6%84%EB%93%A4)를 참고해 주세요.

---

## 인프라 현상황 및 차후 구상

현재 진행 중인 프로젝트의 인프라 구성 및 [설계도](https://github.com/kaestro/ChatApplication/wiki/%EC%8B%9C%EC%8A%A4%ED%85%9C-%EC%84%A4%EA%B3%84%EB%8F%84)는 다음과 같습니다.

![image](https://camo.githubusercontent.com/b0ca2b60dbacab06d3aa600efaec77524fd96b74f9b7059b74288cac6c9ab486/68747470733a2f2f64726976652e676f6f676c652e636f6d2f75633f6578706f72743d646f776e6c6f61642669643d3176483557387a384d6333764a6262384e566b75384f704f687049686e6259477a)

```md
1. github actions
2. docker
3. go web server
4. postgresql
5. redis
```

차후 구상중인 인프라 구성 및 설계도는 다음과 같습니다.

![image](https://github.com/kaestro/ChatApplication/assets/32026095/5d97f107-028d-476d-803e-e64a1f86e078)

```md
1. load balancer
2. message queue
3. nosql
4. 부하 테스트
5. cloud server
```

---
