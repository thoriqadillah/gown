import { defineStore } from "pinia";
import { ref, computed } from "vue";
import { Download } from "../types/download";

export const useDownloads = defineStore('downloads', () => {
  const list = ref<Download[]>(new Array<Download>())
  const defaults = ref<Download[]>(new Array<Download>())
  const search = ref('')

  const ascName = ref(true)
  const ascDate = ref(true)
  const ascSize = ref(true)
  const ascTimeElapsed = ref(true)

  const add = (download: Download) => list.value.push(download)
  const remove = (download: Download) => list.value.splice(list.value.indexOf(download), 1)
  const setData = (data: Download[]) => {
    list.value = data
    defaults.value = data
  }

  const filterByImage = () => list.value = defaults.value.filter(d => d.type.name === 'image')
  const filterByVideo = () => list.value = defaults.value.filter(d => d.type.name === 'video')
  const filterByDocument = () => list.value = defaults.value.filter(d => d.type.name === 'document')
  const filterByMusic = () => list.value = defaults.value.filter(d => d.type.name === 'music')
  const filterByCompressed = ()  => list.value = defaults.value.filter(d => d.type.name === 'compressed')
  const setDefault = ()  => list.value = defaults.value
  const filter = (query: string) =>  {
    if (query.length > 0 && list.value.length !== defaults.value.length) setData(defaults.value)
    return list.value.filter(d => d.name.toLowerCase().includes(query.toLowerCase()))
  }

  const sortByName = () => {
    ascName.value = !ascName.value
    return ascName.value ? list.value.sort((a, b) => a.name.localeCompare(b.name)) : list.value.sort((a, b) => a.name.localeCompare(b.name)).reverse()
  }
  const sortByDate = () => {
    ascDate.value = !ascDate.value
    return ascDate.value ? list.value.sort((a, b) => a.date.localeCompare(b.date)) : list.value.sort((a, b) => a.date.localeCompare(b.date)).reverse()
  }
  const sortBySize = () => {
    ascSize.value = !ascSize.value
    return ascSize.value ? list.value.sort((a, b) => a.size - b.size) : list.value.sort((a, b) => a.size - b.size).reverse()
  }
  const sortByTimeElapsed = () => {
    ascTimeElapsed.value = !ascTimeElapsed.value
    return ascTimeElapsed.value ? list.value.sort((a, b) => a.timeElapsed - b.timeElapsed) : list.value.sort((a, b) => a.timeElapsed - b.timeElapsed).reverse()
  }

  return {
    list,
    search,
    add,
    remove,
    setData,
    filter,
    filterByImage,
    filterByVideo,
    filterByDocument,
    filterByMusic,
    filterByCompressed,
    setDefault,
    sortByName,
    sortByDate,
    sortBySize,
    sortByTimeElapsed,
  }
})