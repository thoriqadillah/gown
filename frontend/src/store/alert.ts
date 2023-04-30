import { defineStore } from "pinia";
import { ref } from "vue";

export const useAlert = defineStore('alert', () => {
  const status = ref(false)
  const message = ref('')

  const open = (msg: string) => {
    status.value = true
    message.value = msg
  }
  
  return {
    status,
    open,
    message
  }
})