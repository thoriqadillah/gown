import { defineStore } from "pinia";
import { ref } from "vue";
import { setting } from "../../wailsjs/go/models";

export const useSettings = defineStore('settings', () => {
  const partsize = ref(0)
  const concurrency = ref(0)
  const maxtries = ref(0)
  const simmultanousNum = ref(0)
  const saveLocation = ref('')
  const dataLocation = ref('')
  const dataFilename = ref('')

  const defaults = ref(new setting.Settings())

  const init = (s: setting.Settings) => {
    partsize.value = s.partsize
    concurrency.value = s.concurrency
    maxtries.value = s.maxtries
    simmultanousNum.value = s.simmultanousNum
    saveLocation.value = s.saveLocation
    dataLocation.value = s.dataLocation
    dataFilename.value = s.dataFilename

    defaults.value = s
  }
  
  return {
    partsize,
    concurrency,
    maxtries,
    simmultanousNum,
    saveLocation,
    dataLocation,
    dataFilename,
    init
  }
})