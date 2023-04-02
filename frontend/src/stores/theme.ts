import { defineStore } from "pinia";
import { ref } from "vue";
import { setting } from '../../wailsjs/go/models'

export const useTheme = defineStore('theme', () => {
  const theme = ref<setting.Themes>()

  function setTheme(themes: setting.Themes) {
    theme.value = themes
  }

  return {
    theme,
    setTheme
  }
})