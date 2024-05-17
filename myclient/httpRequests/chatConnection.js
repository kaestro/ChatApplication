// chatConnection.js
import ws from 'k6/ws';

export class ChatConnection {
    constructor(emailAddress, sessionKey) {
        this.messages = [];
        this.lastReadMessageIndex = -1;
        this.emailAddress = emailAddress;
        this.sessionKey = sessionKey;

        this.socket = null;
        this.connectionPromise = this.connect();
    }

    async connect() {
        return new Promise(async (resolve, reject) => {
            try {
                await this.handleOpen();
                if (this.socket && this.socket.readyState === WebSocket.OPEN) {
                    console.log('WebSocket connection established');
                    resolve();
                } else {
                    throw new Error('Failed to establish WebSocket connection');
                }
            } catch (error) {
                reject(error);
            }
        });
    }

    async handleOpen() {
        try {
            console.log("connecting to websocket server...");
            const response = await this.enterChat();
            if (response.status === 200 && response.data.message === 'entered room successfully') {
                this.socket = new websocket('ws://localhost:9000');
                this.socket.onopen = () => console.log('websocket connection established');
                this.socket.onmessage = (event) => this.handlemessage(event.data);
                this.socket.onclose = () => this.handleclose();
            } else {
                console.error('failed to enter chat room. response:', response);
                throw new error('failed to enter chat room');
            }
        } catch (error) {
            console.error('error during enterchat:', error);
            throw error;
        }
    }

    async enterChat() {
        const url = `ws://localhost:8080/enterChat`;  // 수정된 부분

        try {
            const response = ws.connect(url, {  // 수정된 부분
                headers: {
                    'Session-Key': this.sessionKey,
                    'emailAddress': this.emailAddress
                }
            }, function(socket) {
                socket.on('open', () => console.log('Connection opened!'));
                socket.on('message', (data) => console.log('Message received: ', data));
                socket.on('close', () => console.log('Connection closed!'));
                socket.on('error', (error) => console.log('Error: ', error));
            });

            // Handle response here
            return response;
        } catch (error) {
            console.error('Error during enterChat:', error.message);
            console.error('Error stack:', error.stack);
        }
    }

    handleCreateRoom(roomName, password) {
        const url = `http://localhost:8080/createRoom`;

        const payload = JSON.stringify({
            sessionKey: this.sessionKey,
            emailAddress: this.emailAddress,
            roomName: roomName,
            password: password
        });

        const params =  { headers: { 'Content-Type': 'application/json' } };

        http.post(url, payload, params);
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
