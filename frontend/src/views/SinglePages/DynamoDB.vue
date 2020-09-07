<template>
  <div>
    <h2 class="text-2xl font-bold">{{ info.Name }} (DynamoDB Table)</h2>
    <div>
      <p class="font-bold">Arn</p>
      {{ info.Arn }}
      <p class="font-bold">Region</p>
      {{ info.Region }}
      <p class="font-bold">RCU</p>
      {{ info.RCU }}
      <p class="font-bold">WCU</p>
      {{ info.WCU }}
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
      this.$store.dispatch("getResources", { service: "dynamodb" });
    }
  },
};
</script>
