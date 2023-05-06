import { defineStore } from "pinia";
import { ref } from "vue";

type AlertType = 'info' | 'warning' | 'danger' | 'success'

export const useAlert = defineStore('alert', () => {
  const status = ref(false)
  const message = ref('')
  const type = ref<AlertType>('danger')

  const open = (msg: string, alerttype: AlertType = 'danger') => {
    status.value = true
    message.value = msg
    type.value = alerttype
  }
  
  return {
    status,
    message,
    type,
    open,
  }
})