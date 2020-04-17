<template>
  <q-card>
    <q-card-section>
      <div class="row q-ml-md">
        <div class="col-3 q-mr-sm">
          <q-input
            v-model="inQueIdx"
            outlined
            dense
            disable
            label="Question Index"
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
            @input="updateQuestionType"
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
          label="Save Question"
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
import { ADD_QUESTION_MUTATION } from '../../graphql/mutations'
import { mapMutations, mapGetters, mapActions } from 'vuex'

export default {
  name: 'AddQuiz',
  props: {
    kp: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      inQueIdx: '1',
      inStdSec: '30',
      inQueType: 'M_COM',
      inAnsType: 'SC',
      inComment: '',
      inUpText: '',
      inDownText: '',
      inFormula: '',
      inJsonText: [],
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
      queTypeOptions: ['M_COM', 'M_TABLE', 'M_SHAPE', 'M_CHART'],
      ansTypeOptions: ['SC', 'IT', 'TF', 'MC'],
      answerOptions: ['Option A', 'Option B', 'Option C', 'Option D', 'Option E'],
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
    ...mapGetters('currentUser', ['currentUser']),
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
      tmpUpText.replace(/\n$/, '')
      let tmpDownText = ''
      for (let i = 0; i < newVal.DownTexts.length; i++) {
        tmpDownText += newVal.DownTexts[i] + '\n'
      }
      tmpDownText.replace(/\n$/, '')
      let tmpFormula = ''
      for (let i = 0; i < newVal.Formula.length; i++) {
        tmpFormula += newVal.Formula[i] + '\n'
      }
      tmpFormula.replace(/\n$/, '')

      this.inQueIdx = newVal.QueIdx
      this.inStdSec = newVal.StdSec
      this.inQueType = newVal.QuestionType
      this.inAnsType = newVal.AnswerType
      this.inComment = newVal.Tips
      this.inUpText = tmpUpText
      this.inDownText = tmpDownText
      this.inFormula = tmpFormula
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
        tmpOrgJson = [{}]
      }
      for (let i = 0; i < tmpOrgJson.length; i++) {
        tmpJson.push(JSON.parse(tmpOrgJson[i]))
      }
      this.inJsonText = tmpJson
      if (this.inAnsType === 'SC') {
        this.inAnswer = this.answerOptions[parseInt(newVal.Answers) - 1]
      } else if (this.inAnsType === 'IT') {
        this.inAnswer = newVal.Answers
      }
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
    ...mapActions('questions', ['addNewQuestion']),
    saveQuestion () {
      console.log('start save question by graphql-')
      let tmpOption = ''
      tmpOption += this.inOption1 + '||'
      tmpOption += this.inOption2 + '||'
      tmpOption += this.inOption3 + '||'
      tmpOption += this.inOption4 + '||'
      tmpOption += this.inOption5
      /* tmpOption.push(this.inOption1)
      tmpOption.push(this.inOption2)
      tmpOption.push(this.inOption3)
      tmpOption.push(this.inOption4) */
      let tmpCharts = ''
      let tmpShapes = ''
      let tmpTables = ''
      let tmpClocks = ''
      let tmpTags = ''
      if (this.inQueType === 'M_CHART') {
        for (let i = 0; i < this.inJsonText.length; i++) {
          tmpCharts += JSON.stringify(this.inJsonText[i]) + '||'
        }
      } else if (this.inQueType === 'M_SHAPE') {
        console.log('shapes- ', this.inJsonText)
        for (let i = 0; i < this.inJsonText.length; i++) {
          tmpShapes += JSON.stringify(this.inJsonText[i]) + '||'
        }
      } else if (this.inQueType === 'M_TABLE') {
        for (let i = 0; i < this.inJsonText.length; i++) {
          tmpTables += JSON.stringify(this.inJsonText[i]) + '||'
        }
      } else if (this.inQueType === 'M_CLK') {
        for (let i = 0; i < this.inJsonText.length; i++) {
          tmpClocks += JSON.stringify(this.inJsonText[i]) + '||'
        }
      }
      for (let i = 0; i < this.inTags.length; i++) {
        tmpTags += this.inTags[i] + '||'
      }
      console.log('tmpTags - ', tmpTags)
      // let tmpFormula = []
      // tmpFormula.push(this.inFormula)
      let tmpAnswer = this.answerOptions.indexOf(this.inAnswer) + 1

      this.$apollo
        .mutate({
          mutation: ADD_QUESTION_MUTATION,
          variables: {
            Gid: this.currentUser.Gid,
            QueIdx: parseInt(this.inQueIdx),
            Kp: this.kp,
            StdSec: this.inStdSec,
            AnswerType: this.inAnsType,
            QuestionType: this.inQueType,
            UpTexts: this.inUpText.replace(/\n$/, ''),
            DownTexts: this.inDownText.replace(/\n$/, ''),
            Formula: this.inFormula.replace(/\n$/, ''),
            Charts: tmpCharts,
            Shapes: tmpShapes,
            Tables: tmpTables,
            Clocks: tmpClocks,
            Options: tmpOption,
            Answers: tmpAnswer.toString(),
            Tags: tmpTags
          }
        })
        .then(response => {
          if (!response.errors) {
            if (response.data.AddQuestion) {
              console.log('Add question successful')
              this.updateCurrentQuestion(true)
            }
          } else {
            console.log('reponse error', response.errors)
          }
        })
        .catch(error => {
          console.log(error)
        })
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
    updateQuestionType: function () {
      // set json text example
      this.inJsonText = []
      if (this.inQueType === 'M_TABLE') {
        this.inJsonText.push({
          'tableTitle': 'Table Example',
          'style': 'width: 40%',
          'separator': 'cell',
          'columns': [
            { 'name': 'na', 'label': '', 'field': 'na' },
            { 'name': 'square', 'label': 'Square', 'field': 'square' },
            { 'name': 'circle', 'label': 'Circle', 'field': 'circle' }
          ],
          'tableData': [
            { 'na': 'Black', 'square': '4', 'circle': '6' },
            { 'na': 'White', 'square': '7', 'circle': '5' }
          ]
        })
      } else if (this.inQueType === 'M_CHART') {
        this.inJsonText.push({
          'xAxis': {
            'type': 'value',
            'min': 0,
            'max': 5,
            'splitNumber': 5,
            'name': 'Cups of flour',
            'nameLocation': 'middle',
            'nameGap': 40
          },
          'yAxis': {
            'type': 'value',
            'min': 0,
            'max': 800,
            'splitNumber': 4,
            'name': 'Weight of flour (g)',
            'nameLocation': 'middle',
            'nameGap': 40
          },
          'series': {
            'type': 'line',
            'data': [ [0, 0], [5, 800] ]
          }
        })
      } else if (this.inQueType === 'M_CLK') {
        this.inJsonText[0] = [{
          'type': 'Not Available'
        }]
      } else if (this.inQueType === 'M_SHAPE') {
        this.inJsonText.push({
          'type': 'line',
          'config': {
            'points': [350, 50, 330, 200, 520, 150, 350, 50],
            'stroke': 'black',
            'strokeWidth': 2
          }
        })
        this.inJsonText.push({
          'type': 'circle',
          'config': { 'x': 100, 'y': 30, 'radius': 15, 'stroke': 'black' }
        })
      } else if (this.inQueType === 'M_COM') {
        this.inJsonText = []
      }
      this.updateCurrentQuestion(false)
    },
    updateCurrentQuestion: function (opt) {
      // save in current question, this should be done in editform
      console.log('write new data')
      let tmp = []
      tmp.push(this.inOption1)
      tmp.push(this.inOption2)
      tmp.push(this.inOption3)
      tmp.push(this.inOption4)
      tmp.push(this.inOption5)
      let tmpAnsInt = this.answerOptions.indexOf(this.inAnswer) + 1
      let tmp1 = [tmpAnsInt.toString()]
      let tmpCharts = []
      let tmpClk = []
      let tmpTables = []
      let tmpShapes = []
      let tmpUpText = []
      let tmpStr = this.inUpText.replace(/\n$/, '')
      if (tmpStr !== '') {
        tmpUpText = tmpStr.split('\n')
      }
      let tmpDownText = []
      tmpStr = this.inDownText.replace(/\n$/, '')
      if (tmpStr !== '') {
        tmpDownText = tmpStr.split('\n')
      }
      tmpStr = this.inFormula.replace(/\n$/, '')
      let tmpFormula = tmpStr.split('\n')
      console.log('tmpFormula - ', tmpFormula)
      if (this.inQueType === 'M_TABLE') {
        for (let i = 0; i < this.inJsonText.length; i++) {
          tmpTables.push(JSON.stringify(this.inJsonText[i]))
        }
      } else if (this.inQueType === 'M_CHART') {
        for (let i = 0; i < this.inJsonText.length; i++) {
          tmpCharts.push(JSON.stringify(this.inJsonText[i]))
        }
      } else if (this.inQueType === 'M_CLK') {
        for (let i = 0; i < this.inJsonText.length; i++) {
          tmpClk.push(JSON.stringify(this.inJsonText[i]))
        }
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
          Formula: tmpFormula,
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
      if (opt === true) {
        this.addNewQuestion(
          {
            Kp: this.QuizId,
            QueIdx: this.inQueIdx,
            StdSec: this.inStdSec,
            QuestionType: this.inQueType,
            AnswerType: this.inAnsType,
            UpTexts: tmpUpText,
            DownTexts: tmpDownText,
            Formula: tmpFormula,
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
      }
    },
    clearForm () {
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
