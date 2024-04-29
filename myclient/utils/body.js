export class Body {
    constructor(bodyData) {
        // 입력값이 객체인지 확인
        if (bodyData && typeof bodyData !== 'object') {
            throw new Error('Body data must be an object');
        }
        this.bodyData = bodyData ? { ...bodyData } : {};
    }

    /**
     * 이 메서드는 JSON 형태의 데이터를 한 번에 추가하는 데 사용됩니다.
     * JSON 객체를 인수로 받아, 그 객체의 모든 키-값 쌍을 본문 데이터에 추가합니다.
     * 만약 본문 데이터에 이미 같은 키가 있다면, 새 값이 기존의 값을 덮어씁니다.
     * 
     * @param {object} jsonData - 본문 데이터에 추가할 JSON 객체입니다.
     * 
     * 사용 예시:
     * let body = new Body();
     * body.addJsonData({ 'key1': 'value1', 'key2': 'value2' });
     */
    addJsonData(jsonData) {
        for (let key in jsonData) {
            this.addBodyData(key, jsonData[key]);
        }
    }

    /**
     * 이 메서드는 요청의 본문 데이터에 키-값 쌍을 추가하는 데 사용됩니다.
     * 키와 값을 인수로 받아 기존의 본문 데이터에 추가합니다.
     * 만약 본문 데이터에 이미 같은 키가 있다면, 새 값이 기존의 값을 덮어씁니다.
     * 
     * @param {string} key - 본문 데이터에 추가할 키입니다.
     * @param {any} value - 키와 연결될 값입니다.
     * 
     * 사용 예시:
     * let body = new Body();
     * body.addBodyData('key1', 'value1');
     */
    addBodyData(key, value) {
        this.bodyData[key] = value;
    }

    getBodyData() {
        return JSON.stringify(this.bodyData);
    }
}