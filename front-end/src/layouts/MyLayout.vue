<template>
  <q-layout view="hHh lpR fFf">
    <q-header elevated>
      <q-toolbar
        class="bg-blue-10 text-white"
      >
        <q-toolbar-title>
          <router-link to="/">
            <q-icon
              size="lg"
              name="supervised_user_circle"
              color="white"
            />
            <q-chip
              size="lg"
              color="blue-10"
              text-color="white"
              style="font: bold 100% Cursive;"
            >
              TTM Administration Board
            </q-chip>
          </router-link>
        </q-toolbar-title>
        <q-space />
        <div
          v-if="loginRole!=''"
          class="q-mx-sm"
        >
          <q-chip
            size="sm"
            color="blue-10"
            text-color="red"
            style="font: bold 150% Verdana;"
          >
            Welcom: {{ loginName }}
          </q-chip>
          <q-btn
            flat
            no-caps
            text-color="white"
            @click="signout()"
          >
            <q-icon name="exit_to_app" />
          </q-btn>
        </div>
      </q-toolbar>
    </q-header>
    <q-drawer
      v-model="leftDrawerOpen"
      :width="180"
      bordered
      :content-style="{ backgroundColor: '#2020C0' }"
    >
      <el-menu
        mode="vertical"
        default-active="1"
        class="el-menu-vertical-demo"
        background-color="#2020C0"
        text-color="#fff"
        active-text-color="#ff804b"
        @open="handleOpen"
        @close="handleClose"
        @select="handleSelect"
      >
        <el-menu-item index="1">
          <i class="el-icon-document" />
          <span slot="title"> Dashboard </span>
        </el-menu-item>
        <el-submenu index="2">
          <template slot="title">
            <i class="el-icon-location" />
            <span> Management </span>
          </template>
          <el-menu-item-group>
            <el-menu-item
              index="2-1"
            >
              Staffs
            </el-menu-item>
            <el-menu-item
              index="2-2"
            >
              Operators
            </el-menu-item>
          </el-menu-item-group>
        </el-submenu>
        <el-submenu index="3">
          <template slot="title">
            <i class="el-icon-location" />
            <span> Quiz </span>
          </template>
          <el-menu-item-group>
            <el-menu-item
              index="3-1"
            >
              Math
            </el-menu-item>
            <el-menu-item
              index="3-2"
            >
              English
            </el-menu-item>
          </el-menu-item-group>
        </el-submenu>
        <el-menu-item index="4">
          <i class="el-icon-setting" />
          <span slot="title"> Quiz Log </span>
        </el-menu-item>
        <el-menu-item index="5">
          <i class="el-icon-edit" />
          <span slot="title"> Task Log </span>
        </el-menu-item>
      </el-menu>
    </q-drawer>
    <q-page-container>
      <router-view />
    </q-page-container>
    <q-footer
      elevated
      class="bg-blue-10 text-white"
    >
      <q-toolbar>
        <q-toolbar-title>
          Copyright @DeCom Technology Ltd. 2020
        </q-toolbar-title>
        <q-btn
          outline
          color="white"
          icon-right="mail"
          label="Contact Us"
          class="q-mr-xl"
          @click="sendMessage"
        />
      </q-toolbar>
    </q-footer>
  </q-layout>
</template>

<script>
import { mapGetters, mapMutations } from 'vuex'

export default {
  name: 'MyLayout',
  components: {
  },
  data () {
    return {
      leftDrawerOpen: false, // this.$q.platform.is.desktop,
      allTabs: [
        { name: '1', title: 'Dashboard' },
        { name: '2-1', title: 'Staffs' },
        { name: '2-2', title: 'Operators' },
        { name: '3', title: 'Quiz' },
        { name: '3-1', title: 'Math' },
        { name: '3-2', title: 'English' },
        { name: '4', title: 'Quiz Log' },
        { name: '5', title: 'Task Log' }
      ]
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),

    loginRole: function () {
      return this.currentUser.Role
    },
    loginName: function () {
      return this.currentUser.Name
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
    console.log('layout before mounte - ' + this.$route.path + 'loginRole: ' + this.loginRole)
    if (this.$route.path !== '/' && this.loginRole === 'na') {
      this.$router.push('/')
    }
  },
  mounted () {
    console.log('layout mounted')
    this.manageLayout()
  },
  methods: {
    ...mapMutations('currentUser', ['updateUser', 'addTab']),

    manageLayout () {
      if (this.loginRole && this.loginRole !== 'na') {
        this.leftDrawerOpen = true
      } else {
        this.leftDrawerOpen = false
      }
    },
    handleOpen (index) {
      console.log('open ...', index)
    },
    handleClose (index) {
      console.log('close ...', index)
    },
    handleSelect (index, indexPath) {
      console.log(index, indexPath)
      // set current menu index into store
      for (let i = 0; i < this.allTabs.length; i++) {
        if (index === this.allTabs[i].name) {
          this.addTab(this.allTabs[i])
          break
        }
      }
    },
    sendMessage () {
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
    }
  }
}
</script>

<style lang="stylus" scoped>
</style>
