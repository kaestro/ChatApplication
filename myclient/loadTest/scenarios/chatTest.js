// chatTest.js
import { check } from 'k6';
import { requestEnterChat } from '../../httpRequests/chatRequests.js';
import { requestDeleteAccount, requestLogin, requestSignup } from '../../httpRequests/userRequests.js';

const message = 'Hello, World!';
const iterations = __ENV.ITERATIONS || 100;  // Default to 100 if not provided
const userCount = __ENV.USERS || 3;  // Default to 3 if not provided

let users = [];
for (let i = 0; i < userCount; i++) {
    users.push({ userName: `user${i+1}`, password: `password${i+1}`, emailAddress: `user${i+1}@example.com`, sessionKey: null, chatConnection: null });
}

export default function() {
    for (const user of users) {
        requestSignup(user.userName, user.password, user.emailAddress);
        const loginResponse = requestLogin(user.emailAddress, user.password);
        user.sessionKey = JSON.parse(loginResponse).sessionKey;
        console.log("User session key: " + user.sessionKey);
        user.chatConnection = requestEnterChat(user.emailAddress, user.sessionKey);
    }

    for (let i = 0; i < iterations; i++) {
        for (const user of users) {
            user.chatConnection.sendMessage('roomName', message, user.emailAddress);
        }
    }

    for (const user of users) {
        const messages = user.chatConnection.receiveMessage();
        check(messages.length, { 'received messages': (val) => val === users.length * iterations });
        check(messages.every(msg => msg === message), { 'all messages are correct': (val) => val === true });
    }

    for (const user of users) {
        requestDeleteAccount(user.sessionKey);
    }
}
