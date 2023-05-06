import { defineStore } from "pinia";
import { ref } from "vue";
import { download } from "../../wailsjs/go/models";
import { UpdateData, Delete } from '../../wailsjs/go/main/App'

export type Store = {
  [id: string]: download.Download
}

export type Entries<T> = {
  [K in keyof T]: [K, T[K]];
}[keyof T][];


type DownloadType = 'image' | 'document' | 'video' | 'audio' | 'compressed'

export const useDownloads = defineStore('downloads', () => {
  const list = ref<Store>({})
  const defaults = ref<Store>({})
  const search = ref('')

  const ascName = ref(true)
  const ascDate = ref(true)
  const ascSize = ref(true)
  const ascTimeElapsed = ref(true)

  const filter = (query: string) => {
    if (query.length > 0 && Object.entries(list.value).length !== Object.entries(defaults.value).length) setDefault()
    const filtered: Store = {}
    for (const id of Object.keys(list.value)) {
      if (list.value[id].name.toLowerCase().includes(query.toLowerCase())) {
        filtered[id] = list.value[id]
      }
    }
    return filtered
  }
  const add = (id: string, val: download.Download) => {
    list.value[id] = val
    defaults.value[id] = val
  }
  const remove = async (id: string, fromdisk: boolean) => {
    const target = list.value[id]

    delete list.value[id]
    delete defaults.value[id]
    
    await UpdateData(list.value)
    if (fromdisk) Delete(target.name)
  }
  const setData = (data: Store) => {
    list.value = data
    defaults.value = data
  }
  const updateData = async (data: Store) => {
    list.value = data
    defaults.value = data
    await UpdateData(data)
  }
  
  const filterBy = (type: DownloadType) => {
    const filtered: Store = {}
    for (const id of Object.keys(defaults.value)) {
      if (defaults.value[id].type.name === type) {
        filtered[id] = defaults.value[id]
      }
    }
    list.value = filtered
  }
  
  const setDefault = () => list.value = defaults.value
  const getDefault = () => {
    return defaults.value
  }

  const sortByName = () => {
    ascName.value = !ascName.value
    if (ascName.value) {
      Object.fromEntries(
        Object.entries(list).sort(([k1, v1], [k2, v2]) => v1.name.localeCompare(v2.name))
      )
      return
    }

    Object.fromEntries(
      Object.entries(list).sort(([k1, v1], [k2, v2]) => v2.name.localeCompare(v1.name))
    )
  }

  const sortByDate = () => {
    ascDate.value = !ascDate.value
    if (ascDate.value) {
      Object.fromEntries(
        Object.entries(list).sort(([k1, v1], [k2, v2]) => v1.date.localeCompare(v2.date))
      )
      return
    }
  
    Object.fromEntries(
      Object.entries(list).sort(([k1, v1], [k2, v2]) => v2.date.localeCompare(v1.date))
    )
  }
  
  const sortBySize = () => {
    ascSize.value = !ascSize.value
    if (ascSize.value) {
      Object.fromEntries(
        Object.entries(list).sort(([k1, v1], [k2, v2]) => v1.size - v2.size)
      )
      return
    }
  
    Object.fromEntries(
      Object.entries(list).sort(([k1, v1], [k2, v2]) => v2.size - v1.size)
    )
  }

  const sortByTimeElapsed = () => {
    ascTimeElapsed.value = !ascTimeElapsed.value
    if (ascTimeElapsed.value) {
      Object.fromEntries(
        Object.entries(list).sort(([k1, v1], [k2, v2]) => v1.timeElapsed.localeCompare(v2.timeElapsed))
      )
      return
    }
  
    Object.fromEntries(
      Object.entries(list).sort(([k1, v1], [k2, v2]) => v2.timeElapsed.localeCompare(v1.timeElapsed))
    )
  }

  return {
    list,
    search,
    filter,
    add,
    remove,
    setData,
    updateData,
    filterBy,
    setDefault,
    getDefault,
    sortByName,
    sortByDate,
    sortBySize,
    sortByTimeElapsed,
  }
})