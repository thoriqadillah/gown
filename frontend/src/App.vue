<script lang="ts" setup>
import { useTheme } from 'vuetify/lib/framework.mjs';
import Navigation from './components/Navigation.vue'
import DownloadList from './components/DownloadList.vue';
import Main from './components/Main.vue';
import { useDateFormat } from '@vueuse/shared';
import { computed } from 'vue';
import { useDownloads } from './store/downloads'
import { InitData, InitSetting } from '../wailsjs/go/main/App';
import { download, setting } from '../wailsjs/go/models';
import { useSettings } from './store/setting';

// const now = 
// const desserts: Download[] = [
//   {
//     id: '1',
//     status: {
//       name: 'queued',
//       icon: "mdi-tray-full",
//       color: 'info',
//     },
//     name: 'Frozen Yogurt',
//     timeElapsed: 159,
//     size: 301,
//     date: now,
//     type: {
//       name: 'audio',
//       icon: 'mdi-music-box' ,
//       color: 'purple-accent-2'
//     }
//   },
//   {
//     id: '2',
//     status: {
//       name: 'failed',
//       icon: "mdi-alert-outline",
//       color: 'warning',
//     },
//     name: 'Ice cream sandwich',
//     timeElapsed: 237,
//     size: 310,
//     date: now,
//     type: {
//       name: 'document',
//       icon: 'mdi-file-document',
//       color: 'blue-accent-2'
//     }
//   },
//   {
//     id: '3',
//     status: {
//       name: 'success',
//       icon: "mdi-check-circle-outline",
//       color: 'success',
//     },
//     name: 'Eclair',
//     timeElapsed: 262,
//     size: 50,
//     date: now,
//     type: {
//       name: 'document',
//       icon: 'mdi-file-document',
//       color: 'blue-accent-2'
//     }
//   },
//   {
//     id: '4',
//     status: {
//       name: 'success',
//       icon: "mdi-check-circle-outline",
//       color: 'success',
//     },
//     name: 'Cupcake',
//     timeElapsed: 305,
//     size: 110,
//     date: now,
//     type: {
//       name: 'video',
//       icon: 'mdi-video',
//       color: 'deep-orange-accent-2'
//     }
//   },
//   {
//     id: '5',
//     status: {
//       name: 'failed',
//       icon: "mdi-alert-outline",
//       color: 'warning',
//     },
//     name: 'Gingerbread',
//     timeElapsed: 356,
//     size: 110,
//     date: now,
//     type: {
//       name: 'image',
//       icon: 'mdi-image',
//       color: 'red-accent-2'
//     }
//   },
//   {
//     id: '6',
//     status: {
//       name: 'success',
//       icon: "mdi-check-circle-outline",
//       color: 'success',
//     },
//     name: 'Jelly bean',
//     timeElapsed: 375,
//     size: 220,
//     date: now,
//     type: {
//       name: 'compressed',
//       icon: 'mdi-zip-box',
//       color: 'yellow-accent-4'
//     }
//   },
//   {
//     id: '7',
//     status: {
//       name: 'paused',
//       icon: "mdi-pause-circle-outline",
//       color: '',
//     },
//     name: 'Lollipop',
//     timeElapsed: 392,
//     size: 500,
//     date: now,
//     type: {
//       name: 'compressed',
//       icon: 'mdi-zip-box',
//       color: 'yellow-accent-4'
//     }
//   },
//   {
//     id: '8',
//     status: {
//       name: 'paused',
//       icon: "mdi-pause-circle-outline",
//       color: '',
//     },
//     name: 'Honeycomb',
//     timeElapsed: 408,
//     size: 323,
//     date: now,
//     type: {
//       name: 'image',
//       icon: 'mdi-image',
//       color: 'red-accent-2'
//     }
//   },
//   {
//     id: '9',
//     status: {
//       name: 'paused',
//       icon: "mdi-pause-circle-outline",
//       color: '',
//     },
//     name: 'Donut',
//     timeElapsed: 452,
//     size: 520,
//     date: now,
//     type: {
//       name: 'audio',
//       icon: 'mdi-music-box',
//       color: 'purple-accent-2'
//     }
//   },
//   {
//     id: '10',
//     status: {
//       name: 'success',
//       icon: "mdi-check-circle-outline",
//       color: 'success',
//     },
//     name: 'Lorem',
//     timeElapsed: 518,
//     size: 300,
//     date: now,
//     type: {
//       name: 'audio',
//       icon: 'mdi-music-box',
//       color: 'purple-accent-2'
//     }
//   },
// ]
const theme = useTheme()
const downloads = useDownloads()
const settings = useSettings()

const data = computed(() => downloads.filter(downloads.search))

InitData().then((data: download.Download[]) => {
  downloads.setData(data)
})

InitSetting().then((setting: setting.Settings) => {
  settings.init(setting)
})
</script>

<template>
  <v-app>
    <v-theme-provider :theme="theme.global.name.value">
      <Main>
        <Navigation />
        <DownloadList :list="data"/>
      </Main>
    </v-theme-provider>
  </v-app>
</template>

<style>
body {
  font-family: "Roboto", -apple-system, BlinkMacSystemFont, "Segoe UI",
  "Oxygen", "Ubuntu", "Cantarell", "Fira Sans", "Droid Sans", "Helvetica Neue",
  sans-serif;
}

@font-face {
  font-family: "Roboto";
  font-style: normal;
  font-weight: 400;
  src: local(""),
  url("assets/fonts/roboto-regular.woff2") format("woff2");
}
</style>