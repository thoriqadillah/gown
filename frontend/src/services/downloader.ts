import { download, http } from "../../wailsjs/go/models";
import { Fetch } from "../../wailsjs/go/main/App";
import { Download } from "../../wailsjs/go/main/App";

export default class Downloader {

  private KB = 1024
  private MB = this.KB * this.KB
  private GB = this.MB * this.MB

  async fetch(url: string): Promise<download.DownloadData | undefined> {
    const response = await Fetch(url)

    return response
  }

  async download(res: http.Response): Promise<void> {
    await Download(res)
  }

  parseSize(size: number): string {
    if (size < this.KB) return (size / this.KB).toFixed(2) + " KB"
    if (size > this.KB && size < this.MB) return (size / this.KB).toFixed(2) + " KB"
    if (size > this.KB && size < this.GB) return (size / this.MB).toFixed(2) + " MB"
    if (size > this.GB) return (size / this.GB).toFixed(2) + " GB"

    return '0 KB'
  }
}