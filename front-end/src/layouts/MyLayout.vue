<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated>
      <q-toolbar
        class="bg-white text-red"
      >
        <q-toolbar-title>
          <router-link to="/">
            <q-img
              spinner-color="white"
              style="height: 45px; max-width: 62px"
              src="statics/logo.jpg"
            />
          </router-link>
        </q-toolbar-title>
        <q-chip
          color="white"
          text-color="blue"
          style="font: bold 150% Cursive;"
        >
          TTM Administration Board
        </q-chip>
        <q-space />
      </q-toolbar>
    </q-header>

    <q-drawer
      v-model="leftDrawerOpen"
      :width="250"
      bordered
      content-class="bg-blue-1"
      @mouseover="miniState = false"
      @mouseout="miniState = true"
    >
      <q-list>
        <q-item-label header>
          Management
        </q-item-label>
        <q-item
          v-for="nav in navs"
          :key="nav.label"
          clickable
          exact
          :to="nav.to"
        >
          <q-item-section avatar>
            <q-icon :name="nav.icon" />
          </q-item-section>
          <q-item-section>
            <q-item-label>
              {{ nav.label }}
            </q-item-label>
            <q-item-label caption>
              {{ nav.description }}
            </q-item-label>
          </q-item-section>
        </q-item>
      </q-list>
    </q-drawer>
    <q-page-container>
      <router-view />
    </q-page-container>
    <q-footer
      elevated
      class="bg-blue text-white"
    >
      <q-toolbar>
        <q-toolbar-title
          style="font: 120% Cursive;"
        >
          Copyright @DeCom Technology Ltd. 2020
        </q-toolbar-title>
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from 'vuex'

export default {
  name: 'MyLayout',
  components: {
  },
  data () {
    return {
      leftDrawerOpen: false // this.$q.platform.is.desktop,
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),
    role: function () {
      return this.currentUser.Role
    },
    navs: function () {
      if (this.role === 'admin') {
        return [
          {
            label: 'Staffs',
            icon: 'description',
            to: '/StaffAdmin'
          },
          {
            label: 'Operators',
            icon: 'description',
            to: '/OperatorAdmin'
          },
          {
            label: 'Mentors',
            icon: 'receipt',
            to: '/MentorAdmin'
          },
          {
            label: 'Students',
            icon: 'receipt',
            to: '/StudentAdmin'
          },
          {
            label: 'Quizzes',
            icon: 'receipt',
            to: '/QuizAdmin'
          }
        ]
      } else if (this.role === 'operator') {
        return [
          {
            label: 'Quizzes',
            icon: 'receipt',
            to: '/QuizAdmin'
          }
        ]
      } else { // staff
        return [
          {
            label: 'Operators',
            icon: 'description',
            to: '/UserAdmin'
          },
          {
            label: 'Mentors',
            icon: 'receipt',
            to: '/MentorAdmin'
          },
          {
            label: 'Students',
            icon: 'receipt',
            to: '/StudentAdmin'
          },
          {
            label: 'Quizzes',
            icon: 'receipt',
            to: '/QuizAdmin'
          }
        ]
      }
    }
  },
  watch: {
    '$route' (to, from) {
      console.log(this.$route.path)
      console.log(to.path)
      this.manageLayout()
    }
  },
  beforeMount () {
    console.log('layout before mounte - ' + this.$route.path + 'Role: ' + this.role)
    // load local data into store
    this.loadUserList()
  },
  mounted () {
    console.log('layout mounted - ' + this.$route.path)
    this.manageLayout()
    this.loadDesc()
  },
  methods: {
    ...mapMutations('currentUser', ['updateUser']),
    ...mapActions('localMentors', ['updateLocalMentors']),
    manageLayout () {
      if (this.role && this.role !== 'na') {
        this.leftDrawerOpen = true
      } else {
        this.leftDrawerOpen = false
      }
    },
    toSigninup () {
      console.log('set signin to true')
      this.signinDialog = true
    },
    closeSigninDialog (signinOk) {
      console.log('close signin at - ' + signinOk)
      if (signinOk === true) {
        this.toMentorHome()
      }
      this.signinDialog = false
    },
    signout () {
      this.updateUser({
        Gid: '',
        Name: '',
        Username: '',
        LastLoginTime: '',
        Role: 'na'
      })
      if (this.$route.path !== '/') {
        this.$router.push('/')
      } else {
        this.$router.go(this.$router.currentRoute)
      }
    },
    loadUserList () {
      // localStorage.clear()
      let tmpUser = localStorage.getItem('mentors')
      console.log('localstorage Mentors - ' + tmpUser)
      if (tmpUser) {
        let localUser = JSON.parse(tmpUser)
        for (let i = 0; i < localUser.length; i++) {
          if (localUser[i].email === '' || localUser[i].token === '') {
            localUser.splice(i, 1)
          }
        }
        this.updateLocalMentors(localUser)
        console.log('All Mentors - ' + localUser)
      } else {
        this.updateLocalMentors([])
      }

      tmpUser = localStorage.getItem('students')
      console.log('localstorage students - ' + tmpUser)
      if (tmpUser) {
        let localUser = JSON.parse(tmpUser)
        for (let i = 0; i < localUser.length; i++) {
          if (localUser[i].name === '' || localUser[i].password === '') {
            localUser.splice(i, 1)
          }
        }
        this.updateLocalStudents(localUser)
        console.log('All students - ' + localUser)
      } else {
        this.updateLocalStudents([])
      }

      let usedSpace = unescape(encodeURIComponent(JSON.stringify(localStorage))).length
      // let remSpace = window.localStorage.remainingSpace // For IE
      console.log('LocalStorage Used Bytes: ', usedSpace)
    },
    loadDesc () {
      let tmpDesc = localStorage.getItem('kpDescriptions')
      if (tmpDesc) {
        this.copyDesc(JSON.parse(tmpDesc))
      }
    }
  }
}
</script>

<style lang="stylus" scoped>
</style>
