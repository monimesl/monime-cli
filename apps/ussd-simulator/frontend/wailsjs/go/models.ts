export namespace ussdgateway {
	
	export class ExchangeRequest {
	    networkId: string;
	    sessionId: string;
	    replyData: string;
	    initialUssdCode: string;
	
	    static createFrom(source: any = {}) {
	        return new ExchangeRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.networkId = source["networkId"];
	        this.sessionId = source["sessionId"];
	        this.replyData = source["replyData"];
	        this.initialUssdCode = source["initialUssdCode"];
	    }
	}
	export class ExchangeResponse {
	    sessionId: string;
	    terminate: boolean;
	    responseMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new ExchangeResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.sessionId = source["sessionId"];
	        this.terminate = source["terminate"];
	        this.responseMessage = source["responseMessage"];
	    }
	}

}

