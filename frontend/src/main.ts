import App from './App.vue'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { createVuetify, ThemeDefinition } from 'vuetify'
import { aliases, mdi } from 'vuetify/iconsets/mdi'
import { useDark } from '@vueuse/core'

import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'
import './index.css'

const customDark: ThemeDefinition = {
  dark: true,
  colors: {
    background: '#121212',
    surface: '#121212'
  }
}

const pinia = createPinia()
const vuetify = createVuetify({
  theme: {
    defaultTheme: useDark() ? 'customDark' : 'light',
    themes: {
      customDark
    }
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