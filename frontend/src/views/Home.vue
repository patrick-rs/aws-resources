<template>
  <div>
    <h1 class="text-2xl">Services</h1>
    <span>
      <div
        class="w-48 h-20 my-4 bg-gray-300 border-2 rounded-md"
        v-for="(service, index) in supportedServices"
        :key="index"
      >
        <ServiceCard :service="service" />
      </div>
    </span>
    <ServiceCard />
  </div>
</template>

<script>
// @ is an alias to /src
import { mapState } from "vuex";
import ServiceCard from "../components/ServiceCard";
export default {
  name: "Home",
  components: { ServiceCard },
  computed: {
    ...mapState(["supportedServices"]),
  },
};
</script>
