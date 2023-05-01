import { defineStore } from "pinia";
import { ref } from "vue";
import { download } from "../../wailsjs/go/models";
import { UpdateData, Delete } from '../../wailsjs/go/main/App'

export const useDownloads = defineStore('downloads', () => {
  const list = ref<download.Download[]>([])
  const defaults = ref<download.Download[]>([])
  const search = ref('')

  const ascName = ref(true)
  const ascDate = ref(true)
  const ascSize = ref(true)
  const ascTimeElapsed = ref(true)

  const add = (val: download.Download) => {
    list.value.unshift(val)
    defaults.value = list.value
  }
  const remove = async (id: string, fromdisk: boolean) => {
    const index = list.value.findIndex(el => el.id === id)
    const deleted = list.value.splice(index, 1)
    
    await UpdateData(list.value)
    if (fromdisk) Delete(deleted[0].name)
    defaults.value = list.value
  }
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

  return {
    list,
    search,
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
  }
})