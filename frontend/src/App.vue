<template>
  <div id="app">
    <Contacts />
  </div>
</template>

<script>
import Vue from "vue";
import ElementUI from "element-ui";
Vue.use(ElementUI);
import HelloWorld from "./components/HelloWorld.vue";
import Contacts from "./components/Contacts.vue";
export default {
  name: "app",
  components: {
    HelloWorld,
    Contacts
  }
};
</script>

<style>
#app {
  font-family: "Avenir", Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
