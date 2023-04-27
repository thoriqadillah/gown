<script setup lang="ts">
import DownloadDialog from './DownloadDialog.vue';
import { ref, inject, provide } from 'vue';
import { useTheme } from 'vuetify/lib/framework.mjs';

const drawer = inject('drawer')
const search = inject('search')

const theme = useTheme()
const themeIcon = ref(theme.global.current.value.dark ? 'mdi-white-balance-sunny' : 'mdi-weather-night')
const themeTooltip = ref(theme.global.current.value.dark ? 'Light Mode' : 'Dark Mode')
function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'customDark'
  themeIcon.value = theme.global.current.value.dark ? 'mdi-white-balance-sunny' : 'mdi-weather-night'
  themeTooltip.value = theme.global.current.value.dark ? 'Light Mode' : 'Dark Mode'
}
</script>

<template>
  <nav class="tw-mx-5 tw-my-3 tw-flex tw-justify-between tw-gap-8 tw-items-center">
    <p class="text-button">GOWN</p>

    <v-text-field v-model="search" density="compact" variant="outlined" color="primary" label="Search" append-inner-icon="mdi-magnify" single-line hide-details/>

    <div class="tw-flex tw-gap-5">
      <v-tooltip :text="themeTooltip" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" :icon="themeIcon" @click="toggleTheme()" variant="flat"/>
        </template>
      </v-tooltip>

      <v-divider vertical/>

      <v-tooltip text="Resume All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" icon="mdi-play" variant="flat"/>
        </template>
      </v-tooltip>

      <v-tooltip text="Pause All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" icon="mdi-pause" variant="flat"/>
        </template>
      </v-tooltip>

      <v-tooltip text="Stop All" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" icon="mdi-stop" variant="flat"/>
        </template>
      </v-tooltip>

      <v-divider vertical/>
      
      <v-tooltip text="Add into Queue" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props" density="comfortable" icon="mdi-tray-plus" variant="flat"/>
        </template>
      </v-tooltip>

      <v-tooltip text="New" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn v-bind:="props" variant="flat" density="comfortable" icon>
            <v-icon icon="mdi-plus"></v-icon>
            <DownloadDialog/>
          </v-btn>
        </template>
      </v-tooltip>

      <v-tooltip text="Menu" location="bottom">
        <template v-slot:activator="{ props }">
          <v-btn class="xl:tw-hidden" v-bind="props" density="comfortable" icon="mdi-dots-vertical" variant="flat" @click="drawer = !drawer"/>
        </template>
      </v-tooltip>
    </div>
  </nav>
</template>