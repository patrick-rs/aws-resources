<template>
  <tr>
    <td
      v-for="(column, index) in schema"
      :key="index"
      class="border-2 border-gray-800 rounded-lg p-1"
    >
      {{ resource[column] }}
    </td>
    <td class="border-2 border-gray-800 rounded-lg p-1">
      <router-link
        class="text-blue-400 underline p-4"
        :to="{ name: resource.PK, params: { id: resource.SK } }"
      >
        View
      </router-link>
    </td>
  </tr>
</template>

<script>
export default {
  props: {
    resource: {
      type: Object,
    },
    schema: {
      type: Array,
    },
  },
};
</script>
