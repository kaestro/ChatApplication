import { HttpResponse } from './httpResponse';
import { StatusChecker } from './statusChecker';
 
export function sendRequest(endpoint, method = 'get', headerData, bodyData = null) {
    let response = new HttpResponse(endpoint, method, headerData, bodyData);
    return response.getResponse();
}

export function checkIsSuccess(response, method, endpoint) {
    let checker = new StatusChecker(response, method, endpoint);
    return checker.isSuccess();
}
