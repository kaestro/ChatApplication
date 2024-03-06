myapp
├── api
│   ├── handler
│   │   └── user
│   ├── models
│   └── service
│       └── user
├── internal
│   ├── db
│   ├── session
│   └── password
├── pkg
├── scripts
├── tests
└── main.go

###

위와 같이 구성된 프로젝트 상에 chatroom을 추가하려해. 어떤 모듈의 형태로 작성하면 좋을까?

chatroom은 여러 명의 사람이 socket을 통해 연결돼있으면서, observer의 형태로 chatroom에 변화가 있을 경우에 내부에 메시지를 전달받도록 하고 싶어.

이 때 메시지는 session module을 쓸 생각이야.

채팅에 저장할 내용을 mongodb에 저장할 생각이고.

현재 서버의 스택은 다음 docker-compose를 통해 확인할 수 있어.

###

version: '3'
services:
  postgresql:
    image: postgres:latest
    environment:
      - POSTGRES_DB=postgres
      - POSTGRES_PASSWORD=rootpassword
    ports:
      - "5432:5432"
    volumes:
      - postgresql_data:/var/lib/postgresql/data

  redis:
    image: redis:latest
    command: redis-server --requirepass redisPassword
    ports:
      - "6379:6379"

  mongodb:
    image: mongo:latest
    ports:
      - "27017:27017"
    volumes:
      - mongodb_data:/data/db
  
  main_server:
    image: main_server:latest
    ports:
      - "8080:8080"
    depends_on:
      - postgresql
      - redis
      - mongodb
    environment:
      - REDIS_ADDR=redis:6379
      - DB_URL=postgres://postgres:rootpassword@postgresql:5432/postgres?sslmode=disable

volumes:
  postgresql_data:
  mongodb_data:

###

아직 mongodb_module은 없다고 할 때 mongo_db의 모듈은 또 어디에다가 놓으면 좋을까?

###

어디서부터 시작하면 좋을지 아이디어 및 프로젝트 구성 제안해줘

GitHub Copilot: 이 프로젝트에 채팅 기능을 추가하려면, 다음과 같은 구조로 모듈을 추가하는 것이 좋을 것 같습니다:

1. `api/handler/chatroom`: 이 모듈에서는 채팅방에 관련된 HTTP 요청을 처리합니다. 예를 들어, 채팅방 생성, 채팅방 목록 조회, 채팅방 입장 및 퇴장 등의 요청을 처리할 수 있습니다.

2. `api/service/chatroom`: 이 모듈에서는 채팅방 관련 비즈니스 로직을 처리합니다. 예를 들어, 채팅 메시지를 보내는 기능, 채팅방에 참가하는 사용자를 관리하는 기능 등을 구현할 수 있습니다.

3. `internal/db/mongodb`: MongoDB와의 연결 및 쿼리를 처리하는 모듈입니다. 채팅 메시지를 저장하고 조회하는 기능을 구현할 수 있습니다.

4. `internal/session`: 이미 존재하는 세션 모듈을 사용하여, 사용자의 로그인 세션을 관리하고, 채팅방에 참가한 사용자를 식별할 수 있습니다.

5. `internal/socket`: WebSocket 연결을 관리하는 모듈입니다. 사용자가 채팅방에 입장하면 WebSocket 연결을 생성하고, 채팅 메시지가 도착하면 해당 채팅방에 참가한 모든 사용자에게 메시지를 전송할 수 있습니다.

이렇게 모듈을 구성하면, 각 기능이 잘 분리되어 있어서 유지보수가 쉬워질 것입니다. 또한, 각 모듈이 하나의 책임만 가지므로 단일 책임 원칙(Single Responsibility Principle)을 잘 지킬 수 있습니다.

이제 이러한 모듈을 구현하기 위한 첫 번째 단계는, 각 모듈의 인터페이스를 정의하는 것입니다. 이렇게 하면, 각 모듈이 어떤 기능을 제공해야 하는지 명확하게 알 수 있고, 이를 바탕으로 구현을 시작할 수 있습니다.


---

## TODO

### mongoDB 관련

mongodb부터 작성중.

mongodb의 상세 인터페이스 설계 필요.

mongodb에서 사용하는 model의 위치를 api에 두는 것이 맞는지에 대한 고려 필요.

---

### 채팅방 목록 관련

1. 데이터 모델링 => user, chatroom
2. api 설계 => chatroom list, chatroom create, chatroom delete
3. 세션 관리 => chatroom에 참가한 사용자 관리
4. 채팅 메시지 저장 => mongodb에 저장
5. 채팅 메시지 캐싱 => redis에 저장
6. 채팅 메시지 전송 => WebSocket을 통해 전송