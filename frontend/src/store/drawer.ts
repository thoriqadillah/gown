import { defineStore } from "pinia";
import { ref } from "vue";

export const useDrawerStore = defineStore('drawer', () => {
  const drawer = ref(true)
  const openDrawer = () => drawer.value = true
  const closeDrawer = () => drawer.value = false
  
  return {
    drawer,
    openDrawer,
    closeDrawer
  }
})