<template>
  <div>
    <form @submit.prevent="onSubmit" method="post">
      <div class="form-text">
        <label>Room key</label>
        <input
          type="text"
          name="room_code"
          placeholder="Enter room code"
          v-model="room_code"
        />
      </div>
      <div class="form-text">
        <label>Player's nick</label>
        <input
          type="text"
          name="nickname"
          placeholder="Enter your nickname"
          v-model="nickname"
        />
      </div>
      <button type="submit" name="button">Enter the game</button>
    </form>
  </div>
</template>

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
          localStorage.setItem('response', JSON.stringify(res.data))
        })
        .catch((err) => {
          console.error(err.response.data.message)
          window.alert(err.response.data.message)
        })
    }
  }
}
</script>
