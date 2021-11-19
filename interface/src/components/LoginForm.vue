<template src="../interfaces/login/login.html"></template>

<script>
import axios from 'axios'

export default {
  name: 'LoginForm',
  data () {
    return {
      nickname: '',
      room_code: ''
    }
  },
  methods: {
    onSubmit: function () {
      console.log(this.nickname, this.room_code)

      axios
        .post('http://localhost:8081/join', {
          nickmane: this.nickname,
          room_code: this.room_code
        })
        .then((res) => {
          console.log(res.data)
          localStorage.setItem('interface', JSON.stringify(res.data.interface))
          localStorage.setItem('address', JSON.stringify(res.data.address))
        })
        .catch((err) => {
          console.error(err.response.data.message)
          window.alert(err.response.data.message)
        })
    }
  }
}
</script>

<style lang="scss">
@import '../interfaces/login/login.scss';
</style>
