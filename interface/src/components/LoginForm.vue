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
            room_code: '',
            errorMessage: '' 
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
                    // get error message from body or default to response status
                    const error = (data && data.message) || response.status;
                    console.log(error);
                    //return Promise.reject(error);
                }
                response.json()
            })
            .then(data => console.log(json))
            .catch(error => {
                console.error(error);
            })
        },
        /*displayError(error) {
            this.errorMessage = error.toString();
            console.error('There was an error!', error);
        }*/
    },
}

</script>