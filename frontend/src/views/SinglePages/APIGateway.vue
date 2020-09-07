<template>
  <div>
    <h2 class="text-2xl font-bold">{{ info.Name }} (API Gateway)</h2>
    <p class="font-bold">Description:</p>
    {{ info.Description }}
    <p class="font-bold">ID:</p>
    {{ info.ID }}
  </div>
</template>

<script>
export default {
  props: {
    id: {
      type: String,
    },
  },
  computed: {
    info() {
      return this.$store.getters.getResourceById(this.id);
    },
  },
  created() {
    if (!this.info) {
      this.$store.dispatch("getResources", { service: "lambda" });
    }
  },
};
</script>
