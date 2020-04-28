<template>
  <div>
    <q-card>
      <q-card-section class="q-ml-md">
        <div class="row">
          <q-btn
            label="Copy"
            no-caps
            color="primary"
            style="font: 120% Arial bold"
            @click="copyQuestion()"
          />
          <q-btn
            class="q-ml-md"
            label="Paste"
            no-caps
            color="primary"
            :disable="tmpQue"
            style="font: 120% Arial bold"
            @click="pasteQuestion()"
          />
          <q-btn
            class="q-ml-md"
            label="Clear"
            no-caps
            color="primary"
            :disable="tmpQue"
            style="font: 120% Arial bold"
            @click="clearForm()"
          />
          <q-space />
          <q-btn
            label="Save Question"
            no-caps
            color="primary"
            style="font: 120% Arial bold"
            @click="saveQuestion()"
          />
        </div>
      </q-card-section>
      <q-separator />
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
              v-model="inType"
              outlined
              dense
              label="Types"
              :options="typeOptions"
              @input="updateQuestionType"
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
        <div class="row bg-cyan-1">
          <div class="col-2">
            <q-btn
              color="primary"
              label="Select KPs"
              dense
              no-caps
            >
              <q-menu fit anchor="top right" self="bottom middle">
                <q-list dense style="min-width: 100px">
                  <q-item
                    v-for="(item, key) in kpOptions"
                    :key="key"
                    clickable
                    @mouseover.native="setOver(item.value, true)"
                    @mouseout.native="setOver(item.value, false)"
                  >
                    <q-item-section class="text-deep-orange"> {{ item.label }} </q-item-section>
                    <q-item-section side>
                      <q-icon name="keyboard_arrow_right" />
                    </q-item-section>
                    <q-menu
                      v-model="item.active"
                      anchor="top right"
                      self="top left"
                    >
                      <q-list>
                        <q-item
                          v-for="(child, key) in item.children"
                          :key="key"
                          dense
                          clickable
                          v-close-popup
                          @click="addInTags(child.value)"
                        >
                          <q-item-section class="text-blue">{{ child.label }}</q-item-section>
                        </q-item>
                      </q-list>
                    </q-menu>
                  </q-item>
                </q-list>
              </q-menu>
            </q-btn>
          </div>
          <div class="col-5">
            <q-chip
              v-for="item in inTags"
              :key="item"
              removable
              color="primary"
              text-color="white"
              :label="item"
              @remove="removeTag(item)"
            />
          </div>
        </div>
      </q-card-section>
      <q-separator />
      <q-card-section class="q-ml-md">
        <div class="row">
          <q-btn
            label="Copy"
            no-caps
            color="primary"
            style="font: 120% Arial bold"
            @click="copyQuestion()"
          />
          <q-btn
            class="q-ml-md"
            label="Paste"
            no-caps
            color="primary"
            :disable="tmpQue"
            style="font: 120% Arial bold"
            @click="pasteQuestion()"
          />
          <q-btn
            class="q-ml-md"
            label="Clear"
            no-caps
            color="primary"
            :disable="tmpQue"
            style="font: 120% Arial bold"
            @click="clearForm()"
          />
          <q-space />
          <q-btn
            label="Save Question"
            no-caps
            color="primary"
            style="font: 120% Arial bold"
            @click="saveQuestion()"
          />
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<script>
import VJsoneditor from 'v-jsoneditor'
import { ADD_QUESTION_MUTATION } from '../../graphql/mutations'
import { mapMutations, mapGetters, mapActions } from 'vuex'
import { debounce } from 'quasar'

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
      inType: 'Line Chart',
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
      inTags: [],
      tmpQue: true,
      usedSeconds: 0,
      vjsonOptions: {
      },
      skipQueryAddQuiz: true,
      queTypeOptions: ['M_COM', 'M_TABLE', 'M_SHAPE', 'M_CHART'],
      ansTypeOptions: ['SC', 'IT', 'TF', 'MC'],
      answerOptions: ['Option A', 'Option B', 'Option C', 'Option D', 'Option E'],
      typeOptions: ['Bar Chart', 'Pie Chart', 'Line Chart', 'Coordinate Chart'],
      elprops: {
        multiple: true,
        expandTrigger: 'hover'
      },
      kpOptions: [
        {
          value: '1',
          label: 'Basic Calculation',
          over: false,
          active: false,
          children: [
            { value: 'MA1', label: 'Plus' },
            { value: 'MA2', label: 'Minus' },
            { value: 'MA3', label: 'Multiple' },
            { value: 'MA4', label: 'Division' }
          ]
        },
        {
          value: '2',
          label: 'DataHandling',
          over: false,
          active: false,
          children: [
            { value: 'MF1', label: 'Bar Chart' },
            { value: 'MF2', label: 'Line Chart' },
            { value: 'MF3', label: 'Table' }
          ]
        }
      ],
      kpOptionsOld: [
        { label: 'DataHandling', value: 'MF6' },
        { label: 'Algebra', value: 'MC31' },
        { label: 'TimesTable', value: 'MA24' },
        { label: 'Multiplication', value: 'MA39' },
        { label: 'Division', value: 'MA51' },
        { label: 'FourOperator', value: 'MA54' },
        { label: 'MinusNumber', value: 'MA61' },
        { label: 'Fraction', value: 'MB22' },
        { label: 'Decimal', value: 'MB37' },
        { label: 'Ratio&Percentage', value: 'MB43' },
        { label: 'Measurement', value: 'MD20' },
        { label: 'Shapes', value: 'ME4' },
        { label: 'Rounding', value: 'MH6' },
        { label: 'Average&Mean', value: 'MH12' },
        { label: 'Time&Date', value: 'MI11' }
      ],
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
      } else if (newVal.QuestionType === 'M_CLK') {
        tmpOrgJson = newVal.Clocks
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
            this.alertError('Question Save Failed ! ! !')
          }
        })
        .catch(error => {
          console.log(error)
          this.alertError('Question Save Failed ! ! !')
        })
    },
    alertError (msg) {
      const dialog = this.$q.dialog({
        title: 'Error',
        message: msg,
        html: false,
        ok: {
          push: true
        }
      }).onDismiss(() => {
        clearTimeout(timer)
      })
      const timer = setTimeout(() => {
        dialog.hide()
      }, 3000)
    },
    onJsonChange (value) {
      console.log('value:', value)
    },
    addInTags: function (tag) {
      console.log('added tag: ', tag)
      if (this.inTags.length < 3) {
        let isNew = true
        for (let i = 0; i < this.inTags.length; i++) {
          if (this.inTags[i] === tag) {
            isNew = false
            break
          }
        }
        if (isNew === true) {
          this.inTags.push(tag)
        }
      } else {
        this.alertError('Maximum 3 tags!')
      }
    },
    removeTag: function (tag) {
      console.log('remove tag:', tag)
      for (let i = 0; i < this.inTags.length; i++) {
        if (this.inTags[i] === tag) {
          this.inTags.splice(i, 1)
        }
      }
    },
    checkOver: function (idx, isHover) {
      let tmp = parseInt(idx) - 1
      console.log('idx to tmp:', tmp)
      if (this.kpOptions[tmp].over) {
        this.kpOptions[tmp].active = true
      }
    },
    setOver (idx, isHover) {
      let tmp = parseInt(idx) - 1
      this.kpOptions[tmp].over = isHover
      this.debounceFunc(idx, isHover)
    },
    debounceFunc: debounce(function (idx, isHover) { this.checkOver(idx, isHover) }, 300),
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
            { 'name': 'col1', 'label': '', 'field': 'col1' },
            { 'name': 'col2', 'label': 'col2', 'field': 'col2' },
            { 'name': 'col3', 'label': 'col3', 'field': 'col3' }
          ],
          'tableData': [
            { 'col1': 'Black', 'col2': '4', 'col3': '6' },
            { 'col1': 'White', 'col2': '7', 'col3': '5' }
          ]
        })
      } else if (this.inQueType === 'M_CHART') {
        if (this.inType === 'Line Chart') {
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
        } else if (this.inType === 'Bar Chart') {
          this.inJsonText.push({
            'xAxis': {
              'type': 'category',
              'min': 0,
              'max': 5,
              'splitNumber': 5,
              'name': 'Sports',
              'nameLocation': 'middle',
              'nameGap': 40,
              'data': ['Mon', 'Tue', 'Wed', 'Thu', 'Fri', 'Sat', 'Sun']
            },
            'yAxis': {
              'type': 'value',
              'min': 0,
              'max': 400,
              'splitNumber': 4,
              'name': 'Numbers',
              'nameLocation': 'middle',
              'nameGap': 40
            },
            'series': {
              'type': 'bar',
              'barWidth': 40,
              'data': [120, 200, 150, 80, 70, 110, 130]
            }
          })
        } else if (this.inType === 'Pie Chart') {
          this.inJsonText.push({
            'legend': {
              'orient': 'vertical',
              'left': 10,
              'data': ['6 marks', '7 marks', '8 marks', '9 marks', '10 marks']
            },
            'series': [{
              'type': 'pie',
              'radius': '100%',
              'center': ['50%', '50%'],
              'data': [
                { 'name': '6 marks', 'value': 8 },
                { 'name': '7 marks', 'value': 10 },
                { 'name': '8 marks', 'value': 8 },
                { 'name': '9 marks', 'value': 8 },
                { 'name': '10 marks', 'value': 8 }
              ],
              'animation': false,
              'label': {
                'show': true,
                'color': 'black',
                'position': 'inside',
                'bleedMargin': 5
              },
              'left': '20%',
              'right': '20%',
              'top': '20%',
              'bottom': '20%'
            }]
          })
        } else if (this.inType === 'Coordinate Chart') {
          this.inJsonText.push({
            'grid': {
              'bottom': 10,
              'right': 10,
              'containLabel': true
            },
            'xAxis': {
              'type': 'value',
              'min': -6,
              'max': 5,
              'splitNumber': 11,
              'splitLine': {
                'show': true,
                'lineStyle': {
                  'color': '#999',
                  'type': 'dashed'
                }
              },
              'axisTick': {
                'show': true,
                'length': 10
              },
              'axisLine': {
              }
            },
            'yAxis': {
              'type': 'value',
              'min': -3,
              'max': 4,
              'splitNumber': 7,
              'splitLine': {
                'show': true,
                'lineStyle': {
                  'color': '#999',
                  'type': 'dashed'
                }
              },
              'axisTick': {
                'show': true,
                'length': 10
              },
              'axisLine': {
                'show': true
              }
            },
            'series': {
              'type': 'scatter',
              'data': [ [ -3, 3 ] ],
              'symbolSize': 15
            }
          })
        }
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
        this.inJsonText.push({
          'type': 'text',
          'config': { 'x': 420, 'y': 110, 'text': 'ABC', 'fontSize': 15, 'fill': 'blue' }
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
    copyQuestion () {
      localStorage.setItem('tmpQue', JSON.stringify(this.currentQuestion))
      this.tmpQue = false
    },
    pasteQuestion () {
      if (localStorage.getItem('tmpQue')) {
        let que = JSON.parse(localStorage.getItem('tmpQue'))
        this.doSetCurrentQuestion(
          {
            Kp: this.QuizId,
            QueIdx: this.inQueIdx,
            StdSec: que.StdSec,
            QuestionType: que.QuestionType,
            AnswerType: que.AnswerType,
            UpTexts: que.UpTexts,
            DownTexts: que.DownTexts,
            Formula: que.Formula,
            Options: que.Options,
            Tags: que.Tags,
            Answers: que.Answers,
            Charts: que.Charts,
            Clocks: que.Clocks,
            Tables: que.Tables,
            Shapes: que.Shapes,
            AnswerText: '',
            Helper: false,
            Imgs: [],
            Tips: [],
            Choice: ''
          }
        )
      }
    },
    clearForm () {
      this.doSetCurrentQuestion(
        {
          Kp: this.QuizId,
          QueIdx: this.inQueIdx,
          StdSec: 30,
          QuestionType: 'M_COM',
          AnswerType: 'SC',
          UpTexts: [],
          DownTexts: [],
          Formula: [],
          Options: [],
          Tags: [],
          Answers: [],
          Charts: [],
          Clocks: [],
          Tables: [],
          Shapes: [],
          AnswerText: '',
          Helper: false,
          Imgs: [],
          Tips: [],
          Choice: ''
        }
      )
      localStorage.removeItem('tmpQue')
      this.tmpQue = true
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
