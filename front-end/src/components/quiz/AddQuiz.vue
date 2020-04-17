<template>
  <div>
    <div class="row q-mb-md">
      <div class="col-1 q-mr-sm">
        <el-input
          v-model="inId"
          class="myClass"
          placeholder="Quiz Id"
          clearable
        />
      </div>
      <div class="col-1 q-mr-sm">
        <el-input
          v-model="inGrade"
          class="myClass"
          placeholder="Grade"
          clearable
        />
      </div>
      <el-select
        v-model="inOperator"
        class="myClass q-mr-sm"
        filterable
        placeholder="Operator"
      >
        <el-option
          v-for="item in operatorOptions"
          :key="item.value"
          :label="item.label"
          :value="item.label"
        />
      </el-select>
      <div class="col-4 q-mr-sm">
        <el-input
          v-model="inDesc"
          class="myClass"
          placeholder="Description"
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
        label="Add Quiz"
        no-caps
        color="primary"
        style="font: 100% Arial bold"
        @click="addQuiz()"
      />
    </div>
  </div>
</template>

<script>
import { ADD_QUIZ_MUTATION } from '../../graphql/mutations'
import { mapMutations, mapGetters } from 'vuex'

export default {
  name: 'AddQuiz',
  data () {
    return {
      inId: '',
      inGrade: '',
      inDesc: '',
      inOperator: '',
      inComment: '',
      period: ''
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),
    ...mapGetters('users', ['getUserList']),
    operatorOptions: function () {
      let tmp = []
      for (let i = 0; i < this.getUserList.length; i++) {
        if (this.getUserList[i].Role === 'operator') {
          tmp.push({
            value: i,
            label: this.getUserList[i].Username
          })
        }
      }
      return tmp
    },
    loginRole: function () {
      return this.currentUser.Role
    },
    creator: function () {
      return this.currentUser.Username
    }
  },
  methods: {
    ...mapMutations('quiz', ['addNewQuiz']),
    addQuiz () {
      if (this.inId === '' || this.inGrade === '' || this.inDesc === '') {
        return
      }
      this.$apollo
        .mutate({
          mutation: ADD_QUIZ_MUTATION,
          variables: {
            QuizId: this.inId.trim().toUpperCase(),
            Grade: parseInt(this.inGrade),
            Desc: this.inDesc,
            Operator: this.inOperator,
            Comment: this.inComment
          }
        })
        .then(response => {
          if (!response.errors) {
            if (response.data.AddQuiz) {
              this.$emit('addQuizDone')
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
      this.inId = ''
      this.inGrade = ''
      this.inDesc = ''
      this.inOperator = ''
      this.inSubject = ''
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
