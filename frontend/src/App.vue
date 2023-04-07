<script lang="ts" setup>
import { useTheme } from './stores/theme'
import { Theme } from '../wailsjs/go/main/App'

const store = useTheme()
Theme().then(theme => {
  store.setTheme(theme)
})

</script>

<template>
  <div class="flex w-full justify-center h-screen items-center">
    <p>Center</p>
  </div>
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