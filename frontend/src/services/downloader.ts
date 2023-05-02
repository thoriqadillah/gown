import { download } from "../../wailsjs/go/models";
import { DeleteTempfile, Fetch } from "../../wailsjs/go/main/App";
import { Download } from "../../wailsjs/go/main/App";
import { useDownloads } from "../store/downloads";
import { EventsEmit, EventsOn } from "../../wailsjs/runtime/runtime";

export default class Downloader {

  private static instance: Downloader
  private store = useDownloads()

  public static service(): Downloader {
    if (!Downloader.instance) {
      Downloader.instance = new Downloader()
    }

    return Downloader.instance
  }

  handleDuplication(name: string): string {
    var regex = /\(([^)]+)\)/; // get number inside the parenthesis
    
    let newname = name  
    if (this.store.list.findIndex(el => el.name === name) > -1) {
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
    EventsEmit("stop", id)
    const target = this.store.list[this.store.list.findIndex(el => el.id === id)]
    
    EventsOn("total-bytes", async (...data) => {
      const [index, downloaded] = data
      target.chunks[index].downloaded = downloaded

      target.status.icon = 'mdi-pause-circle-outline'
      target.status.color = ''
      target.status.name = 'Paused'
      
      await this.store.updateData(this.store.list)
    })
    console.log(target);
    // save the range based on file size?

  }

  async resume(id: string) {

  }

  async stop(id: string, deleteTempfile: boolean) {
    EventsEmit("stop", id)
    const target = this.store.list[this.store.list.findIndex(el => el.id === id)]
    target.status.icon = 'mdi-stop-circle-outline'
    target.status.color = 'warning'
    target.status.name = 'Canceled'
    
    await this.store.updateData(this.store.list)
    if (deleteTempfile) DeleteTempfile(target)
  }
}