import check from 'k6';
import { HTTP_SUCCESS } from './config.js';

export class StatusChecker {
    constructor(response, method, endpoint) {
        this.response = response;
        this.method = method;
        this.endpoint = endpoint;
    }

    isSuccess() {
        check(this.response, {
            [`${this.method} ${this.endpoint} status was ${HTTP_SUCCESS}`]: (r) => {
                return r.status === HTTP_SUCCESS;
            },
        });
        return this.response.status === HTTP_SUCCESS;
    }
}