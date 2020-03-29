<template>
  <div>
    <div
      class="row"
    >
      <q-input
        v-model="text1"
        label="I"
        style="width: 20px"
      />
      <q-input
        v-model="text2"
        label="has"
        style="width: 40px"
      />
      <q-input
        v-model="text4"
        label="an"
        style="width: 30px"
      />
      <q-input
        v-model="text3"
        label="dog."
        style="width: 40px; font-size: 10pt"
      />
    </div>
  </div>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'

export default {
  name: 'QuestionEnCommon',
  components: {
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
      tableData: [],
      isLearn: true
    }
  },
  computed: {
    ...mapGetters('questions', ['allQuestions']),
    formula: function () {
      return (this.allQuestions[this.index]) ? this.allQuestions[this.index].UpTexts[0] : ''
    },
    wordList: function () {
      if (this.allQuestions[this.index]) {
        let tmp = JSON.parse(this.allQuestions[this.index].Tag[0])
        return tmp
      }
      return ''
    }
  },
  watch: {
    trigger: function (newVal, oldVal) {
      console.log('En WL question triggered')
      if (this.questionType === 'E_WL') {
        console.log('new words - ', this.learnWords)
      } else if (this.questionType === 'E_WT') {
        console.log('enter E_WT')
      } else {
        console.log('something wrong')
      }
    }
  },
  created () {
    document.addEventListener('keyup', this.handleKeyup)
    console.log('create word learn')
    this.interHandler = setInterval(this.updateTimeCount, 500)
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
      console.log('clear timeout and interval timer is - ', this.timeCount)
    },
    handleKeyup (event) {
      let self = this
      if (event.keyCode === 13 && self.timeCount > 1) {
        console.log('got event key - ', event.keyCode)
        self.showQuestion()
      }
    }
  }
}
</script>

<style>
.q-table {
  font-size: 24px;
}
</style>
