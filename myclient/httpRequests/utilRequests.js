export function ping() {
    const header = createJsonTypeHeader();
    const headerData = header.getHeaderData();
    responseBody = sendRequest('/ping', 'get', headerData);
}
