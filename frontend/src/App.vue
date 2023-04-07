<script lang="ts" setup>
import { useTheme } from 'vuetify/lib/framework.mjs';

const theme = useTheme()
function toggleTheme() {
  theme.global.name.value = theme.global.current.value.dark ? 'light' : 'dark'
}

</script>

<template>
  <v-app>
    <v-theme-provider :theme="theme.global.name.value">
      <div class="tw-flex tw-w-full tw-justify-center tw-h-screen tw-items-center">
        <v-btn color="primary" @click="toggleTheme()">Singo</v-btn>
      </div>
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