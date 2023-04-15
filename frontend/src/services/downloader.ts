import { http } from "../../wailsjs/go/models";
import { Ref, ref } from "vue";
import { Fetch } from "../../wailsjs/go/main/App";
import Dialog from "./download-dialog";


export default class Downloader {

  private KB = 1024
  private MB = this.KB * this.KB
  private GB = this.MB * this.MB

  async fetch(url: string): Promise<http.Response | undefined> {
    const response = await Fetch(url)
    response.size = this.parseSize(response.size)

    return response
  }

  download() {
  }

  private parseSize(size: number): number {
    if (size < this.KB) return parseFloat((size / this.KB).toFixed(2))    
    if (size > this.KB && size < this.GB) return parseFloat((size / this.MB).toFixed(2))
    if (size > this.GB) return parseFloat((size / this.GB).toFixed(2))

    return 0
  }
}