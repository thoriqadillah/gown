import { defineStore } from "pinia";
import { ref } from "vue";
import { setting } from "../../wailsjs/go/models";
import { DefaultSetting, UpdateSetting } from "../../wailsjs/go/store/fileStore";
import { useAlert } from "./alert";

export const useSettings = defineStore('settings', () => {
  const partsize = ref(0)
  const concurrency = ref(0)
  const maxtries = ref(0)
  const simmultanousNum = ref(0)
  const saveLocation = ref('')
  const dataLocation = ref('')
  const dataFilename = ref('')

  const instance = ref(new setting.Settings())
  const alert = useAlert()

  const init = (s: setting.Settings) => {
    partsize.value = s.partsize
    concurrency.value = s.concurrency
    maxtries.value = s.maxtries
    simmultanousNum.value = s.simmultanousNum
    saveLocation.value = s.saveLocation
    dataLocation.value = s.dataLocation
    dataFilename.value = s.dataFilename

    instance.value = s
  }

  const backDefault = async () => instance.value = await DefaultSetting()

  const update = async () => {
    try {
      await UpdateSetting(instance.value)
    } catch (error) {
      alert.open(error as string, 'danger')
    }
  }
  
  return {
    partsize,
    concurrency,
    maxtries,
    simmultanousNum,
    saveLocation,
    dataLocation,
    dataFilename,
    instance,
    backDefault,
    update,
    init
  }
})