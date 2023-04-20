import { ref, Ref } from "vue"
import { DialogState } from "../types/dialog"

export default class Dialog {

  private state: DialogState

  constructor(state: DialogState) {
    this.state = state
  }

  loading() {
    this.state.loading.value = true
  }

  done() {
    this.state.loading.value = true
    this.state.loading.value = false
  }
  
  next() {
    this.state.loaded.value = true
    this.state.onFile.value = true
    this.state.onURL.value = false
  }
  
  prev() {
    this.state.loaded.value = true
    this.state.onFile.value = false
    this.state.onURL.value = true
  }

  close() {
    this.state.onURL.value = true
    this.state.onFile.value = false
    this.state.activator.value = false
    this.state.loading.value = false
    this.state.loaded.value = false
  }
}