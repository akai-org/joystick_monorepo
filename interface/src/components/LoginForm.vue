<template>
<div>
<div id="background"></div>

<div id="logo">
  <i id="logo-gamepad" class="fas fa-gamepad"></i>
  joystick
</div>

<div class="spacer-4"></div>

<div id="form-container">
  <form @submit.prevent="onSubmit" method="post">
    <div class="form-background">
      <div class="spacer-05"></div>

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

      <div class="spacer-1"></div>
    </div>

    <div class="spacer-1"></div>

    <button type="submit" name="button">CONNECT</button>
  </form>

  <div class="spacer-2"></div>

</div>
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
        .post('http://localhost:8081/join', {
          nickname: this.nickname,
          room_code: this.room_code
        })
        .then((res) => {
          localStorage.setItem('gui', JSON.stringify(res.data.gui))
          dispatchSaveInterface(res.data.gui)
          localStorage.setItem('global_id', JSON.stringify(res.data.global_id))
          window.location.href = '/game'
        })
        .catch((err) => {
          console.error(err.response.data.message)
          window.alert(err.response.data.message)
        })
    }
  }
}
</script>

<style lang="scss" src="../app.scss"></style>
