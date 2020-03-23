<template>
  <div>
    <q-card
      v-if="isLearn"
      style="width: 50%"
      no-caps
    >
      <q-card-section>
        Look and Learn:
      </q-card-section>
      <q-card-section
        style="width: 70%"
      >
        <q-table
          v-if="wordType==='tense'"
          class="q-table"
          :data="tableData"
          :columns="columnTense"
          row-key="name"
          separator="cell"
          flat
          bordered
          hide-bottom
        >
          <q-td
            slot="body-cell-col1"
            slot-scope="props"
            :props="props"
            style="color: blue; font-size: 15pt"
          >
            {{ learnWords[0] }}
          </q-td>
          <q-td
            slot="body-cell-col2"
            slot-scope="props"
            :props="props"
            style="color: blue; font-size: 15pt"
          >
            {{ learnWords[1] }}
          </q-td>
          <q-td
            slot="body-cell-col3"
            slot-scope="props"
            :props="props"
            style="color: blue; font-size: 15pt"
          >
            {{ learnWords[2] }}
          </q-td>
        </q-table>
        <q-table
          v-else-if="wordType==='plural'"
          class="q-table"
          :data="tableData"
          :columns="columnPlural"
          row-key="name"
          separator="cell"
          flat
          bordered
          hide-bottom
        >
          <q-td
            slot="body-cell-col1"
            slot-scope="props"
            :props="props"
            style="color: blue; font-size: 15pt"
          >
            {{ learnWords[0] }}
          </q-td>
          <q-td
            slot="body-cell-col2"
            slot-scope="props"
            :props="props"
            style="color: blue; font-size: 15pt"
          >
            {{ learnWords[1] }}
          </q-td>
        </q-table>
        <q-table
          v-else
          class="q-table"
          :data="tableData"
          :columns="columnNewWords"
          row-key="name"
          separator="cell"
          flat
          bordered
          hide-bottom
        >
          <q-td
            slot="body-cell-col1"
            slot-scope="props"
            :props="props"
            style="color: blue; font-size: 15pt"
          >
            {{ learnWords[0] }}
          </q-td>
        </q-table>
      </q-card-section>
      <q-card-section
        class="q-mt-md"
      >
        <q-btn
          style="width: 10%"
          no-caps
          color="primary"
          label="Ok"
          @click="showQuestion"
        />
      </q-card-section>
    </q-card>
    <q-card
      v-else
      class="q-mt-md"
      style="width: 50%"
      no-caps
    >
      <q-card-section
        v-if="wordType === 'normal'"
      >
        Rewrite the Word please.
      </q-card-section>
      <q-card-section
        v-else
      >
        Given the word:
        <q-btn
          style="width: 20%"
          no-caps
          flat
          text-color="blue"
          size="lg"
        >
          {{ formula }}
        </q-btn>
      </q-card-section>
    </q-card>
    <alert
      :active="isUserManual"
      :msg="alertMsg"
      @cancelAlert="alertClose"
    />
  </div>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'

export default {
  name: 'QuestionEnWordLearn',
  components: {
    'alert': require('components/common/AlertMsg.vue').default
  },
  props: {
    index: {
      type: Number,
      default: 0
    },
    trigger: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      columnTense: [
        { name: 'col1', align: 'center', label: 'Verb' },
        { name: 'col2', align: 'center', label: 'Simple Past' },
        { name: 'col3', align: 'center', label: 'Past Participle' }
      ],
      columnPlural: [
        { name: 'col1', align: 'center', label: 'Word' },
        { name: 'col2', align: 'center', label: 'Plural' }
      ],
      columnNewWords: [
        { name: 'col1', align: 'center', label: 'Word' }
      ],
      tableData: [],
      isLearn: true,
      isTense: true,
      isPlural: false,
      interHandler: 0,
      timeCount: 0,
      alertMsg: 'User Instruction: Read and remember the words, question will be shown after 10 seconds.Or, press the "Show Question" button to continue.',
      isUserManual: false
    }
  },
  computed: {
    ...mapGetters('questions', ['getQuestions']),
    formula: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].UpTexts[0] : ''
    },
    learnWords: function () {
      return (this.getQuestions[this.index]) ? JSON.parse(this.getQuestions[this.index].Tag[0]) : ''
    },
    questionType: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].QuestionType : ''
    },
    wordType: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].Tag[1] : ''
    }
  },
  watch: {
    trigger: function (newVal, oldVal) {
      console.log('En WL question triggered')
      if (this.questionType === 'E_WL') {
        console.log('new words - ', this.learnWords)
        this.tableData = [{
          col1: this.learnWords[0],
          col2: this.learnWords[1],
          col3: this.learnWords[2],
          col4: this.learnWords[3]
        }]
        this.isLearn = true
        this.timeCount = 0
      } else if (this.questionType === 'E_WT') {
        console.log('enter E_WT')
        this.showQuestion()
      } else {
        console.log('something wrong')
      }
    }
  },
  created () {
    document.addEventListener('keyup', this.handleKeyup)
    console.log('create word learn')
    if (this.questionType === 'E_WL') {
      this.isLearn = true
      this.isUserManual = true
      this.tableData = [{
        col1: this.learnWords[0],
        col2: this.learnWords[1],
        col3: this.learnWords[2],
        col4: this.learnWords[3]
      }]
    } else if (this.questionType === 'E_WT') {
      this.isLearn = false
      this.isUserManual = false
    } else {
      console.log('wrong questiontype?')
    }
    this.interHandler = setInterval(this.updateTimeCount, 500)
    console.log('new words - ', this.learnWords)
  },
  destroyed () {
    document.removeEventListener('keyup', this.handleKeyup)
    clearInterval(this.interHandler)
  },
  methods: {
    updateTimeCount () {
      this.timeCount++
      if (this.timeCount === 20) {
        this.showQuestion()
      }
    },
    showQuestion () {
      this.$emit('evShowQuestion')
      this.isLearn = false
      console.log('clear timeout and interval timer is - ', this.timeCount)
    },
    handleKeyup (event) {
      let self = this
      if (event.keyCode === 13 && self.timeCount > 1) {
        console.log('got event key - ', event.keyCode)
        self.showQuestion()
      }
    },
    alertClose () {
      this.isUserManual = false
      this.timeCount = 0
    }
  }
}
</script>

<style>
.q-table {
  font-size: 24px;
}
</style>
