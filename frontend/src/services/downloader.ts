import { download } from "../../wailsjs/go/models";
import { Fetch } from "../../wailsjs/go/main/App";
import { Download } from "../../wailsjs/go/main/App";
import { useDownloads } from "../store/downloads";

export default class Downloader {

  private KB = 1024
  private MB = this.KB * this.KB
  private GB = this.MB * this.MB

  handleDuplication(name: string): string {
    var regex = /\(([^)]+)\)/; // get number inside the parenthesis
    const downloads = useDownloads()
    
    let newname = name
    downloads.list.forEach(el => {
      if (el.name == name) {
        const matches = regex.exec(name)        
        if (matches == null) {
          let split = name.split('.')
          split[0] = split[0] + ' (1)'
          
          newname = split.join('.')
          newname = this.handleDuplication(newname)
          return
        }
        
        const number = parseInt(matches![1]) + 1
        name = name.replaceAll(matches![0], '')
        let split = name.split('.')
        split[0] = split[0] + `(${number})`
        newname = split.join('.')
        newname = this.handleDuplication(newname)
        return
      }
    })
    
    return newname
  }

  async fetch(url: string): Promise<download.Download | undefined> {
    const res =  await Fetch(url)
     // parse the / to not consider it with folder and incremental to remove name duplication
    res.name = this.handleDuplication(res.name.replaceAll('/', '-'))

    return res
  }

  async download(toDownload: download.Download): Promise<void> {
    toDownload.date = new Date() // set the start date when we click download
    
    await Download(toDownload)
  }
}