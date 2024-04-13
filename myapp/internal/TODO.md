# TODO

## overall

- fmt를 통해 출력되는 로그들을 logger로 변경해야 한다.
- db에 메시지가 저장하는 기능을 추가한다.
  - 이에 따라 메시지가 전송됐을때 해당 부분을 추가 테스트해야한다.

## ChatManager.go

socket 연결이 살아있는지 확인할 수 있는 기능을 추가한다.

## room.go

Debugging을 위한 debugging message들이 있는데, 이것이 debugging mode일 때는 logger로 출력되도록 하고 아닐 경우에는 출력되지 않도록 한다. 현재는 전체 주석처리 된 상태

room에 client를 3명 추가하고 연결하는데 2초로 부족한데, 이 부분에 대해 생각 필요

## clientManager.go

보유 client 갯수 제한 및 지속 시간 제한을 둘 수 있도록 변경

## session쪽

key, value 변수명이 헷갈리기 좋게 돼있어서, 이 부분이 리팩토링 필요하다.

## chat/model.go

- ChatMessage와 error message를 별개의 파일로 분리한다.
- ChatMessage의 테스트 코드를 작성한다
