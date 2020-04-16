<template>
  <div>
    <div class="row q-mb-md">
      <div class="col-1 q-mr-sm">
        <el-input
          v-model="inName"
          class="myClass"
          placeholder="Name"
          clearable
        />
      </div>
      <div class="col-2 q-mr-sm">
        <el-input
          v-model="inEmail"
          class="myClass"
          placeholder="Email"
          clearable
        />
      </div>
      <div class="col-2 q-mr-sm">
        <el-input
          v-model="inMobile"
          class="myClass"
          placeholder="Mobile"
          clearable
        />
      </div>
      <div class="col-3 q-mr-sm">
        <el-input
          v-model="inComment"
          class="myClass"
          placeholder="Any Comment?"
        />
      </div>
      <q-btn
        class="q-ml-lg"
        label="Add User"
        no-caps
        color="primary"
        style="font: 100% Arial bold"
        @click="addUser()"
      />
    </div>
  </div>
</template>

<script>
import { ADD_USER_MUTATION } from '../../graphql/mutations'
import { mapMutations, mapGetters } from 'vuex'

export default {
  name: 'AddUser',
  props: {
    role: {
      type: String,
      default: 'operator'
    }
  },
  data () {
    return {
      inName: '',
      inRole: '',
      inEmail: '',
      inMobile: '',
      inComment: '',
      roleOptions: [
        { value: '1', label: 'operator' },
        { value: '2', label: 'staff' }
      ],
      period: ''
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),
    loginRole: function () {
      return this.currentUser.Role
    },
    creator: function () {
      return this.currentUser.Username
    }
  },
  methods: {
    ...mapMutations('users', ['addNewUser']),
    addUser () {
      if (this.inName === '' || this.inEmail === '' || this.inMobile === '') {
        return
      }
      this.$apollo
        .mutate({
          mutation: ADD_USER_MUTATION,
          variables: {
            name: this.inName,
            role: this.role,
            email: this.inEmail,
            mobile: this.inMobile,
            comment: this.inComment
          }
        })
        .then(response => {
          if (!response.errors) {
            if (response.data.AddQuiz) {
              this.$emit('addUserDone')
            }
          } else {
            console.log('reponse error', response.errors)
          }
        })
        .catch(error => {
          console.log(error)
        })

      this.clearForm()
    },
    clearForm () {
      this.inName = ''
      this.inRole = ''
      this.inEmail = ''
      this.inMobile = ''
      this.inComment = ''
    }
  }
}
</script>

<style>
.myClass input.el-input__inner {
    border-radius:15px!important;
}
.dialog-input-text >>> .el-input__inner {
  -web-kit-appearance: none;
  -moz-appearance: none;
  font-size: 1.4em;
  height: 2.9em;
  border-radius: 4px;
  border: 1px solid #b6d8f1;
  color: #6a6f77;
  outline: 0;
}
</style>
<style scope>
</style>
