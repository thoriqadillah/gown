<script setup lang="ts">
import { useDrawerStore } from '../store/drawer';

const store = useDrawerStore()

</script>

<template>
  <v-card>
    <v-layout >
      <v-navigation-drawer v-model="store.drawer" elevation="2" expand-on-hover rail :border="0">
        <v-list density="compact" nav class="tw-flex tw-flex-col tw-mt-20">
          <v-list-item prepend-icon="mdi-home" title="Home" value="home"></v-list-item>
          <v-list-item prepend-icon="mdi-alert-outline" title="Failed Download" value="failedDownload"></v-list-item>
          <v-list-item prepend-icon="mdi-tray-full" title="Queued Download" value="queuedDownload"></v-list-item>
          <v-list-item prepend-icon="mdi-cog-outline" title="Settings" value="settings"></v-list-item>
          <v-list-item prepend-icon="mdi-information-outline" title="About" value="about"></v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-main class="tw-h-screen">
        <slot/>
      </v-main>
    </v-layout>
  </v-card>
</template>