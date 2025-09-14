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
