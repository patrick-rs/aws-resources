import Vue from "vue";
import Vuex from "vuex";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    // constants
    supportedServices: ["lambda", "ec2", "s3", "dynamodb", "apigw"],
    displayName: {
      lambda: "Lambda",
      ec2: "EC2",
      s3: "S3",
      dynamodb: "DynamoDB",
      apigw: "API Gateway",
    },
    baseUrl: "https://ndnvxfv8bl.execute-api.us-east-1.amazonaws.com/dev",
    APIKey: "d2e5183861bb4c72b89acde4a2c3c3af",
    tableSchema: {
      lambda: ["Name", "Region", "Runtime"],
      ec2: ["Region", "InstanceType"],
      s3: ["Name"],
      dynamodb: ["Name", "Region"],
      apigw: ["Name"],
    },

    // variables
    loading: false,
    currentResources: [],
  },
  mutations: {
    SET_CURRENT_RESOURCES(state, resources) {
      state.currentResources = resources.data;
    },
  },
  actions: {
    async getResources({ commit, state }, { service }) {
      const response = await fetch(
        `${state.baseUrl}/resources?query=${service}`,
        {
          headers: { "x-api-key": "d2e5183861bb4c72b89acde4a2c3c3af" },
        }
      );
      const data = await response.json();
      commit("SET_CURRENT_RESOURCES", { data });
      state.loading = false;
    },

    async refreshTable({ dispatch, state }, { service }) {
      state.loading = true;
      const body = { regions: ["us-east-1", "us-east-2"] };
      const response = await fetch(
        `${state.baseUrl}/resources?resource=${service}`,
        {
          headers: {
            "x-api-key": "d2e5183861bb4c72b89acde4a2c3c3af",
            "Content-Type": "application/json",
          },
          method: "POST",
          body: JSON.stringify(body),
        }
      );
      if (response.status == 200) {
        dispatch("getResources", { service });
      } else {
        state.loading = false;
      }
    },
  },
  getters: {
    getResourceById: (state) => {
      return (id) => {
        return state.currentResources.find((resource) => resource.SK === id);
      };
    },
    getDisplayName: (state) => {
      return (name) => {
        return state.displayName[name];
      };
    },
  },
});
