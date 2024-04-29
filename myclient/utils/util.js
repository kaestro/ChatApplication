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
// 예: { userName: 'newUser', password: 'newPassword', emailAddress: 'newUser@example.com' }
export function setBody(newBody) {
    return JSON.stringify(newBody);
}

export function sendRequest(endpoint, method = 'get', headers, body = null) {
    let response = new HttpResponse(endpoint, method, headers, body);
    return response.getResponse();
}

export function checkIsSuccess(response, method, endpoint) {
    let checker = new StatusChecker(response, method, endpoint);
    return checker.isSuccess();
}