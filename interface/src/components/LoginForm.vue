<template>
    <div>
        <form @submit.prevent="onSubmit" method="post">
            <div class="form-text">
                <label>Room key</label>
                <input type="text" name="room_code" placeholder="Enter room code" v-model="room_code" />
            </div>
            <div class="form-text">
                <label>Player's nick</label>
                <input type="text" name="nickname" placeholder="Enter your nickname" v-model="nickname" />
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
      const dataForm = new FormData()
      dataForm.append('nickname', this.nickname)
      dataForm.append('room_code', this.room_code)
      // dataForm.getAll('nickname') - use it to get an array of nicknames

      axios.post('http://localhost:8080/join', dataForm)
        .then(res => console.log(res.data))
        .catch(err => {
          console.error(err.message)
        })
    }
  }
}
</script>
