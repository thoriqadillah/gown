import { defineStore } from "pinia";
import { ref } from "vue";
import { download } from "../../wailsjs/go/models";
import { UpdateData} from '../../wailsjs/go/main/App'

export const useDownloads = defineStore('downloads', () => {
  const list = ref<download.Download[]>([])
  const defaults = ref<download.Download[]>([])
  const search = ref('')
  const toDownload = ref(new download.Download())

  const ascName = ref(true)
  const ascDate = ref(true)
  const ascSize = ref(true)
  const ascTimeElapsed = ref(true)

  const add = (val: download.Download) => {
    list.value.unshift(val)
    toDownload.value = val
  }
  const remove = (download: download.Download) => list.value.splice(list.value.indexOf(download), 1)
  const setData = (data: download.Download[]) => {
    list.value = data
    defaults.value = data
  }
  const updateData = async (data: download.Download[]) => {
    list.value = data
    defaults.value = data
    await UpdateData(data)
  }

  const filterByImage = () => list.value = defaults.value.filter(d => d.type.name === 'image')
  const filterByVideo = () => list.value = defaults.value.filter(d => d.type.name === 'video')
  const filterByDocument = () => list.value = defaults.value.filter(d => d.type.name === 'document')
  const filterByMusic = () => list.value = defaults.value.filter(d => d.type.name === 'audio')
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
    return ascTimeElapsed.value ? list.value.sort((a, b) => a.timeElapsed.localeCompare(b.timeElapsed)) : list.value.sort((a, b) => a.timeElapsed.localeCompare(b.timeElapsed)).reverse()
  }

  const KB = 1024
  const MB = KB * KB
  const GB = MB * MB
  const parseSize = (size: number): string => {
    if (size < KB) return (size / KB).toFixed(2) + " KB"
    if (size > KB && size < MB) return (size / KB).toFixed(2) + " KB"
    if (size > KB && size < GB) return (size / MB).toFixed(2) + " MB"
    if (size > GB) return (size / GB).toFixed(2) + " GB"

    return '0 KB'
  }

  const parseElapsedTime = (start: Date): string => {
    const begin = new Date(start)
    const end = new Date()

    const elapsed = new Date(end.getTime() - begin.getTime()).getSeconds()
    let s = elapsed % 60
    let m = (elapsed / 60) % 60
    let h = elapsed / 3600

    return `${h.toFixed(0)}h : ${m < 10 ? '0'+m.toFixed(0) : m.toFixed(0)}m : ${s < 10 ? '0'+s.toFixed(0) : s.toFixed(0)}s`
  }

  return {
    list,
    search,
    toDownload,
    add,
    remove,
    setData,
    updateData,
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
    parseSize,
    parseElapsedTime
  }
})