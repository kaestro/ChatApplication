import { ChatConnection } from "../httpRequests/chatConnection.js";
import { Body } from '../utils/body.js';
import { Key } from '../utils/constants.js';
import { createJsonTypeHeader } from '../utils/header.js';
import { sendRequest } from '../utils/util.js';

export function requestEnterChat(emailAddress, sessionKey) {
    return new ChatConnection(emailAddress, sessionKey);
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
    const responseBody = sendRequest('/enterRoom', 'post', headerData, bodyData);
    return responseBody;
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
    const responseBody = sendRequest('/createRoom', 'post', headerData, bodyData);
    return responseBody;
}

export function requestRoomList(emailAddress, password, sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    let body = new Body();
    body.addBodyData(Key.emailAddress, emailAddress);
    body.addBodyData(Key.password, password);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    const responseBody = sendRequest('/getRoomList', 'get', headerData, bodyData);
    return responseBody;
}

export function requestSendMessage(roomName, message, emailAddress, sessionKey) {
    let header = createJsonTypeHeader();
    header.addSessionKey(sessionKey);

    let body = new Body();
    body.addBodyData(Key.roomName, roomName);
    body.addBodyData(Key.message, message);
    body.addBodyData(Key.emailAddress, emailAddress);

    const headerData = header.getHeaderData();
    const bodyData = body.getBodyData();
    const responseBody = sendRequest('/sendMessage', 'post', headerData, bodyData);
    return responseBody;
}