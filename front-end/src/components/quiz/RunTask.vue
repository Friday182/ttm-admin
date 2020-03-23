<template>
  <div class="q-pa-md q-gutter-xl">
    <q-dialog
      v-model="active"
      persistent
      :maximized="true"
      transition-show="slide-up"
      transition-hide="slide-down"
    >
      <q-card
        class="bg-gray text-black"
        align="center"
        style="background: radial-gradient(circle, #FFFFFF 75%, #C0F6F2 100%)"
      >
        <q-card-section>
          <div class="row q-mx-md">
            <q-space />
            <q-btn
              class="q-pa-sm glossy"
              round
              color="primary"
            >
              {{ taskLog.correctPerc }}%
              <q-tooltip
                content-class="bg-white text-primary"
              >
                Correct Percentage
              </q-tooltip>
            </q-btn>
            <q-btn
              class="q-ml-sm q-mr-xl q-pa-sm glossy"
              round
              color="amber"
            >
              {{ taskLog.usedSeconds }}
              <q-tooltip
                content-class="bg-white text-primary"
              >
                Used Time in Seconds
              </q-tooltip>
            </q-btn>
            <q-btn
              align="right"
              dense
              flat
              round
              icon="close"
              color="red"
              @click="closeTask"
            >
              <q-tooltip
                content-class="bg-white text-primary"
              >
                Close
              </q-tooltip>
            </q-btn>
          </div>
        </q-card-section>
        <q-card-section>
          <div class="q-px-xl">
            <div class="q-px-xl">
              <q-linear-progress
                stripe
                rounded
                style="height: 10px"
                :value="progressValue"
                color="primary"
              />
            </div>
          </div>
        </q-card-section>
        <q-card-section>
          <div class="row text-h8">
            <q-space />
            Question - {{ taskLog.queIndex }}
            <q-space />
            <q-btn
              v-if="isHelper"
              class="q-mr-md"
              round
              dense
              color="primary"
              icon="help_outline"
              @click="showHelper"
            />
          </div>
        </q-card-section>
        <q-card-section
          v-if="question_type === 'M_COM'"
          key="que_type"
        >
          <question-math-common
            :index="localQueIdx"
          />
        </q-card-section>
        <q-card-section
          v-if="question_type === 'M_CLK'"
          key="que_type"
        >
          <question-math-clock
            :index="localQueIdx"
          />
        </q-card-section>
        <q-card-section
          v-if="question_type === 'M_MONEY'"
          key="que_type"
        >
          <question-math-money
            :index="localQueIdx"
            :status="ansStatus"
            :sta-color="statusColor"
            @ansCorrect="getAnswer"
          />
        </q-card-section>
        <q-card-section
          v-else-if="isWordQuestion"
          key="que_type"
        >
          <question-en-wl
            :index="localQueIdx"
            :trigger="wlTrigger"
            @evShowQuestion="showEnWlQuesion"
          />
        </q-card-section>
        <q-card-section
          v-if="question_type === 'E_COM'"
          key="que_type"
        >
          <question-en-common
            :index="localQueIdx"
          />
        </q-card-section>
        <q-card-section
          v-if="question_type === 'M_CHART'"
          key="que_type"
        >
          <question-math-chart
            :index="localQueIdx"
          />
        </q-card-section>
        <q-card-section
          v-if="question_type === 'M_TABLE'"
          key="que_type"
        >
          <question-math-table
            :index="localQueIdx"
          />
        </q-card-section>
        <q-card-section
          v-if="question_type === 'M_SHAPE'"
          key="que_type"
        >
          <question-math-shape
            :index="localQueIdx"
          />
        </q-card-section>

        <q-card-section style="width: 50%">
          <q-linear-progress
            v-if="question_type !=='M_MONEY'"
            :value="ansStatus"
            :color="statusColor"
            class="q-mt-sm"
          />
        </q-card-section>

        <q-card-section
          v-if="answer_type === 'SC'"
          key="answer-type"
          class="q-mt-lg"
        >
          <answer-sc
            :useradio="useRadio"
            :index="localQueIdx"
            @ansCorrect="getAnswer"
          />
        </q-card-section>
        <q-card-section
          v-else-if="answer_type === 'IT'"
          key="answer-type"
          class="q-mt-lg"
        >
          <answer-it
            :index="localQueIdx"
            @ansCorrect="getAnswer"
          />
        </q-card-section>
        <q-card-section
          v-else-if="answer_type === 'IT_FRAC'"
          key="answer-type"
          class="q-mt-lg"
        >
          <answer-it-frac
            :index="localQueIdx"
            @ansCorrect="getAnswer"
          />
        </q-card-section>
        <q-card-section
          v-else-if="answer_type === 'TF'"
          key="answer-type"
          class="q-mt-xl"
        >
          <answer-tf
            :index="localQueIdx"
            @ansCorrect="getAnswer"
          />
        </q-card-section>
        <q-card-section
          v-else-if="answer_type === 'CS'"
          key="answer-type"
          class="q-mt-xl"
        />
        <q-card-section
          v-else
          key="answer-type"
          class="q-mt-xl"
        >
          <q-spinner-hourglass
            color="blue"
            size="4em"
          />
        </q-card-section>
        <q-toggle
          v-if="devMode"
          v-model="developSwitch"
          dense
          color="orange"
          label="Dev Mode"
        />
        <q-card-section v-if="developSwitch">
          <q-btn
            class="q-mt-md"
            style="width: 50%"
            no-caps
          >
            <vue-mathjax
              :formula="formula"
              :options="{}"
            />
          </q-btn>
        </q-card-section>
        <q-card-section v-if="developSwitch">
          <q-btn
            flat
            color="primary"
          >
            Current Task: {{ taskid }}
          </q-btn>
          <q-input
            v-model="formula"
            outlined
            label="formula"
            style="width: 300px"
          />
        </q-card-section>
      </q-card>
    </q-dialog>
    <alert
      :active="alert"
      :msg="alertMsg"
      @cancelAlert="alertClose"
    />
  </div>
</template>
<script>
import { TASK_QUESTIONS_QUERY } from '../../graphql/queries'
import { ADD_TASKLOG_MUTATION } from '../../graphql/mutations'
import { mapGetters, mapActions } from 'vuex'
import { VueMathjax } from 'vue-mathjax'
import { apolloProvider } from '../../boot/apollo'

export default {
  components: {
    'vue-mathjax': VueMathjax,
    'question-math-common': require('components/student/QuestionMathCommon.vue').default,
    'question-math-clock': require('components/student/QuestionMathClock.vue').default,
    // 'question-math-clock': require('components/student/QuestionMathClock1.vue').default,
    'question-en-wl': require('components/student/QuestionEnWordLearn.vue').default,
    'question-math-money': require('components/student/QuestionMathMoney.vue').default,
    'question-math-chart': require('components/student/QuestionMathChart.vue').default,
    'question-math-table': require('components/student/QuestionMathTable.vue').default,
    'question-math-shape': require('components/student/QuestionMathShape.vue').default,
    'answer-sc': require('components/student/AnswerSc.vue').default,
    'answer-it': require('components/student/AnswerIt.vue').default,
    'answer-it-frac': require('components/student/AnswerItFrac.vue').default,
    'answer-tf': require('components/student/AnswerTf.vue').default,
    'question-en-common': require('components/student/QuestionEnCommon.vue').default,
    'alert': require('components/common/AlertMsg.vue').default
  },
  props: {
    active: {
      type: Boolean,
      default: false
    },
    taskid: {
      type: String,
      default: ''
    },
    subject: {
      type: Number,
      default: 0
    },
    mode: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      devMode: true,
      developSwitch: false,
      isChart: false,
      isHelper: false,
      isFomula: true,
      numQue: 0,
      ansStatus: 0,
      incFactor: 1,
      demoCount: 0,
      alert: false,
      alertMsg: '',
      useRadio: false,
      statusColor: '',
      progressValue: 0,
      maximizedToggle: true,
      milliSec: 0,
      updateStatusMs: 0,
      formula: '',
      skipQueryQuestions: true,
      question_type: '',
      answer_type: '',
      localQueIdx: 0,
      task_done_count: 0,
      studenGid: '',
      queSecond: 0,
      queList: [],
      taskLog: {
        queTotal: 0,
        queIndex: 0,
        correctNum: 0,
        wrongNum: 0,
        correctPerc: 0,
        usedSeconds: 0,
        totalScore: 0,
        results: [] // log result for each question: task, que_id, time_sec, correct, tags
      },
      details: [],
      finishMsg: '',
      finishAlert: false
    }
  },
  computed: {
    ...mapGetters('student', ['currentStudent']),
    enQueIdx: function () {
      return this.taskLog.queIndex
    },
    isWordQuestion: function () {
      if (this.question_type === 'E_WL' || this.question_type === 'E_WT') {
        return true
      } else {
        return false
      }
    }
  },
  watch: {
    active: function (newVal, oldVal) {
      console.log('taskid in watch: ' + this.taskid, 'mode - ' + this.mode)
      if (newVal === true) {
        if (this.mode === 1) {
          // run Demo
          this.numQue = 3
        } else {
          // use default 50
          this.numQue = 0
        }
        if (this.taskid === 'EVM') {
          console.log('Run EVM')
          this.runEvaluaton()
        } else {
          this.cleanTask()
          this.fetchNewQues()
        }
      }
    }
  },
  created () {
    setInterval(this.updateSecs, 100)
  },
  destroyed () {
    console.log('destroyed?')
  },
  methods: {
    ...mapActions('questions', ['setQuestions', 'clearQuestions']),
    ...mapActions('student', ['updateStudent']),
    // Display next questions
    showQuestion (queIdx) {
      console.log('show question - ' + queIdx + 'total: ' + this.taskLog.queTotal)
      console.log('question type - ', this.queList[queIdx].QuestionType)
      if (queIdx >= this.taskLog.queTotal) {
        // Complete all question, send taskLog and exit
        this.finishTaskRun(1)
      }
      this.localQueIdx = queIdx
      this.question_type = this.queList[queIdx].QuestionType
      if (this.question_type === 'E_WL' || this.question_type === 'E_WT') {
        this.wlTrigger = queIdx + 1
        console.log('triggered?', this.wlTrigger)
      }
      if (this.queList[queIdx].QuestionType === 'E_WL') {
        this.answer_type = 'na'
      } else {
        this.answer_type = this.queList[queIdx].AnswerType
      }
      this.formula = this.queList[queIdx].Formula

      // todo: will update helper
      if (this.queList[queIdx].Helper === '') {
        this.isHelper = true
      } else {
        this.isHelper = true
      }
      this.taskLog.queIndex = queIdx + 1
      this.queSecond = 0
      if (queIdx === this.taskLog.queTotal - 2) {
        // this.fetchNewQues()
      }
    },
    getAnswer (isCorrect) {
      console.log('Process returned answer ')
      if (isCorrect === true) {
        this.statusColor = 'green'
        this.taskLog.correctNum++
      } else {
        this.statusColor = 'red'
        this.taskLog.wrongNum++
      }
      this.ansStatus = 1
      this.updateStatusMs = this.milliSec
      this.taskLog.correctPerc = Math.floor((this.taskLog.correctNum / this.taskLog.queIndex) * 100)
      if (this.queList[this.taskLog.queIndex - 1].QueIdx !== 0) {
        this.taskLog.results.push({
          kp: this.queList[this.taskLog.queIndex - 1].Kp,
          queIdx: this.queList[this.taskLog.queIndex - 1].QueIdx,
          isCorrect: isCorrect,
          usedSec: this.queSecond,
          stdSec: this.queList[this.taskLog.queIndex - 1].StdSec,
          tags: ''
        })
      }
      if (this.updateScore(isCorrect) === true) {
        this.showQuestion(this.taskLog.queIndex)
      }
    },
    closeTask () {
      this.finishTaskRun(0)
    },
    showEnWlQuesion () {
      this.answer_type = this.queList[this.taskLog.queIndex - 1].AnswerType
      console.log('answer type changed - ', this.answer_type)
    },
    updateSecs: function () {
      this.milliSec++
      if (this.milliSec % 10 === 0) {
        this.queSecond++
        this.taskLog.usedSeconds++
        if (this.taskLog.usedSeconds > 1860) {
          this.finishTaskRun(3)
        }
      }
      if (this.milliSec - this.updateStatusMs === 10) {
        this.ansStatus = 0
        // this.statusColor = 'grey-1'
      }
    },
    updateScore (result) {
      let increament = 0
      let isNextQue = true
      let queIdx = this.taskLog.queIndex - 1
      if (this.mode === 1) {
        this.taskLog.totalScore += 35
      } else if ((this.taskid.charAt(1) === 'Q') || this.currentStudent.subjects[this.subject].level > 6) {
        increament = 100 / this.taskLog.queTotal
      } else {
        if (result === true) {
          increament = 1
          increament += Math.round(this.queList[queIdx].StdSec / 10)
          if (this.taskLog.correctPerc > 85) {
            if (this.taskLog.usedSeconds < this.queList[queIdx].StdSec * this.taskLog.queIndex) {
              increament++
            }
          }
          if (increament > 5) {
            increament = increament - 2
          }
          this.taskLog.totalScore += increament * this.incFactor
        } else {
          if (this.taskLog.totalScore > 3) {
            this.taskLog.totalScore -= 3
          }
        }
      }
      console.log('increament - ', increament, 'totalScore - ', this.taskLog.totalScore)
      // this is only for debug/testing purpose
      // this.taskLog.totalScore += 8
      if (this.taskLog.totalScore >= 100) {
        this.taskLog.totalScore = 100
        isNextQue = false
        // Complete this task and send taskLog to server
        this.finishTaskRun(2)
      } else {
        this.progressValue = this.taskLog.totalScore / 100
        console.log('progress value: ' + this.progressValue)
      }
      return isNextQue
    },
    finishTaskRun (opt) {
      if (this.mode === 1) {
        // This is DEMO RUN
        this.demoCount++
        if (this.demoCount % 5 === 0) {
          this.alertMsg = 'Well Done! Time to Sign Up!'
          this.alert = true
          opt = 0
        }
      } else {
        if (opt === 1) {
          // no more question? send result
          if (this.taskLog.queIndex > 8) {
            this.addTaskLog()
          }
          this.alertMsg = 'Something wrong, please report to your mentor.'
          this.alert = true
        } else if (opt === 2) {
          // Complete this task, send result
          this.addTaskLog()
          this.alertMsg = 'Task Completed, Well Done!'
          this.alert = true
        } else if (opt === 3) {
          // no more question? send result
          if (this.taskLog.queIndex > 8) {
            this.addTaskLog()
          }
          this.alertMsg = 'Well done, but no time for today, try it next time.'
          this.alert = true
        } else {
          // Closed by user, check if send result
          if (this.taskLog.queIndex > 3) {
            this.addTaskLog()
          }
          this.alertMsg = 'Task Not finish, are you sure to quit now?'
          this.alert = false
        }
      }
      this.cleanTask()
      this.$emit('completed', opt)
    },
    alertClose () {
      this.alertMsg = ''
      this.alert = false
    },
    showHelper () {
      console.log('show helper info')
    },
    fetchNewQues () {
      if (this.taskid !== '') {
        console.log('Run query: ', this.taskid, 'numberQue - ', this.numQue)
        this.studenGid = this.currentStudent.gid
        if (this.mode === 1) {
          this.studenGid = '1000'
        }
        this.skipQueryQuestions = false
        this.$apollo.queries.GetQuestions.refetch({
          gid: this.studenGid,
          sub: this.subject,
          kp: this.taskid,
          numQues: this.numQue
        })
      } else {
        this.alertMsg = 'This task is not available yet!'
        this.alert = true
      }
    },
    // Add new questions in the queue
    updateQuestons (newQues) {
      console.log('in update questons: ' + newQues[0].Kp, 'length - ', newQues.length)
      if (newQues.length > 2) {
        this.shuffleArray(newQues)
        for (let i = 0; i < newQues.length; i++) {
          this.queList.push(newQues[i])
        }
        this.taskLog.queTotal += newQues.length
        // Store question list in store, so it will be available for other components
        this.setQuestions(this.queList)
        if (this.taskLog.queIndex === 0) {
          // Show the first question
          this.showQuestion(this.taskLog.queIndex)
        }
      } else {
        console.log('NO enough questions')
        this.finishTaskRun(0)
      }
    },
    shuffleArray (array) {
      for (let i = array.length - 1; i > 0; i--) {
        let j = Math.floor(Math.random() * (i + 1))
        let temp = array[i]
        array[i] = array[j]
        array[j] = temp
      }
    },
    addTaskLog () {
      let kpArray = this.taskid.split(',')
      let details = []
      for (let k = 0; k < kpArray.length; k++) {
        details.push({
          kp: kpArray[k],
          numQue: 0,
          correctQue: 0,
          maxSec: 0,
          minSec: 0,
          resultList: [],
          totalSec: 0,
          stdSec: 0
        })
      }
      for (let i = 0; i < this.taskLog.results.length; i++) {
        for (let j = 0; j < details.length; j++) {
          if (this.taskLog.results[i].kp === details[j].kp) {
            details[j].numQue++
            if (this.taskLog.results[i].isCorrect === true) {
              details[j].correctQue++
            }
            details[j].resultList.push(this.taskLog.results[i])
            details[j].totalSec += this.taskLog.results[i].usedSec
            details[j].stdSec += this.taskLog.results[i].stdSec
            if (details[j].maxSec < this.taskLog.results[i].usedSec) {
              details[j].maxSec = this.taskLog.results[i].usedSec
            }
            if (details[j].minSec > this.taskLog.results[i].usedSec) {
              details[j].minSec = this.taskLog.results[i].usedSec
            }
          }
        }
      }
      this.$apollo
        .mutate({
          mutation: ADD_TASKLOG_MUTATION,
          variables: {
            gid: this.currentStudent.gid,
            kps: this.taskid,
            totalQue: this.taskLog.wrongNum + this.taskLog.correctNum,
            totalCorrect: this.taskLog.correctNum,
            totalSec: this.taskLog.usedSeconds,
            totalScore: this.taskLog.totalScore,
            details: JSON.stringify(details)
          }
        })
        .then(response => {
          if (response.data.AddTaskLog.Gid === this.currentStudent.gid) {
            console.log('get back stickers -', response.data.AddTaskLog.Stickers)
            if (response.data.AddTaskLog.Stickers !== this.currentStudent.stickers) {
              this.updateStudent({
                stickers: response.data.AddTaskLog.Stickers,
                starts: response.data.AddTaskLog.Stars
              })
            }
          } else {
            this.alertMsg = 'Add task log failed, please report to your mentor.'
            this.alert = true
          }
        })
    },
    cleanTask () {
      // Note: Reset apollo cache here, otherwise, the previous query data may be returned
      apolloProvider.defaultClient.cache.reset()
      this.ansStatus = 0
      this.statusColor = ''
      this.progressValue = 0
      this.milliSec = 0
      this.updateStatusMs = 0
      this.formula = ''
      this.skipQueryQuestions = true
      this.answer_type = ''
      this.queList = []
      this.taskLog.queTotal = 0
      this.taskLog.queIndex = 0
      this.taskLog.correctNum = 0
      this.taskLog.wrongNum = 0
      this.taskLog.correctPerc = 0
      this.taskLog.usedSeconds = 0
      this.taskLog.totalScore = 0
      // clear data in store
      this.clearQuestions()
    },
    runEvaluaton () {
      let age = this.currentStudent.age
      if (age < 7) {
        this.taskid = 'MA5'
      }
    }
  },
  apollo: {
    GetQuestions: {
      query: TASK_QUESTIONS_QUERY,
      variables () {
        return {
          gid: this.studenGid,
          sub: this.subject,
          kp: this.taskid,
          numQues: this.numQue
        }
      },
      error (error) {
        console.error('We\'ve got an error!', error)
      },
      skip () {
        return this.skipQueryQuestions
      },
      result (data, key) {
        this.skipQueryQuestions = true
        if (data.data.GetQuestions) {
          console.log('questions - ', data.data.GetQuestions[0].Kp)
          this.updateQuestons(data.data.GetQuestions)
        } else {
          this.finishTaskRun(1)
        }
      }
    }
  }
}
</script>

<style lang="stylus" scoped>
.my-card
  width 100%
  max-width 1000px
.q-input {
  font-size: 24px;
}
</style>
