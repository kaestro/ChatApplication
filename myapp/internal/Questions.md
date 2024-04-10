# Questions

## Chat모듈

- client.go
  - client의 sessionID를 계속해서 검증하는 작업이 반복되는 중
    - 중복되는 코드를 줄이기 위해 middleware로 분리하는 방법의 검토

- clientSession.go
  - 현재 모든 client - room 간의 session이 go를 통해 관리 되고 있다.
  - 이것이 대용량에 대해서도 문제가 없는지, 어느 정도의 규모까지 성능이 유지되는지 확인 필요

- clientManager.go
  - Question
    - How can I make sure that ClientManager won't be calling garbage collection on the Client object?
    - should I assure garbage collection won't be called
    - Should I limit the number of clients to be stored inside clientManager?
    - How does garbage collection work in Go?
    - Is making ClientManager a singleton a good idea?

- room.go
  - 현재 room.go의 AddClient는 객체를 (Client*, websocket.conn)을 받고 이를 통해 roomclienthandler를 생성하는데, 이를 roomclienthandler에서 먼저 처리하도록 해야하는가?
  - 이미 close된 room에 대한 요청이 들어올 때, 이를 모든 method에서 확인하는 코드를 반복해서 작성하는 것은 문제가 있다. 이를 해결하기 위한 방법은?

- roomManager.go
  - room을 관리하는 객체인 rooms를 slice와 map 중 무엇으로 관리하는 것이 옳은가?
    - 방이 삭제되고 추가하는 작업이 생길 경우 room들에 대한 indexing을 다시 다 새로 해야하므로 map을 사용하는 것이 좋다.

- ChatManager.go
  - 채팅과 관련한 모든 요청이 ChatManager를 통하도록 할 거면, ChatManager의 메소드를 제외하고는 전부 private으로 만들어야 하는가?
