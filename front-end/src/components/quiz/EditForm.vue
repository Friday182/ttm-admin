<template>
  <q-card>
    <q-card-section>
      <div class="row q-ml-md">
        <div class="col-3 q-mr-sm">
          <q-input
            v-model="inQueIdx"
            outlined
            dense
            label="Question Index"
            @input="updateCurrentQuestion"
          />
        </div>
        <div class="col-2 q-mr-sm">
          <q-input
            v-model="inStdSec"
            outlined
            dense
            type="number"
            label="Std Second"
            @input="updateCurrentQuestion"
          />
        </div>
        <div class="col-3 q-mr-sm">
          <q-select
            v-model="inQueType"
            outlined
            dense
            label="Question Type"
            :options="queTypeOptions"
            @input="updateCurrentQuestion"
          />
        </div>
        <div class="col-3 q-mr-sm">
          <q-select
            v-model="inAnsType"
            outlined
            dense
            label="Answer Type"
            :options="ansTypeOptions"
            @input="updateCurrentQuestion"
          />
        </div>
      </div>
    </q-card-section>
    <q-card-section class="q-ml-md">
      <q-input
        v-model="inUpText"
        outlined
        type="textarea"
        label="Up Text"
      >
        <template v-slot:append>
          <q-btn
            round
            dense
            color="blue"
            icon="send"
            @click="updateCurrentQuestion"
          />
        </template>
      </q-input>
    </q-card-section>
    <q-card-section class="q-ml-md">
      <q-input
        v-model="inFormula"
        outlined
        label="Formula"
        @input="updateCurrentQuestion"
      />
    </q-card-section>
    <q-card-section class="q-ml-md">
      <q-btn
        class="q-mb-xs"
        color="blue"
        icon="send"
        label="Save"
        @click="updateCurrentQuestion"
      />
      <v-jsoneditor
        v-model="inJsonText"
        :plus="false"
        height="600px"
        :options="vjsonOptions"
      />
    </q-card-section>
    <q-card-section class="q-ml-md">
      <q-input
        v-model="inDownText"
        outlined
        type="textarea"
        label="Down Text"
        @input="updateCurrentQuestion"
      >
        <template v-slot:append>
          <q-btn
            round
            dense
            color="blue"
            icon="send"
            @click="updateCurrentQuestion"
          />
        </template>
      </q-input>
    </q-card-section>
    <q-card-section class="q-ml-md">
      <q-input
        v-model="inOption1"
        outlined
        label="Option A"
        @input="updateCurrentQuestion"
      />
      <q-input
        v-model="inOption2"
        class="q-mt-xs"
        outlined
        label="Option B"
        @input="updateCurrentQuestion"
      />
      <q-input
        v-model="inOption3"
        class="q-mt-xs"
        outlined
        label="Option C"
        @input="updateCurrentQuestion"
      />
      <q-input
        v-model="inOption4"
        class="q-mt-xs"
        outlined
        label="Option D"
        @input="updateCurrentQuestion"
      />
      <q-input
        v-model="inOption5"
        class="q-mt-xs"
        outlined
        label="Option E"
        @input="updateCurrentQuestion"
      />
    </q-card-section>
    <q-card-section class="q-ml-md">
      <div class="row">
        <div class="col-3">
          <q-select
            v-model="inAnswer"
            outlined
            dense
            label="Answer"
            :options="answerOptions"
            @input="updateCurrentQuestion"
          />
        </div>
      </div>
    </q-card-section>
    <q-card-section class="q-ml-md">
      <q-select
        outlined
        v-model="inTags"
        multiple
        :options="kpOptions"
        use-chips
        stack-label
        label="KP Dependency"
        @input="updateCurrentQuestion"
      />
    </q-card-section>
    <q-card-section class="q-ml-md">
      <div class="row">
        <q-space />
        <q-btn
          label="Save Quiz"
          no-caps
          color="primary"
          style="font: 120% Arial bold"
          @click="saveQuestion()"
        />
        <q-btn
          class="q-ml-xl"
          label="Clear"
          no-caps
          color="primary"
          style="font: 120% Arial bold"
          @click="clearForm()"
        />
      </div>
    </q-card-section>
  </q-card>
</template>

<script>
import VJsoneditor from 'v-jsoneditor'
import { ADD_QUIZ_MUTATION } from '../../graphql/mutations'
import { mapMutations, mapGetters } from 'vuex'

export default {
  name: 'AddQuiz',
  props: {
    index: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      inQueIdx: '',
      inStdSec: '',
      inQueType: '',
      inAnsType: '',
      inComment: '',
      inUpText: '',
      inDownText: '',
      inFormula: '',
      inJsonText: '',
      inOption1: '',
      inOption2: '',
      inOption3: '',
      inOption4: '',
      inOption5: '',
      inAnswer: '',
      inTags: ['MA24'],
      usedSeconds: 0,
      vjsonOptions: {
      },
      skipQueryAddQuiz: true,
      queTypeOptions: ['M_COM', 'M_TABLE', 'M_SHAPE', 'M_CLK', 'M_CHART', 'M_MONEY', 'M_IMG'],
      ansTypeOptions: ['SC', 'IT', 'TF', 'MC'],
      answerOptions: ['A', 'B', 'C', 'D', 'E'],
      kpOptions: ['MA24', 'MC16', 'MF5'],
      exampleTable: [{
        tableTitle: '',
        columns: ['columns'],
        tableData: ['tableData'],
        style: 'width: 40%',
        separator: 'cell'
      }],
      exampleShape: [{
        tableTitle: '',
        columns: ['columns'],
        tableData: ['tableData'],
        style: 'width: 40%',
        separator: 'cell'
      }],
      period: ''
    }
  },
  computed: {
    ...mapGetters('currentQuestion', ['currentQuestion']),
    editQuestion: function () {
      return this.currentQuestion
    }
  },
  watch: {
    editQuestion: function (newVal, oldVal) {
      console.log('uptext objects - ', newVal.UpTexts.length)
      let tmpUpText = ''
      for (let i = 0; i < newVal.UpTexts.length; i++) {
        tmpUpText += newVal.UpTexts[i] + '\n'
      }
      let tmpDownText = ''
      for (let i = 0; i < newVal.DownTexts.length; i++) {
        tmpDownText += newVal.DownTexts[i] + '\n'
      }
      this.inQueIdx = newVal.QueIdx
      this.inStdSec = newVal.StdSec
      this.inQueType = newVal.QuestionType
      this.inAnsType = newVal.AnswerType
      this.inComment = newVal.Tips
      this.inUpText = tmpUpText
      this.inDownText = tmpDownText
      this.inFormula = newVal.Formula
      this.inOption1 = newVal.Options[0]
      this.inOption2 = newVal.Options[1]
      this.inOption3 = newVal.Options[2]
      this.inOption4 = newVal.Options[3]
      this.inOption5 = newVal.Options[4]
      this.inTags = newVal.Tags
      let tmpJson = []
      let tmpOrgJson = []
      if (newVal.QuestionType === 'M_TABLE') {
        tmpOrgJson = newVal.Tables
      } else if (newVal.QuestionType === 'M_CHART') {
        tmpOrgJson = newVal.Charts
      } else if (newVal.QuestionType === 'M_SHAPE') {
        tmpOrgJson = newVal.Shapes
      } else if (newVal.QuestionType === 'M_COM') {
        tmpOrgJson = newVal.Formula
      } else {
        tmpOrgJson = ['']
      }
      for (let i = 0; i < tmpOrgJson.length; i++) {
        tmpJson.push(JSON.parse(tmpOrgJson[i]))
      }
      this.inJsonText = tmpJson
    }
  },
  components: {
    VJsoneditor
  },
  created () {
    setInterval(this.updateSecs, 1000)
  },
  methods: {
    ...mapMutations('currentQuestion', ['doSetCurrentQuestion']),
    // ...mapActions('currentQuestion', ['setCurrentQuestion']),
    saveQuestion () {
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
              this.updateAddQuiz()
            }
          } else {
            console.log('reponse error', response.errors)
          }
        })
        .catch(error => {
          console.log(error)
        })
    },
    updateAddQuiz () {
      console.log('add quiz')
      this.addNewQuiz({
        QuizId: this.inId.trim().toUpperCase(),
        Grade: parseInt(this.inGrade),
        Desc: this.inDesc,
        Operator: this.inOperator,
        Comment: this.inComment
      })
      this.clearForm()
    },
    onJsonChange (value) {
      console.log('value:', value)
    },
    updateSecs: function () {
      this.usedSeconds++
    },
    jsonChanged: function (opt) {
      console.log('write new data', opt)
    },
    updateCurrentQuestion: function () {
      // save in current question, this should be done in editform
      console.log('write new data')
      let tmp = []
      tmp.push(this.inOption1)
      tmp.push(this.inOption2)
      tmp.push(this.inOption3)
      tmp.push(this.inOption4)
      let tmp1 = []
      tmp1.push(this.inAnswer)
      let tmpCharts = []
      let tmpClk = []
      let tmpTables = []
      let tmpShapes = []
      let tmpUpText = this.inUpText.split('\n')
      let tmpDownText = this.inDownText.split('\n')
      if (this.inQueType === 'M_TABLE') {
        tmpTables.push(this.inJsonText)
      } else if (this.inQueType === 'M_CHART') {
        tmpCharts.push(this.inJsonText)
      } else if (this.inQueType === 'M_CLK') {
        tmpClk.push(this.inJsonText)
      } else if (this.inQueType === 'M_SHAPE') {
        for (let i = 0; i < this.inJsonText.length; i++) {
          tmpShapes.push(JSON.stringify(this.inJsonText[i]))
        }
      }
      console.log('Save to current question')
      this.doSetCurrentQuestion(
        {
          Kp: this.QuizId,
          QueIdx: this.inQueIdx,
          StdSec: this.inStdSec,
          QuestionType: this.inQueType,
          AnswerType: this.inAnsType,
          UpTexts: tmpUpText,
          DownTexts: tmpDownText,
          Formula: this.inFormula,
          Options: tmp,
          Tags: this.inTags,
          Answers: tmp1,
          Charts: tmpCharts,
          Clocks: tmpClk,
          Tables: tmpTables,
          Shapes: tmpShapes,
          AnswerText: '',
          Helper: false,
          Imgs: [],
          Tips: this.inComment,
          Choice: ''
        }
      )
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
