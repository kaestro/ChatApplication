import { check } from 'k6';
import http from 'k6/http';

const BASE_URL = 'http://localhost:8080';
const HTTP_SUCCESS = 200;

// HTTP 요청을 보내고 응답을 확인하는 함수입니다.
// endpoint: 요청을 보낼 엔드포인트입니다.
// method: 사용할 HTTP 메서드입니다. 기본값은 'get'입니다.
// headers: 요청에 사용할 헤더입니다.
// body: 요청 본문입니다. 기본값은 null입니다.
export function requestAndCheck(endpoint, method = 'get', headers, body = null) {
    let response = http[method](`${BASE_URL}${endpoint}`, body, { headers: headers });
    check(response, {
        [`${method} ${endpoint} status was ${HTTP_SUCCESS}`]: (r) => {
            return r.status === HTTP_SUCCESS;
        },
    });
}
// 새로운 헤더를 설정하고 반환하는 함수입니다.
// newHeaders: 새로 설정할 헤더를 나타내는 객체입니다. 
// 여러 개의 헤더를 한 번에 설정할 수 있습니다.
// 예: { 'Authorization': 'Bearer token', 'Custom-Header': 'Custom value' }
export function setHeaders(newHeaders) {
    return newHeaders;
}

// 새로운 본문을 설정하고 반환하는 함수입니다.
// newBody: 새로 설정할 본문을 나타내는 객체입니다. 
// 여러 개의 본문 항목을 한 번에 설정할 수 있습니다.
// 예: { username: 'newUser', password: 'newPassword', email: 'newUser@example.com' }
export function setBody(newBody) {
    return JSON.stringify(newBody);
}


export default function sampleTest() {
    const sampleHeaders = setHeaders({ 'Authorization': 'Bearer token', 'Custom-Header': 'Custom value' });
    const sampleBody = setBody({ username: 'newUser', password: 'newPassword', email: 'newUser@example.com' });

    requestAndCheck('/signup', 'post', sampleHeaders, sampleBody);
    requestAndCheck('/login', 'post', sampleHeaders, sampleBody);
    requestAndCheck('/logout', 'post', sampleHeaders);
    requestAndCheck('/deleteAccount', 'post', sampleHeaders);
    requestAndCheck('/enterChat', 'get', sampleHeaders);
    requestAndCheck('/enterRoom', 'post', sampleHeaders, sampleBody);
    requestAndCheck('/createRoom', 'post', sampleHeaders, sampleBody);
    requestAndCheck('/getRoomList', 'get', sampleHeaders);
    requestAndCheck('/ping', 'get', sampleHeaders);
}