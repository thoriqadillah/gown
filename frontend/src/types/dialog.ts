import { Ref } from "vue"

export type DialogState = {
  activator: Ref<boolean>
  loading: Ref<boolean>
  loaded: Ref<boolean>
  onURL: Ref<boolean>
  onFile: Ref<boolean>
}