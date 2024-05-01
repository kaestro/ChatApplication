// chatConnection.js
import WebSocket from 'ws';

export class ChatConnection {
    constructor(emailAddress, sessionKey) {
        this.ws = new WebSocket('ws://localhost:8080');
        this.messages = [];
        this.lastReadMessageIndex = -1;

        this.ws.on('open', () => {
            this.ws.send(JSON.stringify({
                type: 'enterChat',
                sessionKey: sessionKey,
                emailAddress: emailAddress
            }));
        });

        this.ws.on('message', (data) => {
            this.messages.push(data);
        });
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