<template>
  <div class="register">
    <h1>Registration Page</h1>
    <form >
        <label for="fname">First Name</label> <br>
        <input v-model="fname" id="fname" placeholder="ex: John"> <br>
        <br>
        <label for="lname">Last Name</label> <br>
        <input v-model="lname" id="lname" placeholder="ex: Smith"> <br>
        <br>
        <label for="ssn">Social Security Number</label> <br>
        <input v-model="ssn" id="ssn" placeholder="ex: 123456"> <br>
        <br>
        <button type="button" value="Submit" v-on:click="registerVoter()" >Submit</button>
    </form>
    
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Register',
  props: {
    msg: String
  },
  data: function() {
      return {
          fname: '',
          lname: '',
          ssn: ''
      };
  },
  methods: {
      registerVoter: function() {
          const data = {
              first_name: this.fname,
              last_name: this.lname,
              ssn: this.ssn
          }
          // console.log(data);
          axios.
            post('http://localhost:8040/register', data)
            .then(function (response) {
                console.log(response);
            })
            .catch(error => {
                this.errorMessage = error.message;
                console.error("There was an error!", error);
            });

        this.fname = '';
        this.lname = '';
        this.ssn = '';
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
