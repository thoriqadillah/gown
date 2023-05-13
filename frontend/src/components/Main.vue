<script setup lang="ts">
import { provide, ref } from 'vue';
import { useDisplay } from 'vuetify/lib/framework.mjs';
import SettingDialog from './setting/SettingDialog.vue';

const display = useDisplay()
const drawer = ref(display.width.value !== 815)
provide('drawer', drawer)

const reload = () => location.reload()

const settingDialog = ref(false)

</script>

<template>
  <v-card>
    <v-layout >
      <v-navigation-drawer v-model="drawer" :elevation="0" expand-on-hover rail :border="0" location="right">
        <v-list density="compact" nav class="tw-flex tw-flex-col tw-mt-12">
          <v-list-item active-color="primary" prepend-icon="mdi-brightness-4" title="Theme" value="theme"></v-list-item>
          <v-list-item active-color="primary" prepend-icon="mdi-bug-outline" title="Log" value="log"></v-list-item>
          <v-list-item active-color="primary" prepend-icon="mdi-replay" title="Refresh" value="refresh" @click="reload()"></v-list-item>
          <v-list-item active-color="primary" prepend-icon="mdi-cog-outline" :active="settingDialog" v-model="settingDialog" title="Settings" value="settings">
            <setting-dialog :dialog="settingDialog" @update-dialog="(newval: boolean) => settingDialog = newval"/>
          </v-list-item>
          <v-list-item active-color="primary" prepend-icon="mdi-information-outline" title="About" value="about"></v-list-item>
        </v-list>
      </v-navigation-drawer>

      <v-main class="tw-overflow-scroll tw-h-screen tw-w-fit">
        <slot/>
      </v-main>
    </v-layout>
  </v-card>
</template>