import { defineStore } from "pinia";
import { ref } from "vue";
import { setting } from "../../wailsjs/go/models";

export const useSettings = defineStore('settings', () => {
  const value = ref(new setting.Settings())

  const init = (s: setting.Settings) => {
    value.value = s
  }
  
  return {
    value,
    init
  }
})