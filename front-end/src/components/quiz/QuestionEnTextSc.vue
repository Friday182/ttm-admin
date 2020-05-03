<template>
  <div>
    <q-card
      class="bg-gray text-blue q-px-lg"
      no-caps
      style="width: 70%; font-size: 16pt"
    >
      <q-card-section align="left">
        <div
          v-for="(item,key) in upTexts"
          :key="key"
        >
          {{ item }}
        </div>
      </q-card-section>
      <q-card-section align="left">
        <div
          v-for="(item,idx) in downTexts"
          :key="idx"
        >
          <q-btn
            v-for="(word,index) in item"
            :key="idx * 10 + index"
            :color="getColor(idx, index)"
            :disable="noEdit"
            size="lg"
            no-caps
            flat
            push
            dense
            @click="setChoice(idx, index)"
          >
            {{ word }}
          </q-btn>
        </div>
      </q-card-section>
      <q-card-section>
        <q-btn
          glossy
          no-caps
          size="lg"
          label="Submit"
          @click="submitAnswer"
        />
      </q-card-section>
    </q-card>
  </div>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'

export default {
  name: 'QuestionEnTextSc',
  components: {
  },
  data () {
    return {
      btnColor: [],
      noEdit: false
    }
  },
  computed: {
    ...mapGetters('questions', ['getQuestions']),
    ...mapGetters('currentQuestion', ['currentQuestion']),
    upTexts: function () {
      return (this.currentQuestion) ? this.currentQuestion.UpTexts : ''
    },
    downTexts: function () {
      // return [['this', 'is', 'a', 'good', 'story', 'to', 'read', 'dad', 'said.', 'this', 'is', 'a', 'good', 'story', 'to', 'read', 'dad', 'said.', 'this', 'is', 'a', 'good', 'story', 'to', 'read', 'dad', 'said.'], ['Yes, I', 'think so', ',"said Jamie.']]
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.DownTexts.length; i++) {
          let str = this.currentQuestion.DownTexts[i].split(' ')
          tmp.push(str)
        }
      }
      return tmp
      // return (this.currentQuestion) ? this.currentQuestion.DownTexts : ''
    }
  },
  watch: {
    que: function (newVal, oldVal) {
      console.log('downtext was changed')
      for (let i = 0; i < newVal.length; i++) {
        for (let j = 0; j < newVal[i].length; j++) {
          this.btnColor.push('info')
        }
      }
    }
  },
  mounted () {
    for (let i = 0; i < this.downTexts.length; i++) {
      for (let j = 0; j < this.downTexts[i].length; j++) {
        this.btnColor.push('primary')
      }
    }
  },
  methods: {
    setChoice: function (idx, index) {
      console.log('select idx: ', idx, index)
      let tmpIdx = 0
      for (let i = 0; i < idx; i++) {
        tmpIdx += this.downTexts[i].length
      }
      tmpIdx += index
      if (this.btnColor[tmpIdx] === 'red') {
        this.btnColor[tmpIdx] = 'primary'
      } else {
        this.btnColor[tmpIdx] = 'red'
      }
      this.btnColor.push('')
      this.btnColor.splice(this.btnColor.length, 1)
      console.log('button color: ', tmpIdx, this.btnColor[tmpIdx])
    },
    getColor: function (idx, index) {
      let tmpIdx = 0
      for (let i = 0; i < idx; i++) {
        tmpIdx += this.downTexts[i].length
      }
      tmpIdx += index
      return this.btnColor[tmpIdx]
    },
    submitAnswer: function () {
      this.noEdit = true
    }
  }
}
</script>

<style>
</style>
