import { requestSignup } from '../../myclient/httpRequests/userRequests.js';
import { setBody, setHeaders } from '../../myclient/utils/util.js';

export default function sampleTest() {
    const sampleHeaders = setHeaders({ 'Authorization': 'Bearer token', 'Custom-Header': 'Custom value' });
    const sampleBody = setBody({ userName: 'newUser', password: 'newPassword', emailAddress: 'newUser@example.com' });

    response = requestSignup('newUser', 'newPassword', 'newEmail@gmail.com');
}