# Questions

## Chat모듈

- client.go
  - client의 sessionID를 계속해서 검증하는 작업이 반복되는 중
    - 중복되는 코드를 줄이기 위해 middleware로 분리하는 방법의 검토

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
