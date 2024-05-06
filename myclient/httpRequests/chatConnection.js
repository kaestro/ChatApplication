// chatConnection.js
import { WebSocket } from 'k6/ws';

export class ChatConnection {
    constructor(emailAddress, sessionKey) {
        this.ws = new WebSocket('ws://localhost:9000');
        this.messages = [];
        this.lastReadMessageIndex = -1;

        this.ws.on('open', () => this.handleOpen(emailAddress, sessionKey));
        this.ws.on('message', (data) => this.handleMessage(data));
        this.ws.on('close', () => this.handleClose(emailAddress, sessionKey));
    }

    handleOpen(emailAddress, sessionKey) {
        this.ws.send(JSON.stringify({
            type: 'enterChat',
            sessionKey: sessionKey,
            emailAddress: emailAddress
        }));
    }

    handleMessage(data) {
        this.messages.push(data);
    }

    handleClose(emailAddress, sessionKey) {
        this.ws.send(JSON.stringify({
            type: 'exitChat',
            sessionKey: sessionKey,
            emailAddress: emailAddress
        }));
    }

    sendMessage(roomName, message, emailAddress) {
        if (this.ws.readyState === WebSocket.OPEN) {
            this.ws.send(JSON.stringify({
                type: 'message',
                roomName: roomName,
                message: message,
                emailAddress: emailAddress
            }));
        }
    }

    receiveMessage() {
        const newMessages = this.messages.slice(this.lastReadMessageIndex + 1);
        this.lastReadMessageIndex = this.messages.length - 1;
        return newMessages;
    }
}
