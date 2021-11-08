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
            <button type="submit">Enter the game</button>
        </form>
        <span class="error">

        </span>
    </div>
</template>

<script>
export default {
    name: 'LoginForm',
    data() {
        return {
            nickname: '',
            room_code: ''
        }
    },
    methods: {
        onSubmit() {
            //console.log({this.nickname, this.room_code})
            fetch('http://servergo/join', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    nickname: this.nickname,
                    room_code: this.room_code
                })
            })
            .then(response => {
                // check for error in response
                if(!response.ok) {
                    console.log('There was an error!');
                }
                response.json()
            })
            .then(data => console.log(data))
            .catch(error => {
                console.error(error);
            })
        },
    },
}

</script>