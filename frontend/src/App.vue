<template>
  <div id="app" class="font-medium min-h-screen bg-gray-100">
    <div class="bg-blue-700 w-full">
      <router-link class="text-4xl py-4 p-20 text-white" :to="{ name: 'home' }"
        >AWS Resources</router-link
      >
    </div>
    <router-view class="px-20 max-h-screen" />
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
@tailwind base;
@tailwind components;
@tailwind utilities;
</style>
