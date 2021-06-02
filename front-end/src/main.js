import Vue from 'vue'
// import App from './App.vue'
import HelloWorld from './components/HelloWorld.vue'
import Register from './components/Register.vue'
import Vote from './components/Vote.vue'
import ConfirmRegister from './components/ConfirmRegister.vue'
import ConfirmVote from './components/ConfirmVote.vue'


Vue.config.productionTip = false

const NotFound = { template: '<p>Page not found</p>' }

const routes = {
  '/': HelloWorld,
  '/register': Register,
  '/vote': Vote,
  '/confirm-register': ConfirmRegister,
  '/confirm-vote': ConfirmVote
}

// new Vue({
//   render: h => h(App),
// }).$mount('#app')

new Vue({
  el: '#app',
  data: {
    currentRoute: window.location.pathname
  },
  computed: {
    ViewComponent () {
      return routes[this.currentRoute] || NotFound
    }
  },
  render (h) { return h(this.ViewComponent) }
})