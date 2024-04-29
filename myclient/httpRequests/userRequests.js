import { sendRequest, setBody, setHeaders } from '../utils/util.js';

export function requestSignup(userName, password, emailAddress) {
    const headers = setHeaders({ 'Content-Type': 'application/json' });
    const body = setBody({ userName: userName, password: password, emailAddress: emailAddress });
    responseBody = sendRequest('/signup', 'post', headers, body);
}