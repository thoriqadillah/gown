import App from './App.vue'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify } from 'vuetify'
import { aliases, mdi } from 'vuetify/iconsets/mdi'
import { useDark } from '@vueuse/core'

import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import './index.css'

const pinia = createPinia()
const vuetify = createVuetify({
  theme: {
    defaultTheme: useDark() ? 'dark' : 'light'
  },
  icons: {
    defaultSet: 'mdi',
    aliases,
    sets: {
      mdi
    }
  }
})

createApp(App)
  .use(pinia)
  .use(vuetify)
  .mount('#app')

console.log()
