<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import { useTheme } from './stores/theme'
import { Theme } from '../wailsjs/go/main/App'

const store = useTheme()
Theme().then(theme => {
  store.setTheme(theme)
})

</script>

<template>
  <HelloWorld/>
</template>