<template>
  <div class="q-px-sm">
    <add-user
      v-if="addUserActive"
      :role="role"
      @addUserDone="fetchUser"
    />
    <q-table
      title="Users Table"
      :data="tableData"
      :columns="columns"
      :visible-columns="visibleColumns"
      separator="cell"
      wrap-cells
      row-key="id"
      :table-header-style="{ backgroundColor: 'lightblue' }"
      fixed-center
      bordered
      dense
      :loading="loading"
    >
      <template v-slot:top="props">
        <div class="col-2 q-table__title">
          Avaialble Users
        </div>
        <q-space />
        <q-btn
          class="q-ml-md"
          flat
          round
          dense
          color="blue"
          :icon="props.inFullscreen ? 'fullscreen_exit' : 'fullscreen'"
          @click="props.toggleFullscreen"
        />
      </template>
      <q-td
        slot="body-cell-delete"
        slot-scope="props"
        :props="props"
      >
        <q-btn
          flat
          icon="delete"
          @click="toDeleteUser(props.row.Gid)"
        />
      </q-td>
      <template v-slot:no-data="{ icon, message, filter }">
        <div class="full-width row flex-center text-accent q-gutter-sm">
          <span>
            You don't have any student, add student by press
          </span>
          <q-icon
            size="2em"
            name="add"
          />
        </div>
      </template>
    </q-table>
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapMutations } from 'vuex'
import { GET_USERS_QUERY } from '../../graphql/queries'
import { DEL_USER_MUTATION } from '../../graphql/mutations'

export default {
  name: 'UserTab',
  components: {
    'add-user': require('components/admin/AddUser.vue').default
  },
  props: {
    role: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      loading: true,
      quizId: '',
      inRole: 'operator',
      roleOptions: ['operator', 'staff'],
      columns: [
        {
          name: 'Gid',
          required: true,
          align: 'center',
          label: 'User GId',
          field: 'Gid',
          sortable: true
        },
        {
          name: 'Username',
          required: true,
          align: 'center',
          label: 'Username',
          field: 'Username',
          sortable: true
        },
        {
          name: 'Password',
          align: 'center',
          label: 'Password',
          field: 'Password',
          sortable: false
        },
        {
          name: 'Name',
          align: 'center',
          label: 'Name',
          field: 'Name',
          sortable: true
        },
        {
          name: 'Role',
          align: 'center',
          label: 'Role',
          field: 'Role',
          sortable: false
        },
        {
          name: 'Email',
          align: 'center',
          label: 'Email',
          field: 'Email',
          sortable: false
        },
        {
          name: 'Mobile',
          align: 'center',
          label: 'Mobile',
          field: 'Mobile',
          sortable: true
        },
        {
          name: 'delete',
          align: 'center',
          label: 'Delete',
          field: 'delete',
          sortable: false
        }
      ]
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),
    ...mapGetters('users', ['getUserList']),
    tableData: function () {
      console.log('user list = ', this.getUserList.length)
      let tmp = []
      for (let i = 0; i < this.getUserList.length; i++) {
        if (this.getUserList[i].Role === this.role) {
          tmp.push(this.getUserList[i])
        }
      }
      return tmp
    },
    activeTab: function () {
      return this.currentUser.activeTab
    },
    addUserActive: function () {
      if (this.currentUser.Role) {
        if (this.currentUser.Role === 'admin' || this.currentUser.Role === 'staff') {
          return true
        }
      }
      return false
    },
    userRole: function () {
      return this.currentUser.Role
    },
    visibleColumns: function () {
      if (this.addUserActive === true) {
        return ['Username', 'Password', 'Role', 'Name', 'Email', 'Mobile', 'delete']
      } else {
        return ['Username', 'Password', 'Role', 'Name', 'Email', 'Mobile']
      }
    }
  },
  watch: {
    activeTab: function (newVal) {
      console.log('newvalue userTab role - ', newVal)
      if (newVal === '2-1' && this.role === 'staff') {
        this.fetchUser()
      } else if (newVal === '2-2' && this.role === 'operator') {
        this.fetchUser()
      }
    }
  },
  mounted () {
    this.fetchUser()
  },
  methods: {
    ...mapMutations('users', ['addNewUser', 'removeUser']),
    toEditQuiz (opt) {
      this.quizId = opt
    },
    fetchUser () {
      console.log('Get user list', this.role)
      this.$apollo
        .query({
          query: GET_USERS_QUERY,
          variables: {
            role: this.role
          }
        })
        .then(response => {
          if (response.data.GetUsers) {
            this.updateUser(response.data.GetUsers)
            console.log('Read user list complete')
          } else {
            console.log('Get users failed')
          }
        })
        .catch(error => {
          console.log(error)
        })
    },
    updateUser (newList) {
      console.log('in update users length - ', newList)
      if (newList.length > 0) {
        this.loading = false
        for (let i = 0; i < newList.length; i++) {
          this.addNewUser(newList[i])
        }
      } else {
        console.log('NO User Available')
      }
    },
    toDeleteUser (opt) {
      this.$apollo
        .mutate({
          mutation: DEL_USER_MUTATION,
          variables: {
            Gid: opt
          }
        })
        .then(response => {
          if (!response.errors) {
            if (response.data.DelUser) {
              console.log('Del quiz successful')
              this.removeUser({
                Gid: opt
              })
            }
          } else {
            console.log('reponse error', response.errors)
          }
        })
        .catch(error => {
          console.log(error)
        })
    }
  }
}
</script>

<style>
</style>
