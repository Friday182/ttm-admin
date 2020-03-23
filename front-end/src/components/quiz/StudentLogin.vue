<template>
  <q-form>
    <q-dialog
      v-model="showlogin"
      persistent
    >
      <q-card style="width: 450px; max-width: 80vw;">
        <q-card-actions align="right">
          <div class="row">
            <q-space />
            <q-btn
              flat
              dense
              icon="cancel"
              color="primary"
              @click="signinClose()"
            />
          </div>
        </q-card-actions>
        <q-card-section align="center">
          <q-btn
            flat
            no-caps
            dense
            size="lg"
            style="color: green;"
            label="Student Login"
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
            label="User Nmae:"
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
            q-ma-md
            rounded
            outlined
            color="green"
            label="Password:"
            :type="isPwd ? 'password' : 'text'"
            hint="* At least 6 mixed letters and numbers"
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
          class="bg-white text-teal"
        >
          <q-btn
            label="OK"
            @click="toStudentLogin(5000)"
          />
        </q-card-actions>
        <q-card-actions
          class="bg-white text-teal"
          align="center"
        >
          <q-checkbox
            v-model="rememberUser"
            keep-color
            left-label
            color="blue"
            label="Remember you login details on this computer?"
          />
        </q-card-actions>
        <q-card-actions
          align="center"
          class="text-blue"
        >
          <signin-list
            :user-type="listName"
            @listSelected="toStudentLogin"
          />
        </q-card-actions>
      </q-card>
      <alert-msg
        :active="alert"
        :msg="alertMsg"
        @cancelAlert="alertClose"
      />
    </q-dialog>
  </q-form>
</template>

<script type="text/javascript">
import { mapGetters, mapActions, mapMutations } from 'vuex'
import { STUDENT_LOGIN_MUTATION } from '../../graphql/mutations'

export default {
  name: 'StudentLogin',
  components: {
    'alert-msg': require('components/common/AlertMsg.vue').default,
    'signin-list': require('components/common/SigninList.vue').default
  },
  props: {
    showlogin: Boolean
  },
  data () {
    return {
      listName: 'student',
      alert: false,
      alertMsg: '',
      rememberUser: true,
      password: '',
      isPwd: true,
      skipQueryStudent: true,
      token: '',
      loginName: '',
      mockStudents: [
        {
          id: 1,
          mentorEmail: '',
          name: 'Lucas Yin',
          password: '',
          token: ''
        },
        {
          id: 2,
          mentorEmail: '',
          name: 'Jamie Yin',
          password: '',
          token: ''
        },
        {
          id: 3,
          mentorEmail: '',
          name: 'Juncheng Yin',
          password: '',
          token: ''
        }
      ],
      student: {
        mentorId: 0,
        mentorEmail: '',
        name: '',
        age: 0,
        city: '',
        country: '',
        school: '',
        stickers: 0,
        stars: 0,
        membershipDate: '',
        expireDate: '',
        lastLoginDate: '',
        subjects: [],
        stickerLog: [],
        contacts: []
      }
    }
  },
  computed: {
    ...mapGetters('localStudents', ['localStudents']),
    localStudentList: function () {
      if (this.localStudents.length > 0) {
        return this.localStudents
      } else {
        return []
      }
    }
  },
  mounted () {
    // console.log('student login mounted, find local students: ' + this.localStudentList.length)
  },
  methods: {
    ...mapActions('student', ['updateStudent']),
    ...mapActions('currentUser', ['updateUser']),
    ...mapMutations('localStudents', ['removeStudent', 'addStudent']),
    ...mapMutations('logs', ['copyLogs']),
    ...mapActions('localStudents', ['updateLocalStudents', 'updateLocalStudentDate']),

    addByEnter (opt) {
      if (this.loginName !== '' && this.password !== '') {
        this.studentLogin()
      } else {
        if (this.loginName !== '') {
          this.$refs.pwInput.focus()
        } else {
          this.$refs.nameInput.focus()
        }
      }
    },
    toStudentLogin (key) {
      console.log('login/read student data for key - ' + key)
      if (key !== 5000) {
        this.loginName = this.localStudentList[key].name
        this.password = this.localStudentList[key].password
        this.studentLogin()
      } else {
        if (!this.loginName || !this.password) {
          console.log('Invalid name or password')
        } else {
          this.studentLogin()
        }
      }
    },
    studentLogin () {
      this.$apollo
        .mutate({
          mutation: STUDENT_LOGIN_MUTATION,
          variables: {
            name: this.loginName,
            password: this.password,
            token: ''
          }
        })
        .then(response => {
          if (!response.errors) {
            if (response.data.StudentLogin.Gid) {
              console.log('studentLogin for - ' + response.data.StudentLogin.Name)
              this.setCurrentStudent(response.data.StudentLogin)
            } else {
              this.alertMsg = response.errors.message
              this.alert = true
            }
          } else {
            console.log('reponse error', response.errors)
          }
        })
        .catch(error => {
          this.alertMsg = 'Invalid Name or Password, Check with your mentor please.'
          this.alert = true
          console.log(error)
        })
    },
    setCurrentStudent (obj) {
      this.skipQueryStudent = true
      if (obj.MentorEmail !== '') {
        console.log('Write current student to store - ' + obj.Name)
        this.updateStudent({
          gid: obj.Gid,
          mentorEmail: obj.MentorEmail,
          name: obj.Name,
          age: obj.Age,
          city: obj.City,
          country: obj.Country,
          school: obj.School,
          stickers: obj.Stickers,
          stars: obj.Stars,
          membershipActive: obj.IsMember,
          membershipDate: obj.MembershipStart,
          expireDate: obj.MembershipStop,
          lastLoginDate: obj.LastLoginTime,
          subjects: JSON.parse(obj.SubjectState),
          quizzes: [],
          quizDone: false,
          stickerLog: obj.StickerLog,
          contacts: obj.Contacts
        })
        this.updateUser({
          id: obj.Gid,
          name: obj.Name,
          contacts: obj.Contacts,
          sessionLogin: 'student'
        })
        this.saveUserLocal()
        this.loadTaskLogs(obj.Gid)
        if (this.$route.path !== '/StudentHome') {
          this.$router.push('/StudentHome')
        }
        this.signinClose()
      } else {
        console.log('something wrong')
      }
    },
    saveUserLocal () {
      let newUser = true
      for (let i = 0; i < this.localStudentList.length; i++) {
        if (this.localStudentList[i].name === this.loginName) {
          console.log('find student in localStudentList name - ' + this.loginName)
          newUser = false
          if (this.rememberUser === false) {
            this.removeStudent(i)
          } else {
            this.updateLocalStudentDate(i)
          }
        }
      }
      if (newUser === true && this.rememberUser === true) {
        this.addStudent({
          name: this.loginName,
          password: this.password,
          letter: this.loginName[0],
          lastSigninDate: Date()
        })
      }

      // Save to local storage and set session storage for apollo
      let students = JSON.stringify(this.localStudentList)
      console.log('save in local students - ' + students)
      localStorage.setItem('students', students)
    },
    loadTaskLogs (gid) {
      let item = gid + '-logs'
      let tmpLog = localStorage.getItem(item)
      console.log('localstorage logs - ', item, 'logs: ', tmpLog)
      if (tmpLog) {
        let localLogs = JSON.parse(tmpLog)
        if (localLogs.length > 0) {
          this.copyLogs(localLogs)
        }
      }
    },
    signinClose () {
      this.loginName = ''
      this.password = ''
      this.$emit('studentLoginClose')
    },
    alertClose () {
      this.alert = false
      this.alertMsg = ''
    }
  }
}
</script>

<style>
</style>
