import { download } from "../../wailsjs/go/models";
import { Fetch } from "../../wailsjs/go/main/App";
import { Download } from "../../wailsjs/go/main/App";
import { useDownloads } from "../store/downloads";

export default class Downloader {
 
  async fetch(url: string): Promise<download.Download | undefined> {
    const res =  await Fetch(url) 
    res.name = res.name.replaceAll('/', '-') // parse the / to not consider it with folder

    return res
  }

  async download(toDownload: download.Download): Promise<void> {
    toDownload.date = new Date() // set the start date when we click download
    
    await Download(toDownload)
  }
}