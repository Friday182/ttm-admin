<template>
  <div class="q-px-sm">
    <q-card>
      <q-page-sticky
        position="right"
        :offset="[100, 100]"
      >
        <q-btn
          rounded
          no-caps
          glossy
          color="blue"
          label="提交"
          size="lg"
          @click="changeDone(0)"
        />
      </q-page-sticky>
      <q-card-section
        align="left"
        style="width:30%"
      >
        <q-input
          v-model="name"
          bottom-slots
          label="姓名"
          dense
        >
          <template v-slot:prepend>
            <q-icon name="face" />
          </template>
        </q-input>
        <q-input
          v-model="id"
          label="身份证号"
          class="q-mb-md"
          dense
        />
        <el-date-picker
          v-model="value2"
          type="datetime"
          placeholder="选择日期时间"
          align="right"
          :picker-options="pickerOptions"
        />
      </q-card-section>
      <q-separator />
      <q-card-section
        align="left"
        style="width:50%"
      >
        输入JSON数据：
        <pre>
          {{ jsonData }}
        </pre>
      </q-card-section>
      <q-separator />
      <q-card-section
        align="left"
        style="width:50%"
      >
        <q-btn
          label="测试注册"
          style="color:red"
          @click="testSignup"
        />
        <q-btn
          class="q-ml-md"
          label="测试登录"
          style="color:red"
          @click="testLogin"
        />
      </q-card-section>
      <q-card-section
        align="left"
        style="width:50%"
      >
        服务器返回数据：
        <pre>
          {{ info }}
        </pre>
      </q-card-section>
    </q-card>
  </div>
</template>

<script type="text/javascript">
import { mapMutations } from 'vuex'
import { signup, login, addRecord } from '../../api/api'

export default {
  name: 'StudentProfile',
  data () {
    return {
      alert: false,
      active: false,
      info: '',
      name: '',
      id: '',
      pickerOptions: {
        shortcuts: [{
          text: '今天',
          onClick (picker) {
            picker.$emit('pick', new Date())
          }
        }, {
          text: '昨天',
          onClick (picker) {
            const date = new Date()
            date.setTime(date.getTime() - 3600 * 1000 * 24)
            picker.$emit('pick', date)
          }
        }, {
          text: '一周前',
          onClick (picker) {
            const date = new Date()
            date.setTime(date.getTime() - 3600 * 1000 * 24 * 7)
            picker.$emit('pick', date)
          }
        }]
      },
      value2: ''
    }
  },
  computed: {
    jsonData: function () {
      return (
        {
          name: this.name,
          gender: '男',
          ethnic: '汉',
          id: this.id,
          dob: '20010101',
          block: '高大上小区',
          enter: '进入',
          deviceid: 'device-123',
          datetime: this.value2
        }
      )
    }
  },
  methods: {
    ...mapMutations('visitor', ['addVisitor']),
    ...mapMutations('currentInfo', ['setActiveTab']),
    changeDone (opt) {
      console.log('datetime = ', this.value2)
      // this.addVisitor(this.jsonData)
      // this.setActiveTab('1-1')
      this.postData()
    },
    postData () {
      console.log('post data')
      addRecord({ 'vistor': this.jsonData }).then(response => (this.info = response))
    },
    testSignup () {
      console.log('signup')
      signup({ 'username': this.name, 'password': 'test123' }).then(response => (this.info = response))
      /* let config = {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
        },
        withCredentials: true
      }
      this.$axios
        .post('http://127.0.0.1:8082/signup', { 'username': this.name, 'password': 'test123' }, config)
        .then(response => (this.info = response))
      */
    },
    testLogin () {
      console.log('login')
      /* let config = {
        headers: {
          'Content-Type': 'application/x-www-form-urlencoded;charset=UTF-8'
        },
        withCredentials: true
      }
      this.$axios
        .post('http://127.0.0.1:8082/login', { 'username': this.name, 'password': 'test123' }, config)
        .then(response => (this.info = response))
    }
    */
      login({ 'username': this.name, 'password': 'test123' })
        .then(response => {
          this.info = response
          if (response.data.code === 200) {
            localStorage.setItem('access_token', response.data.token)
            console.log('set token ok - ', response.data.token)
          } else {
            console.log('not ok')
          }
        })
    }
  }
}

</script>

<style>
</style>
