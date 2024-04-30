import { Key } from '../constants.js';
import { Body } from '../utils/body.js';
import { createJsonTypeHeader } from '../utils/header.js';
import { sendRequest } from '../utils/util.js';

export function requestEnterChat(emailAddress, sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);
    header.addHeader(Key.emailAddress, emailAddress);

    const body = new Body();

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    responseBody = sendRequest('/enterChat', 'get', headerData, bodyData);
}

export function requestEnterRoom(roomName, emailAddress, password, sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    let body = new Body();
    body.addBodyData(Key.roomName, roomName);
    body.addBodyData(Key.emailAddress, emailAddress);
    body.addBodyData(Key.password, password);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    responseBody = sendRequest('/enterRoom', 'post', headerData, bodyData);
}

export function createRoom(roomName, emailAddress, password, sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    let body = new Body();
    body.addBodyData(Key.roomName, roomName);
    body.addBodyData(Key.emailAddress, emailAddress);
    body.addBodyData(Key.password, password);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    responseBody = sendRequest('/createRoom', 'post', headerData, bodyData);
}

export function requestRoomList(emailAddress, password, sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    let body = new Body();
    body.addBodyData(Key.emailAddress, emailAddress);
    body.addBodyData(Key.password, password);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    responseBody = sendRequest('/getRoomList', 'get', headerData, bodyData);
}
