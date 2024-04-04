# Chat Application

## 개요

백엔드 개발에서 요구되는 능력들 중 **실시간 통신, 데이터베이스 및 API 서버 구축, 인증 및 보안** 등의 능력 향상을 목적으로 한 채팅 어플리케이션입니다.

**기초적인 채팅** 메시지 전송 및 수신, 사용자 인증, 사용자 정보 관리에서 시작해, 최종적으로는 **대규모 인원의 동시 접속 및 트래픽을 처리**할 수 있는 서버를 구축하는 것을 목표로 하고 있습니다.

---

## 목차

* 팀원 구성 및 협업 방식
* MVP(Mininum Viable Product)
* 기술 스택
* 확장 계획
* 인프라 AS-IS 및 TO-BE

---

## 팀원 구성 및 협업 방식

총 3명의 팀원([TLOWAC](https://github.com/TLOWAC), [neuma](https://github.com/neuma573), [kaestro](https://github.com/kaestro))으로 구성되어 있으며, 원활한 협업을 위해 노력하고 있습니다.

이를 위해 **[notion](https://www.notion.so/lthek55/Golang-Chat-Backend-f308886d9d834d1a9059d42545066c46), github actions, trunk based development** 등을 통해 생산성과 협업능력을 향상하고 있습니다. 세부적으로는 매 주 1회 **회의**를 통해 진행 상황 공유 및 방향성을 정하였으며, **unit test**를 포함해서 직접적인 **push가 불가능**한 **main branch**에 merge하기 전에는 **code review**를 거치도록 하였습니다.

---

## MVP(Mininum Viable Product)

현재 목표로 하는 채팅 프로그램의 가장 **기본적인 기능**은 다음과 같습니다.

```md
1. 로그인 및 회원가입
2. 채팅방 생성 및 입장
3. 채팅 메시지 전송 및 수신
4. 채팅 기록 저장
```

---

## 기술 스택

현재 진행 중인 프로젝트의 기술 스택은 다음과 같습니다.

* 웹서버
  * goLang
  * gin
* 데이터베이스
  * postgresql
* 세션 관리
  * redis
* 컨테이너
  * docker

이와 관련해서 선택한 이유는 [다음](https://kaestro.github.io/%EA%B0%9C%EB%B0%9C%EC%9D%BC%EC%A7%80/2024/03/19/Chat-Application-5%EC%A3%BC%EC%B0%A8-review.html)에서 확인하실 수 있습니다.

---

## 확장 계획

**목표**로하고 있는 대규모 인원과 트래픽은 다음과 같이 **정의**했습니다.

```md
1. 15000명 이상의 동시 접속자
2. 분당 7000건 이상의 메시지 전송
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

## 인프라 현상황 및 확장 계획

현재 진행 중인 프로젝트의 인프라 구성 및 [설계도](https://github.com/kaestro/ChatApplication/wiki/%EC%8B%9C%EC%8A%A4%ED%85%9C-%EC%84%A4%EA%B3%84%EB%8F%84)는 다음과 같습니다.

![image](https://drive.google.com/uc?export=download&id=1vH5W8z8Mc3vJbb8NVku8OpOhpIhnbYGz)

```md
1. github actions
2. docker
3. 개인 서버
4. go web server
5. postgresql
6. redis
```

확장 계획이 완성된 이후의 인프라 구성 및 설계도는 다음과 같습니다.

![image](https://github.com/kaestro/ChatApplication/assets/32026095/5d97f107-028d-476d-803e-e64a1f86e078)

```md
1. load balancer
2. message queue
3. nosql
4. 부하 테스트
5. cloud server
```

---
