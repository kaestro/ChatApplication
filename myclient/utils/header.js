import { KEYS } from "./constants.js";

export class Header {
    constructor(headers) {
        this.headers = headers || {};
    }

    addSessionKey(sessionKey) {
        this.addHeader(KEYS.SESSION_KEY, sessionKey);
    }

    setContentTypeToJSON() {
        this.addHeader('Content-Type', 'application/json');
    }

    addHeader(key, value) {
        this.headers[key] = value;
    }

    setHeaders(newHeaders) {
        this.headers = Object.assign({}, this.headers, newHeaders);
    }

    getHeaderData() {
        return this.headers;
    }

}

export function createJsonTypeHeader() {
    let jsonHeader = new Header();
    jsonHeader.setContentTypeToJSON();
    return jsonHeader;
}