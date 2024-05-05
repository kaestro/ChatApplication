import { requestSignup } from '../httpRequests/userRequests.js';
import { setHeaders } from '../utils/header.js';
import { setBody } from '../utils/util.js';

export default function sampleTest() {
    const sampleHeaders = setHeaders({ 'Authorization': 'Bearer token', 'Custom-Header': 'Custom value' });
    const sampleBody = setBody({ userName: 'newUser', password: 'newPassword', emailAddress: 'newUser@example.com' });

    response = requestSignup('newUser', 'newPassword', 'newEmail@gmail.com');
}
