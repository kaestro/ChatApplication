// chatConnection.js
import ws from 'k6/ws';

export class ChatConnection {
    constructor(emailAddress, sessionKey) {
        this.messages = [];
        this.lastReadMessageIndex = -1;
        this.emailAddress = emailAddress;
        this.sessionKey = sessionKey;

        this.connectionPromise = this.connect();
    }

    async connect() {
        const url = 'ws://localhost:9000';
        const params = { tags: { my_tag: 'hello' } };

        return new Promise((resolve, reject) => {
            ws.connect(url, params, (socket) => {
                this.socket = socket;
                socket.on('open', () => {
                    if (socket.readyState === WebSocket.OPEN) {
                        this.handleOpen();
                        resolve();
                    } else {
                        reject(new Error('WebSocket connection failed'));
                    }
                });
                socket.on('message', (data) => this.handleMessage(data));
                socket.on('close', () => this.handleClose());
            });
        });
    }

    handleOpen() {
        this.send(JSON.stringify({
            type: 'enterChat',
            sessionKey: this.sessionKey,
            emailAddress: this.emailAddress
        }));
    }

    handleCreateRoom(roomName, password) {
        this.send(JSON.stringify({
            type: 'createRoom',
            sessionKey: this.sessionKey,
            emailAddress: this.emailAddress,
            roomName: roomName,
            password: password
        }));
    }

    handleEnterRoom(roomName) {
        this.send(JSON.stringify({
            type: 'enterRoom',
            sessionKey: this.sessionKey,
            emailAddress: this.emailAddress,
            roomName: roomName
        }));
    }

    handleMessage(data) {
        this.messages.push(data);
    }

    handleClose() {
        this.send(JSON.stringify({
            type: 'exitChat',
            sessionKey: this.sessionKey,
            emailAddress: this.emailAddress
        }));
    }

    async sendMessage(roomName, message) {
        await this.connectionPromise;
        this.send(JSON.stringify({
            type: 'message',
            roomName: roomName,
            message: message,
            emailAddress: this.emailAddress
        }));
    }

    receiveMessage() {
        const newMessages = this.messages.slice(this.lastReadMessageIndex + 1);
        this.lastReadMessageIndex = this.messages.length - 1;
        return newMessages;
    }

    toString() {
        return `ChatConnection: { emailAddress: ${this.emailAddress},
            sessionKey: ${this.sessionKey}, messages: ${JSON.stringify(this.messages)} }`;
    }

    async send(message) {
        await this.connectionPromise;
        if (this.socket) {
            this.socket.send(message);
        } else {
            console.error('WebSocket connection is not established');
        }
    }
}
