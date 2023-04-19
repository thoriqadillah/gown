import { download } from "../../wailsjs/go/models";
import { Fetch } from "../../wailsjs/go/main/App";
import { Download } from "../../wailsjs/go/main/App";

export default class Downloader {

  private KB = 1024
  private MB = this.KB * this.KB
  private GB = this.MB * this.MB

  async fetch(url: string): Promise<download.Download | undefined> {
    return await Fetch(url)
  }

  async download(toDownload: download.Download): Promise<void> {
    await Download(toDownload)
  }
}