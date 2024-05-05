// chatTest.js
import { requestEnterChat } from '../../httpRequests/chatRequests.js';
import { requestDeleteAccount, requestLogin, requestLogout, requestSignup } from '../../httpRequests/userRequests.js';

describe('Chat test', () => {
    const users = [
        { userName: 'user1', password: 'password1', emailAddress: 'user1@example.com', sessionKey: null, chatConnection: null },
        { userName: 'user2', password: 'password2', emailAddress: 'user2@example.com', sessionKey: null, chatConnection: null },
        { userName: 'user3', password: 'password3', emailAddress: 'user3@example.com', sessionKey: null, chatConnection: null }
    ];
    const message = 'Hello, World!';

    beforeAll(async () => {
        for (const user of users) {
            await requestSignup(user.userName, user.password, user.emailAddress);
            const loginResponse = await requestLogin(user.emailAddress, user.password);
            user.sessionKey = loginResponse.sessionKey;
            user.chatConnection = requestEnterChat(user.emailAddress, user.sessionKey);
        }
    });

    it('should send and receive 100 messages', async () => {
        for (let i = 0; i < 100; i++) {
            for (const user of users) {
                user.chatConnection.sendMessage('roomName', message, user.emailAddress);
            }
        }

        for (const user of users) {
            const messages = user.chatConnection.receiveMessage();
            expect(messages.length).toBe(users.length * 100);
            expect(messages.every(msg => msg === message)).toBe(true);
        }
    });

    afterAll(async () => {
        for (const user of users) {
            await requestLogout(user.sessionKey);
            await requestDeleteAccount(user.sessionKey);
        }
    });
});
