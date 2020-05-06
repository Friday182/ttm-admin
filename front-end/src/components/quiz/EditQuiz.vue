<template>
  <q-page class="flex-center">
    <q-dialog
      v-model="active"
      persistent
      :maximized="true"
      transition-show="slide-up"
      transition-hide="slide-down"
    >
      <q-splitter
        v-model="splitterModel"
        unit="px"
        :limits="[94, 1024]"
        horizontal
      >
        <template v-slot:before>
          <q-card class="bg-blue-1 text-white">
            <q-toolbar
              class="bg-blue-10 text-white"
            >
              <q-btn
                color="bg-blue-10"
                no-caps
                disable
                class="q-mx-lg glossy"
              >
                {{ quizId }}
              </q-btn>
              <q-btn
                icon="add"
                color="amber"
                class="q-mx-lg glossy"
                label="Add Question"
                @click="addQuestion"
              />
              <q-space />
              <q-btn
                class="q-mr-lg glossy"
                color="amber"
              >
                {{ usedSeconds }}
              </q-btn>
              <q-btn
                icon="close"
                color="red"
                class="q-mx-lg glossy"
                label="Finish"
                @click="closeTask"
              />
            </q-toolbar>
            <div class="q-pa-xs flex flex-center">
              <q-pagination
                v-model="slide"
                :max="totalQue"
                :direction-links="true"
              />
            </div>
          </q-card>
        </template>
        <template v-slot:after>
          <q-splitter
            v-model="insideModel"
          >
            <template v-slot:before>
              <edit-form
                :kp="quizId"
              />
            </template>
            <template v-slot:after>
              <render-quiz />
            </template>
          </q-splitter>
        </template>
      </q-splitter>
    </q-dialog>
  </q-page>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'
import { TASK_QUESTIONS_QUERY } from '../../graphql/queries'

export default {
  name: 'EditQuiz',
  components: {
    'edit-form': require('components/quiz/EditForm.vue').default,
    'render-quiz': require('components/quiz/RenderQuiz.vue').default
  },
  props: {
    active: {
      type: Boolean,
      default: false
    },
    quizId: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      splitterModel: 15,
      insideModel: 40,
      slide: 1,
      usedSeconds: 20
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),
    ...mapGetters('questions', ['allQuestions']),
    totalQue: function () {
      return this.allQuestions.length
    }
  },
  watch: {
    quizId: function (newVal, oldVal) {
      console.log('require refresh: ', newVal, oldVal)
      if (newVal !== '') {
        this.GetQuestions()
      }
    },
    slide: function (newVal, oldVal) {
      console.log('next question - ', newVal)
      if (newVal > 0 && newVal <= this.allQuestions.length) {
        this.localSetCurrentQuestion(this.allQuestions[newVal - 1])
      }
    }
  },
  created () {
    setInterval(this.updateSecs, 1000)
  },
  methods: {
    ...mapActions('currentQuestion', ['setCurrentQuestion']),
    ...mapActions('questions', ['setQuestions', 'addNewQuestion', 'clearQuestions']),
    // Add new questions in the queue
    updateQuestons (newQues) {
      let queList = []
      if (newQues.length > 0) {
        for (let i = 0; i < newQues.length; i++) {
          queList.push(newQues[i])
        }
        // Store question list in store, so it will be available for other components
        this.setQuestions(queList)
      } else {
        console.log('NO enough questions')
        this.clearQuestions()
      }
    },
    localSetCurrentQuestion (que) {
      this.setCurrentQuestion(que)
    },
    addQuestion () {
      console.log('add question, current length = ', this.allQuestions.length)
      this.addNewQuestion({
        QueIdx: this.allQuestions.length + 1,
        Kp: '',
        StdSec: 30,
        AnswerType: 'SC',
        QuestionType: 'M_COM',
        UpTexts: [],
        DownTexts: [],
        Formula: [],
        Options: [],
        Answers: [],
        Tags: [],
        Charts: [],
        Clocks: [],
        Tables: [],
        Shapes: [],
        AnswerText: '',
        Helper: false,
        Imgs: [],
        Tips: [],
        Choice: '',
        Status: 0
      })
      this.slide = this.allQuestions.length
    },
    GetQuestions () {
      console.log('user and qid - ', this.currentUser.Gid, this.quizId)
      this.$apollo
        .query({
          query: TASK_QUESTIONS_QUERY,
          variables: {
            gid: this.currentUser.Gid,
            kp: this.quizId
          }
        })
        .then(response => {
          if (!response.errors) {
            if (response.data.GetQuestions) {
              this.updateQuestons(response.data.GetQuestions)
            }
          } else {
            console.log('reponse error', response.errors)
          }
        })
        .catch(error => {
          console.log(error)
        })
    },
    updateSecs: function () {
      this.usedSeconds++
    },
    closeTask () {
      console.log('close quiz edit')
      this.$emit('quizEditCompleted')
    }
  }
}
</script>

<style>
</style>
