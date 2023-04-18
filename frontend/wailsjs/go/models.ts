export namespace download {
	
	export class DownloadType {
	    name: string;
	    icon: string;
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new DownloadType(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.icon = source["icon"];
	        this.color = source["color"];
	    }
	}
	export class DownloadStatus {
	    name: string;
	    icon: string;
	    color: string;
	
	    static createFrom(source: any = {}) {
	        return new DownloadStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.icon = source["icon"];
	        this.color = source["color"];
	    }
	}
	export class Download {
	    id: string;
	    name: string;
	    timeElapsed: number;
	    size: number;
	    // Go type: time
	    date: any;
	    status: DownloadStatus;
	    type: DownloadType;
	
	    static createFrom(source: any = {}) {
	        return new Download(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.name = source["name"];
	        this.timeElapsed = source["timeElapsed"];
	        this.size = source["size"];
	        this.date = this.convertValues(source["date"], null);
	        this.status = this.convertValues(source["status"], DownloadStatus);
	        this.type = this.convertValues(source["type"], DownloadType);
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
	export class DownloadData {
	    response?: http.Response;
	    data: Download;
	
	    static createFrom(source: any = {}) {
	        return new DownloadData(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.response = this.convertValues(source["response"], http.Response);
	        this.data = this.convertValues(source["data"], Download);
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
	    dataLocation: string;
	    dataFilename: string;
	
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
	        this.dataLocation = source["dataLocation"];
	        this.dataFilename = source["dataFilename"];
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

