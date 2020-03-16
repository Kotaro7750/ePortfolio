<template>
<div>
  <b-navbar toggleable="lg" type="dark" variant="dark">
    <b-navbar-brand href="/">ePortfolio</b-navbar-brand>

    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

    <b-collapse id="nav-collapse" is-nav>
      <!-- Right aligned nav items -->
      <b-navbar-nav class="ml-auto">
        <b-nav-item-dropdown right>
          <template v-slot:button-content>
            <em>資産</em>
          </template>
          <b-dropdown-item href="/">資産</b-dropdown-item>
        </b-nav-item-dropdown>

        <b-nav-item href="/purchase">株式購入</b-nav-item>
        <b-nav-item href="/ticker">ティッカー</b-nav-item>

        <b-nav-item @click="signOut">サインアウト</b-nav-item>
      </b-navbar-nav>
    </b-collapse>
  </b-navbar>
</div>
</template>

<script>
import firebase from 'firebase/app'
export default {
  name: 'Header',
  methods: {
    signOut: function () {
      firebase.auth().signOut().then(() => {
        this.$router.push('/signin')
      })
    }
  }
}
</script>
