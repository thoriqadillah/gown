export namespace download {
	
	export class Metadata {
	    url: string;
	    cansplit: boolean;
	    totalpart: number;
	
	    static createFrom(source: any = {}) {
	        return new Metadata(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.url = source["url"];
	        this.cansplit = source["cansplit"];
	        this.totalpart = source["totalpart"];
	    }
	}
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
	    timeElapsed: string;
	    size: number;
	    // Go type: time
	    date: any;
	    progressbar: number[];
	    status: DownloadStatus;
	    progress: number;
	    type: DownloadType;
	    metadata: Metadata;
	
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
	        this.progressbar = source["progressbar"];
	        this.status = this.convertValues(source["status"], DownloadStatus);
	        this.progress = source["progress"];
	        this.type = this.convertValues(source["type"], DownloadType);
	        this.metadata = this.convertValues(source["metadata"], Metadata);
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
	
	export class Theme {
	    textColor: string;
	    backgroundColor: string;
	
	    static createFrom(source: any = {}) {
	        return new Theme(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.textColor = source["textColor"];
	        this.backgroundColor = source["backgroundColor"];
	    }
	}
	export class Settings {
	    themes: Theme;
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
	        this.themes = this.convertValues(source["themes"], Theme);
	        this.partsize = source["partsize"];
	        this.concurrency = source["concurrency"];
	        this.maxtries = source["maxtries"];
	        this.simmultanousNum = source["simmultanousNum"];
	        this.saveLocation = source["saveLocation"];
	        this.dataLocation = source["dataLocation"];
	        this.dataFilename = source["dataFilename"];
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

