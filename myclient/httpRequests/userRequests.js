import { Body } from '../utils/body.js';
import { KEYS } from '../utils/constants.js';
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
    let responseBody = sendRequest('/signup', 'post', headerData, bodyData);

    return responseBody;
}

export function requestLogin(emailAddress, password, sessionKey=null) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    let body = new Body();
    body.addBodyData(KEYS.EMAIL_ADDRESS, emailAddress);
    body.addBodyData(KEYS.PASSWORD, password);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    let responseBody = sendRequest('/login', 'post', headerData, bodyData);
    let responseString = JSON.stringify(responseBody);

    console.log(responseString);

    return responseString;
}

export function requestLogout(sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    const headerData = header.getHeaderData();
    let responseBody = sendRequest('/logout', 'post', headerData);

    return responseBody;
}

export function requestDeleteAccount(sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    const headerData = header.getHeaderData();
    let responseBody = sendRequest('/deleteAccount', 'post', headerData);

    return responseBody;
}
