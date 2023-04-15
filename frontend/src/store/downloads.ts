import { defineStore } from "pinia";
import { ref } from "vue";
import { Download } from "../types/download";

export const useDownloads = defineStore('downloads', () => {
  const list = ref<Download[]>(new Array<Download>())
  const search = ref('')

  const add = (download: Download) => list.value.push(download)
  const remove = (download: Download) => list.value.splice(list.value.indexOf(download), 1)
  const setData = (data: Download[]) => list.value = data
  const filter = (query: string) => list.value.filter(d => d.name.toLowerCase().includes(query.toLowerCase()))
  
  return {
    list,
    search,
    add,
    remove,
    setData,
    filter
  }
})