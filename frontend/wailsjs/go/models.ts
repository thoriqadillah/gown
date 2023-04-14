export namespace http {
	
	export class Response {
	    textColor: string;
	    backgroundColor: string;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.textColor = source["textColor"];
	        this.backgroundColor = source["backgroundColor"];
	    }
	}

}

export namespace setting {
	
	export class Themes {
	    textColor: string;
	    backgroundColor: string;
	
	    static createFrom(source: any = {}) {
	        return new Themes(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.textColor = source["textColor"];
	        this.backgroundColor = source["backgroundColor"];
	    }
	}

}

