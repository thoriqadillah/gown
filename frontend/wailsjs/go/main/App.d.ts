// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {download} from '../models';
import {setting} from '../models';

export function Delete(arg1:string):Promise<void>;

export function Download(arg1:any):Promise<void>;

export function Fetch(arg1:string):Promise<any>;

export function InitData():Promise<Array<download.Download>>;

export function InitSetting():Promise<setting.Settings>;

export function StopDownload(arg1:string):Promise<void>;

export function Theme():Promise<setting.Theme>;

export function UpdateData(arg1:Array<download.Download>):Promise<void>;

export function UpdateName(arg1:string,arg2:string):Promise<void>;
