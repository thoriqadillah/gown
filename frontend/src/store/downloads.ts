import { defineStore } from "pinia";
import { ref } from "vue";
import { download } from "../../wailsjs/go/models";
import { UpdateData, UpdateName, Delete } from '../../wailsjs/go/main/App'

export const useDownloads = defineStore('downloads', () => {
  const list = ref<download.Download[]>([])
  const defaults = ref<download.Download[]>([])
  const search = ref('')
  const toDownload = ref(new download.Download())
  const ids = ref<string[]>([])
  const names = ref<string[]>([])

  const ascName = ref(true)
  const ascDate = ref(true)
  const ascSize = ref(true)
  const ascTimeElapsed = ref(true)

  const add = (val: download.Download) => {
    list.value.unshift(val)
    ids.value.unshift(val.id)
    names.value.unshift(val.name)

    defaults.value = list.value
    toDownload.value = val
  }
  const remove = async (id: string, fromdisk: boolean) => {
    const index = ids.value.indexOf(id)
    
    const deleted = list.value.splice(index, 1)
    ids.value.splice(index, 1)
    names.value.splice(index, 1)
    
    await UpdateData(list.value)
    if (fromdisk) Delete(deleted[0].name)
    defaults.value = list.value
  }
  const setData = (data: download.Download[]) => {
    list.value = data
    defaults.value = data

    ids.value = list.value.map(el => el.id)
    names.value = list.value.map(el => el.name)
  }
  const updateData = async (data: download.Download[]) => {
    list.value = data
    defaults.value = data
    await UpdateData(data)
  }
  const updateName = async (oldval: string, newval: string, id: string) => {
    for (const el of list.value) {
      if (id == el.id) {
        await UpdateName(oldval, newval)
        break
      }
    }
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
    list.value = ascName.value ? list.value.sort((a, b) => a.name.localeCompare(b.name)) : list.value.sort((a, b) => b.name.localeCompare(a.name))
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
  const GB = MB * KB
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

    const elapsed = new Date(end.getTime() - begin.getTime())
    let s = elapsed.getSeconds() % 60
    let m = elapsed.getMinutes() % 60
    let h = ((elapsed.getMinutes() / 60) % 24)

    return `${h < 10 ? '0'+h.toFixed(0) : h.toFixed(0)}h : ${m < 10 ? '0'+m : m}m : ${s < 10 ? '0'+s : s}s`
  }

  return {
    list,
    search,
    toDownload,
    names,
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
    parseElapsedTime,
    updateName
  }
})