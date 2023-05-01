import { download } from "../../wailsjs/go/models";
import { Fetch } from "../../wailsjs/go/main/App";
import { Download } from "../../wailsjs/go/main/App";
import { useDownloads } from "../store/downloads";
import { EventsEmit } from "../../wailsjs/runtime/runtime";

export default class Downloader {

  private static instance: Downloader
  private downloads = useDownloads()

  public static service(): Downloader {
    if (!Downloader.instance) {
      Downloader.instance = new Downloader()
    }

    return Downloader.instance
  }

  handleDuplication(name: string): string {
    var regex = /\(([^)]+)\)/; // get number inside the parenthesis
    
    let newname = name  
    if (this.downloads.list.findIndex(el => el.name === name) > -1) {
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
  async pause(id: string) {

  }

  async resume(id: string) {

  }

  async stop(id: string) {
    EventsEmit("stop", id)
  }
}