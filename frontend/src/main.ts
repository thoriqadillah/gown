import App from './App.vue'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify } from 'vuetify'
import { useDark } from '@vueuse/core'

import 'vuetify/styles'
import './index.css'

const pinia = createPinia()
const vuetify = createVuetify({
  theme: {
    defaultTheme: useDark() ? 'dark' : 'light'
  }
})

createApp(App)
  .use(pinia)
  .use(vuetify)
  .mount('#app')

console.log()
