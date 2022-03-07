<template>

<div id="solid-bg"></div>
<div id="image-bg"></div>

<div id="form-container">
<form @submit.prevent="onSubmit" method="post">
  <div id="logo">
    <div id="logo-container">
      <i id="logo-rook" class="fa-solid fa-chess-rook"></i>
      <i id="logo-gamepad" class="fas fa-gamepad"></i>
    </div>
    joystick
  </div>

  <div class="form-element">
    <label>ROOM CODE</label>
      <input
        type="text"
        name="room_code"
        placeholder="CODE"
        v-model="room_code"
        required
        maxlength="10"
        autocomplete="off"
      />
  </div>

  <div class="form-element">
    <label>NICKNAME</label>
    <input
      type="text"
      name="nickname"
      placeholder="NAME"
      v-model="nickname"
      required
      maxlength="10"
    />
  </div>

  <button id="connect-btn" type="submit" name="button">
    <div id="connect-btn-1">CONNECT</div>
    <div id="connect-btn-2"></div>
  </button>
</form>
</div>

</template>

<script>
import axios from 'axios'
import { dispatchSaveInterface } from '../store/actionCreator'

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
        .post('http://192.168.1.104:8081/join', {
          nickname: this.nickname,
          room_code: this.room_code
        })
        .then((res) => {
          localStorage.setItem('gui', JSON.stringify(res.data.gui))
          dispatchSaveInterface(res.data.gui)
          console.log(res.data.gui)
          localStorage.setItem('global_id', JSON.stringify(res.data.global_id))
          this.$router.push('/game')
        })
        .catch((err) => {
          console.error(err.response.data.message)
          window.alert(err.response.data.message)
        })
    }
  }
}
</script>
