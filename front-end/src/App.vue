<template>
  <div id="app">
    <HelloWorld msg="Hello World" />
    <Register msg="Welcome to Your Registration Page"/>
    <Vote msg="Welcome to Your Voting Page"/>
    <ConfirmRegister msg="Welcome to Your Confirm Voter Registration Page"/>
    <ConfirmVote msg="Welcome to Your Confirm Vote Page"/>
  </div>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'
import Register from './components/Register.vue'
import Vote from './components/Vote.vue'
import ConfirmRegister from './components/ConfirmRegister.vue'
import ConfirmVote from './components/ConfirmVote.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    Register,
    Vote,
    ConfirmRegister,
    ConfirmVote
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
