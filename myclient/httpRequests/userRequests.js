import { KEYS } from '../constants.js';
import { Body } from '../utils/body.js';
import { Header } from '../utils/header.js';
import { sendRequest } from '../utils/util.js';

export function requestSignup(userName, password, emailAddress) {
    let header = new Header();
    header.setContentTypeToJSON();

    let body = new Body();
    body.addBodyData(KEYS.USERNAME, userName);
    body.addBodyData(KEYS.PASSWORD, password);
    body.addBodyData(KEYS.EMAIL_ADDRESS, emailAddress);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    responseBody = sendRequest('/signup', 'post', headerData, bodyData);
}

export function requestLogin(emailAddress, password) {
    let header = new Header();
    header.setContentTypeToJSON();
    header.addSessionKey(responseBody[KEYS.SESSION_KEY]);

    let body = new Body();
    body.addBodyData(KEYS.EMAIL_ADDRESS, emailAddress);
    body.addBodyData(KEYS.PASSWORD, password);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    responseBody = sendRequest('/login', 'post', headerData, bodyData);
}