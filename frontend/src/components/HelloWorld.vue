<script lang="ts" setup>
import { ref } from 'vue'
import { useTheme } from '../stores/theme';

const store = useTheme()
const textColor = store.theme?.textColor
const backgroundColor = store.theme?.backgroundColor
</script>

<template>
  text color is {{ textColor }}
  background color is {{ backgroundColor }}
</template>
