<template>
  <div>
    <router-link
      :to="{ name: 'service-table', params: { serviceName: service } }"
    >
      <h2 class="text-xl text-center my-4">{{ displayName }}</h2>
    </router-link>
  </div>
</template>

<script>
export default {
  props: {
    service: {
      type: String,
    },
  },
  computed: {
    displayName() {
      return this.$store.getters.getDisplayName(this.service);
    },
  },
};
</script>
