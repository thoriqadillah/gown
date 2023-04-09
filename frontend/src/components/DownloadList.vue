<script setup lang="ts">
let date = new Date()
const now = date.toLocaleDateString()

const desserts = [
          {
            status: "Success",
            name: 'Frozen Yogurt',
            timeElapsed: 159,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'Ice cream sandwich',
            timeElapsed: 237,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'Eclair',
            timeElapsed: 262,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'Cupcake',
            timeElapsed: 305,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'Gingerbread',
            timeElapsed: 356,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'Jelly bean',
            timeElapsed: 375,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'Lollipop',
            timeElapsed: 392,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'Honeycomb',
            timeElapsed: 408,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'Donut',
            timeElapsed: 452,
            size: 300,
            date: now, 
          },
          {
            status: "Success",
            name: 'KitKat',
            timeElapsed: 518,
            size: 300,
            date: now, 
          },
        ]
</script>

<template>
  <div class="tw-overflow-x-auto tw-mx-5 xl:tw-mx-0 xl:tw-ml-5">
    <v-table density="compact">
    <thead>
      <tr>
        <th class="text-left tw-cursor-pointer">
          <div class="tw-flex tw-justify-between tw-items-center">
            <span>Name</span>
            <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
          </div>
        </th>
        <th class="text-left tw-cursor-pointer">
          <div class="tw-flex tw-justify-between tw-items-center">
            <span>Time Elapsed</span>
            <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
          </div>
        </th>
        <th class="text-left tw-cursor-pointer">
          <div class="tw-flex tw-justify-between tw-items-center">
            <span>Size</span>
            <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
          </div>
        </th>
        <th class="text-left tw-cursor-pointer">
          <div class="tw-flex tw-justify-between tw-items-center">
            <span>Date</span>
            <v-icon icon="mdi-arrow-up-down" class="tw-text-sm"></v-icon>
          </div>
        </th>
      </tr>
    </thead>
    <tbody>
      <tr
        v-for="item in desserts"
        :key="item.name"
      >
        <td class="tw-rounded-sm">
          <v-icon icon="mdi-check-circle-outline" color="success" class="tw-text-sm tw-mr-2"></v-icon>
          {{ item.name }}
        </td>
        <td class="tw-rounded-sm text-left">{{ item.timeElapsed }}</td>
        <td class="tw-rounded-sm text-left">{{ item.size }}</td>
        <td class="tw-rounded-sm text-left">{{ item.date }}</td>
      </tr>
    </tbody>
  </v-table>
  </div>
</template>

<style>
.v-table .v-table__wrapper > table > tbody > tr > td {
  border-bottom: none !important;
}
.v-table .v-table__wrapper > table > tbody > tr:nth-child(even) {
  background: rgb(87, 83, 78, 0.05);
}
</style>