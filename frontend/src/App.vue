<script lang="ts" setup>
import { useTheme } from 'vuetify/lib/framework.mjs';
import Navigation from './components/Navigation.vue'
import DownloadList from './components/DownloadList.vue';
import Main from './components/Main.vue';
import { computed, provide, ref } from 'vue';
import { useDownloads } from './store/downloads'
import { InitData, InitSetting } from '../wailsjs/go/main/App';
import { download, setting } from '../wailsjs/go/models';
import { useSettings } from './store/setting';

const theme = useTheme()
const downloads = useDownloads()
const settings = useSettings()

const search = ref('')
provide('search', search)

const data = computed(() => downloads.filter(search.value))

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