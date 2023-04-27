<script lang="ts" setup>
import { InitData, InitSetting } from '../wailsjs/go/main/App';
import { useDownloads } from './store/downloads'
import { useSettings } from './store/setting';
import { useTheme } from 'vuetify/lib/framework.mjs';
import DownloadList from './components/DownloadList.vue';
import Navigation from './components/Navigation.vue'
import Main from './components/Main.vue';

const theme = useTheme()
const downloads = useDownloads()
const settings = useSettings()

InitData().then(data => downloads.setData(data))
InitSetting().then(setting => settings.init(setting))
</script>

<template>
  <v-app>
    <v-theme-provider :theme="theme.global.name.value">
      <Main>
        <Navigation />
        <DownloadList />
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