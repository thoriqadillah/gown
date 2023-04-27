import { download } from "../../wailsjs/go/models";
import { Fetch } from "../../wailsjs/go/main/App";
import { Download } from "../../wailsjs/go/main/App";
import { useDownloads } from "../store/downloads";

export default class Downloader {

  handleDuplication(name: string): string {
    var regex = /\(([^)]+)\)/; // get number inside the parenthesis
    const downloads = useDownloads()
    
    let newname = name  
    if (downloads.names.indexOf(name) > -1) {
      const matches = regex.exec(name)        
      if (matches == null) {
        let split = name.split('.')
        if (split.length > 2) {
          split[split.length - 2] += ' (1)'
        } else {
          split[0] += ' (1)'
        }
        
        newname = split.join('.')
        newname = this.handleDuplication(newname)
        return newname
      }
      
      const number = parseInt(matches![1]) + 1
      name = name.replaceAll(matches![0], '')
      let split = name.split('.')
      if (split.length > 2) {
        split[split.length - 2] += `(${number})`
      } else {
        split[0] += `(${number})`
      }
      newname = split.join('.')
      newname = this.handleDuplication(newname)
      return newname
    }
    
    return newname
  }
 
  async fetch(url: string): Promise<download.Download | undefined> {
    const res =  await Fetch(url) 
    res.name = res.name.replaceAll('/', '-') // parse the / to not consider it with folder
    res.name = this.handleDuplication(res.name) // handle duplication

    return res
  }

  async download(toDownload: download.Download): Promise<void> {
    toDownload.date = new Date() // set the start date when we click download
    await Download(toDownload)
  }
  
  // TODO: implement resume download
}