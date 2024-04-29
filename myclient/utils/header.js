import { SESSION_KEY } from "./config";

export class Header {
    constructor(headers) {
        this.headers = headers || {};
    }

    addSessionKey(sessionKey) {
        this.addHeader(SESSION_KEY, sessionKey);
    }

    setContentTypeToJSON() {
        this.addHeader('Content-Type', 'application/json');
    }

    addHeader(key, value) {
        this.headers[key] = value;
    }

    setHeaders(newHeaders) {
        this.headers = { ...this.headers, ...newHeaders };
    }

    getHeaderData() {
        return this.headers;
    }
}