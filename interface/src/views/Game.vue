<template>
<div>
<Home/>
</div>
</template>

<script>
import LoginForm from '@/components/LoginForm.vue'
import test1 from '@/testing/test1.vue'
import Home from './Home.vue'

export default {
  name: 'Game',
  data () {
    return {
      version: ''
    }
  },
  component: {
    LoginForm,
    test1,
    Home
  }
}
</script>
