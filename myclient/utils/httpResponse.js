import http from 'k6/http';
import { BASE_URL } from './config.js';

// getResponse: HTTP 요청을 보내고 응답을 확인하는 함수입니다.
// endpoint: 요청을 보낼 엔드포인트입니다.
// method: 사용할 HTTP 메서드입니다. 기본값은 'get'입니다.
// headers: 요청에 사용할 헤더입니다.
// body: 요청 본문입니다. 기본값은 null입니다.
export class HttpResponse {
    constructor(endpoint, method = 'get', headers, body = null) {
        this.endpoint = endpoint;
        this.method = method;
        this.headers = headers;
        this.body = body;
    }

    getResponse() {
        return http[this.method](`${BASE_URL}${this.endpoint}`, this.body, { headers: this.headers });
    }
}
