export namespace ussdgateway {
	
	export class ReplyRequest {
	    sessionId: string;
	    replyMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new ReplyRequest(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.sessionId = source["sessionId"];
	        this.replyMessage = source["replyMessage"];
	    }
	}
	export class ReplyResponse {
	    sessionId: string;
	    terminate: boolean;
	    responseMessage: string;
	
	    static createFrom(source: any = {}) {
	        return new ReplyResponse(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.sessionId = source["sessionId"];
	        this.terminate = source["terminate"];
	        this.responseMessage = source["responseMessage"];
	    }
	}

}

