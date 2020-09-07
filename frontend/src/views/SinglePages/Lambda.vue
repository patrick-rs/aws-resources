<template>
  <div>
    <h2 class="text-2xl font-bold">{{ info.Name }} (Lambda Function)</h2>
    <div>
      <p class="font-bold">Description:</p>
      {{ info.Description }}
      <p class="font-bold">ARN:</p>
      {{ info.Arn }}
      <p class="font-bold">Region:</p>
      {{ info.Region }}
      <p class="font-bold">Runtime:</p>
      {{ info.Runtime }}
    </div>
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
