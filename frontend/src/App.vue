<script lang="ts" setup>
import { useTheme } from 'vuetify/lib/framework.mjs';
import Navigation from './components/Navigation.vue'
import DownloadList from './components/DownloadList.vue';
import Main from './components/Main.vue';
import { useDateFormat } from '@vueuse/shared';
import { useNow } from '@vueuse/core';
import { Download } from './types/download';
import { computed } from 'vue';
import { useDownloads } from './store/downloads'

const theme = useTheme()

const now = useDateFormat(useNow(), 'MMMM DD, YYYY HH:mm').value
const desserts: Download[] = [
  {
    status: {
      name: 'queued',
      icon: "mdi-tray-full",
      color: 'info',
    },
    name: 'Frozen Yogurt',
    timeElapsed: 159,
    size: 300,
    date: now,
    type: {
      name: 'music',
      icon: 'mdi-music-box' ,
      color: 'purple-accent-2'
    }
  },
  {
    status: {
      name: 'failed',
      icon: "mdi-alert-outline",
      color: 'warning',
    },
    name: 'Ice cream sandwich',
    timeElapsed: 237,
    size: 300,
    date: now,
    type: {
      name: 'document',
      icon: 'mdi-file-document',
      color: 'blue-accent-2'
    }
  },
  {
    status: {
      name: 'success',
      icon: "mdi-check-circle-outline",
      color: 'success',
    },
    name: 'Eclair',
    timeElapsed: 262,
    size: 300,
    date: now,
    type: {
      name: 'document',
      icon: 'mdi-file-document',
      color: 'blue-accent-2'
    }
  },
  {
    status: {
      name: 'success',
      icon: "mdi-check-circle-outline",
      color: 'success',
    },
    name: 'Cupcake',
    timeElapsed: 305,
    size: 300,
    date: now,
    type: {
      name: 'video',
      icon: 'mdi-video',
      color: 'deep-orange-accent-2'
    }
  },
  {
    status: {
      name: 'failed',
      icon: "mdi-alert-outline",
      color: 'warning',
    },
    name: 'Gingerbread',
    timeElapsed: 356,
    size: 300,
    date: now,
    type: {
      name: 'image',
      icon: 'mdi-image',
      color: 'red-accent-2'
    }
  },
  {
    status: {
      name: 'success',
      icon: "mdi-check-circle-outline",
      color: 'success',
    },
    name: 'Jelly bean',
    timeElapsed: 375,
    size: 300,
    date: now,
    type: {
      name: 'compressed',
      icon: 'mdi-zip-box',
      color: 'yellow-accent-4'
    }
  },
  {
    status: {
      name: 'paused',
      icon: "mdi-pause-circle-outline",
      color: '',
    },
    name: 'Lollipop',
    timeElapsed: 392,
    size: 300,
    date: now,
    type: {
      name: 'compressed',
      icon: 'mdi-zip-box',
      color: 'yellow-accent-4'
    }
  },
  {
    status: {
      name: 'paused',
      icon: "mdi-pause-circle-outline",
      color: '',
    },
    name: 'Honeycomb',
    timeElapsed: 408,
    size: 300,
    date: now,
    type: {
      name: 'image',
      icon: 'mdi-image',
      color: 'red-accent-2'
    }
  },
  {
    status: {
      name: 'paused',
      icon: "mdi-pause-circle-outline",
      color: '',
    },
    name: 'Donut',
    timeElapsed: 452,
    size: 300,
    date: now,
    type: {
      name: 'music',
      icon: 'mdi-music-box',
      color: 'purple-accent-2'
    }
  },
  {
    status: {
      name: 'success',
      icon: "mdi-check-circle-outline",
      color: 'success',
    },
    name: 'Lorem',
    timeElapsed: 518,
    size: 300,
    date: now,
    type: {
      name: 'music',
      icon: 'mdi-music-box',
      color: 'purple-accent-2'
    }
  },
]

const downloads = useDownloads()
downloads.setData(desserts)
const data = computed(() => downloads.filter(downloads.search))

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