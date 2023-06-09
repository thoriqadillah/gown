import { defineStore } from "pinia";
import { ref } from "vue";
import { download } from "../../wailsjs/go/models";
import { UpdateAllData, DeleteFile } from '../../wailsjs/go/store/fileStore'
import { useAlert } from "./alert";

export type Store = {
  [id: string]: download.Download
}

type DownloadType = 'image' | 'document' | 'video' | 'audio' | 'compressed'

export const useDownloads = defineStore('downloads', () => {
  const list = ref<Store>({})
  const defaults = ref<Store>({})
  const search = ref('')

  const ascName = ref(true)
  const ascDate = ref(true)
  const ascSize = ref(true)
  const ascTimeElapsed = ref(true)

  const alert = useAlert()

  const filter = (query: string) => {
    if (query.length > 0 && Object.entries(list.value).length !== Object.entries(defaults.value).length) setDefault()

    return Object.fromEntries(
      Object.entries(list.value).filter(([_, el]) => el.name.toLowerCase().includes(query.toLowerCase()))
    )
  }
  const add = async (id: string, val: download.Download) => {
    try {
      list.value[id] = val
      defaults.value[id] = val
      await UpdateAllData(list.value)
    } catch (error) {
      alert.open(error as string, 'danger')
    }
  }

  const remove = async (id: string, fromdisk: boolean) => {
    try {
      const target = list.value[id]
  
      delete list.value[id]
      delete defaults.value[id]
      
      await UpdateAllData(list.value)
      if (fromdisk) DeleteFile(target)
    } catch (error) {
      alert.open(error as string, 'danger')
    }
  }

  const updateData = async (data: Store) => {
    try {
      list.value = data
      defaults.value = data
      await UpdateAllData(data)
    } catch (error) {
      alert.open(error as string, 'danger')
    }
  }

  const setData = (data: Store) => {
    list.value = Object.fromEntries(
      Object.entries(data).sort(([,v1], [,v2]) => v1.date.localeCompare(v2.date))
    )
    defaults.value = list.value
  }
  
  const filterBy = (type: DownloadType) => {
    list.value = Object.fromEntries(
      Object.entries(defaults.value).filter(([_, el]) => el.type.name === type)
    )
  }
  
  const setDefault = () => list.value = defaults.value

  const sortByName = () => {
    ascName.value = !ascName.value
    list.value = Object.fromEntries(
      Object.entries(list.value).sort(([, v1], [, v2]) => {
        return ascName.value 
          ? v1.name.localeCompare(v2.name)
          : v2.name.localeCompare(v1.name)
      })
    )
  }

  const sortByDate = () => {
    ascDate.value = !ascDate.value
    list.value = Object.fromEntries(
      Object.entries(list.value).sort(([, v1], [, v2]) => {
        return ascDate.value 
          ? v1.date.localeCompare(v2.date)
          : v2.date.localeCompare(v1.date)
      })
    )
  }
  
  const sortBySize = () => {
    ascSize.value = !ascSize.value
    list.value = Object.fromEntries(
      Object.entries(list.value).sort(([, v1], [, v2]) => {
        return ascSize.value 
          ? v1.size - v2.size 
          : v2.size - v1.size
      })
    )
  }

  const sortByTimeElapsed = () => {
    ascTimeElapsed.value = !ascTimeElapsed.value
    list.value = Object.fromEntries(
      Object.entries(list.value).sort(([, v1], [, v2]) => {
        return ascTimeElapsed.value 
          ? v1.timeElapsed.localeCompare(v2.timeElapsed)
          : v2.timeElapsed.localeCompare(v1.timeElapsed)
      })
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
    sortByName,
    sortByDate,
    sortBySize,
    sortByTimeElapsed,
  }
})