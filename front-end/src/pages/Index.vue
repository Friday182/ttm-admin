<template>
  <q-page class="flex-center" style="background: lightblue;">
    <q-card
      class="absolute-center"
      style="width: 450px; max-width: 80vw;background: lightblue;"
    >
      <q-card-section align="center">
        <q-btn
          flat
          no-caps
          dense
          size="lg"
          style="color: green;"
          label="Login"
        />
      </q-card-section>
      <q-card-section class="q-px-xl q-mx-xl q-mt-sm">
        <q-input
          ref="nameInput"
          v-model="loginName"
          rounded
          outlined
          autofocus
          color="green"
          label="Username:"
          @keyup.enter.native="addByEnter(0)"
        >
          <template v-slot:append>
            <q-icon
              name="face"
              color="orange"
            />
          </template>
        </q-input>
      </q-card-section>
      <q-card-section class="q-px-xl q-mx-xl q-mb-md">
        <q-input
          ref="pwInput"
          v-model="password"
          rounded
          outlined
          color="green"
          label="Password:"
          :type="isPwd ? 'password' : 'text'"
          @keyup.enter.native="addByEnter(1)"
        >
          <template v-slot:append>
            <q-icon
              color="blue"
              :name="isPwd ? 'visibility_off' : 'visibility'"
              class="cursor-pointer"
              @click="isPwd = !isPwd"
            />
          </template>
        </q-input>
      </q-card-section>
      <q-card-actions
        align="center"
        class="text-teal"
      >
        <q-btn
          size="lg"
          label="OK"
          @click="toLogin"
        />
      </q-card-actions>
      <q-card-actions
        class="text-teal"
        align="center"
      >
        <q-checkbox
          v-model="rememberUser"
          keep-color
          left-label
          color="blue"
          label="Remember me on this computer？"
        />
      </q-card-actions>
    </q-card>
  </q-page>
</template>

<script>
import { LOGIN_MUTATION } from '../graphql/mutations'
import { mapMutations } from 'vuex'

export default {
  components: {
    // 'comment': require('components/common/Comment.vue').default
  },
  data () {
    return {
      loginName: 'Phoenix',
      password: 'jiangbo007',
      isPwd: true,
      windowHeight: '',
      rememberUser: true,
      loginForm: {
        username: 'test',
        password: 'test'
      }
    }
  },
  created () {
    this.windowHeight = window.innerHeight + 'px'
  },
  methods: {
    ...mapMutations('currentUser', ['updateUser']),
    toLogin () {
      this.$apollo
        .mutate({
          mutation: LOGIN_MUTATION,
          variables: {
            username: this.loginName,
            password: this.password
          }
        })
        .then(response => {
          if (!response.errors) {
            if (response.data.UserLogin) {
              console.log('studentLogin for - ' + response.data.UserLogin.Name)
              this.setCurrentUser(response.data.UserLogin)
            }
          } else {
            console.log('reponse error', response.errors)
          }
        })
        .catch(error => {
          console.log(error)
        })
    },
    addByEnter (opt) {
      if (this.loginName !== '' && this.password !== '') {
        this.toLogin()
      } else {
        if (this.loginName !== '') {
          this.$refs.pwInput.focus()
        } else {
          this.$refs.nameInput.focus()
        }
      }
    },
    setCurrentUser (obj) {
      if (obj.Gid !== '') {
        console.log('Write current user to store - ' + obj.Role)
        this.updateUser({
          Gid: obj.Gid,
          Name: obj.Name,
          Username: obj.Username,
          LastLoginTime: '',
          Role: obj.Role,
          Comment: obj.Comment,
          tabs: [{
            name: '1',
            title: 'Dashboard'
          }],
          activeTab: '1',
          wsOk: false
        })
        if (this.$route.path !== '/Navigator') {
          this.$router.push('/Navigator')
        }
      } else {
        console.log('something wrong')
      }
    }
  }
}
</script>

<style>
</style>
