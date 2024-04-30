import { KEYS } from '../constants.js';
import { Body } from '../utils/body.js';
import { createJsonTypeHeader } from '../utils/header.js';
import { sendRequest } from '../utils/util.js';

export function requestSignup(userName, password, emailAddress) {
    let header = createJsonTypeHeader();
    header.setContentTypeToJSON();

    let body = new Body();
    body.addBodyData(KEYS.USERNAME, userName);
    body.addBodyData(KEYS.PASSWORD, password);
    body.addBodyData(KEYS.EMAIL_ADDRESS, emailAddress);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    responseBody = sendRequest('/signup', 'post', headerData, bodyData);
}

export function requestLogin(emailAddress, password, sessionKey=null) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    let body = new Body();
    body.addBodyData(KEYS.EMAIL_ADDRESS, emailAddress);
    body.addBodyData(KEYS.PASSWORD, password);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    responseBody = sendRequest('/login', 'post', headerData, bodyData);
}

export function requestLogout(sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    const headerData = header.getHeaderData();
    responseBody = sendRequest('/logout', 'post', headerData);
}

export function requestDeleteAccount(sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    const headerData = header.getHeaderData();
    responseBody = sendRequest('/deleteAccount', 'post', headerData);
}
