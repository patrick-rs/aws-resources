<template>
  <div>
    <div class="flex">
      <h2 class="text-2xl font-bold justify-center py-2 pr-8">
        {{ displayName }}
      </h2>
      <button
        @click="getResources"
        class="m-2 px-4 p-2 rounded-lg text-gray-100"
        :class="{ 'bg-green-600': !loading, 'bg-gray-600': loading }"
      >
        Refresh
      </button>
    </div>
    <table>
      <thead>
        <tr>
          <th
            class="px-4 py-1 border-2 bold text-white border-gray-800 bg-blue-500"
            v-for="(column, index) in schema"
            :key="index"
          >
            {{ column }}
          </th>
          <th
            class="px-4 py-1 border-2 bold text-white border-gray-800 bg-blue-500"
          >
            Details
          </th>
        </tr>
      </thead>
      <tbody>
        <ServiceRow
          v-for="(resource, index) in currentResources"
          :key="index"
          :resource="resource"
          :schema="schema"
        />
      </tbody>
    </table>
  </div>
</template>

<script>
import ServiceRow from "../components/ServiceRow";
import { mapState } from "vuex";
export default {
  components: {
    ServiceRow,
  },
  props: {
    serviceName: {
      type: String,
    },
  },
  computed: {
    ...mapState(["loading", "currentResources"]),
    schema() {
      return this.$store.state.tableSchema[this.serviceName];
    },
    displayName() {
      return this.$store.getters.getDisplayName(this.serviceName);
    },
  },
  created() {
    this.$store.dispatch("getResources", { service: this.serviceName });
  },
  methods: {
    getResources() {
      this.$store.dispatch("refreshTable", { service: this.serviceName });
    },
  },
};
</script>
