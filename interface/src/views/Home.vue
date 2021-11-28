<template>
  <div class="home">
    <img alt="Vue logo" src="../assets/logo.png" />
    <WebsocketTestMessenger msg="Input your data to send" />
  </div>
</template>

<script>
// @ is an alias to /src
import WebsocketTestMessenger from '@/components/WebsocketTestMessenger.vue';

export default {
  name: 'Home',
  components: {
    WebsocketTestMessenger,
  },
};
</script>
