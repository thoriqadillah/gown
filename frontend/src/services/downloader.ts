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

  parseSize(size: number): string {
    if (size < this.KB) return (size / this.KB).toFixed(2) + " KB"
    if (size > this.KB && size < this.MB) return (size / this.KB).toFixed(2) + " KB"
    if (size > this.KB && size < this.GB) return (size / this.MB).toFixed(2) + " MB"
    if (size > this.GB) return (size / this.GB).toFixed(2) + " GB"

    return '0 KB'
  }
}