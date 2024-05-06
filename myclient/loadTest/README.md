# LoadTest Service of ChatApplication

부하테스트를 위한 서비스를 제공하며, docker에 grafana/k6가 설치되어 있어야 합니다.

k6를 이용해 테스트하고 싶은 script를 .js 형태로 작성한 뒤, runScripts.ps1 {script.js}를 실행시 results/{script_results.json}에 결과가 저장됩니다.

## 사용 방법

```powershell
k6 run --env USERS=5 --env ITERATIONS=200 .\loadTest\scenarios\chatTest.js
```

다음과 같이 유저의 숫자와 동일한 메시지를 반복해서 보내는 횟수를 설정할 수 있습니다.
