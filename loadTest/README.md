# LoadTest Service of ChatApplication

부하테스트를 위한 서비스를 제공하며, docker에 grafana/k6가 설치되어 있어야 합니다.

k6를 이용해 테스트하고 싶은 script를 .js 형태로 작성한 뒤, runScripts.ps1 {script.js}를 실행시 results/{script_results.json}에 결과가 저장됩니다.
