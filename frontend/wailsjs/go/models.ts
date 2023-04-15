export namespace http {
	
	export class Response {
	    url: string;
	    size: number;
	    contentType: string;
	    cansplit: boolean;
	    totalpart: number;
	    filename: string;
	    settings?: setting.Settings;
	
	    static createFrom(source: any = {}) {
	        return new Response(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.size = source["size"];
	        this.contentType = source["contentType"];
	        this.cansplit = source["cansplit"];
	        this.totalpart = source["totalpart"];
	        this.filename = source["filename"];
	        this.settings = this.convertValues(source["settings"], setting.Settings);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

export namespace setting {
	
	export class Settings {
	    textColor: string;
	    backgroundColor: string;
	    partsize: number;
	    concurrency: number;
	    maxtries: number;
	    simmultanousNum: number;
	    saveLocation: string;
	
	    static createFrom(source: any = {}) {
	        return new Settings(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.textColor = source["textColor"];
	        this.backgroundColor = source["backgroundColor"];
	        this.partsize = source["partsize"];
	        this.concurrency = source["concurrency"];
	        this.maxtries = source["maxtries"];
	        this.simmultanousNum = source["simmultanousNum"];
	        this.saveLocation = source["saveLocation"];
	    }
	}
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

