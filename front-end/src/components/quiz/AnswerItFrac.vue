<template>
  <div>
    <div class="row">
      <div class="col-6" />
      <div class="col-1 q-ml-md">
        <q-input
          ref="input1"
          v-model="ansText[1]"
          label="Numerator"
          outlined
          dense
          style="width: 100px; font-size: 18px"
          @keyup.enter="selected()"
        />
      </div>
    </div>
    <div class="row">
      <div class="col-5" />
      <div class="col-1">
        <q-input
          ref="input0"
          v-model="ansText[0]"
          label="Whole Number"
          outlined
          autofocus
          dense
          style="width: 100px; font-size: 18px"
          @keyup.enter="selected()"
        />
      </div>
      <div
        class="col-1 q-ml-sm q-mt-xs"
        style="width: 130px;color:black;"
      >
        _____________________
      </div>
    </div>
    <div class="row">
      <div class="col-6" />
      <div class="col-1 q-ml-md">
        <q-input
          ref="input2"
          v-model="ansText[2]"
          label="Denominator"
          outlined
          dense
          style="width: 100px; font-size: 18px"
          @keyup.enter="selected()"
        />
      </div>
    </div>
    <q-btn
      label="Next"
      outlined
      color="primary"
      class="q-my-xl"
      @click="selectedDone()"
    />
  </div>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'

export default {
  name: 'InputTextAnswer',
  props: {
    index: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      ansCorrect: false,
      ansText: ['', '', '', '']
    }
  },
  computed: {
    ...mapGetters('currentQuestion', ['currentQuestion']),
    answer1_text: function () {
      return (this.allQuestions[this.index]) ? this.allQuestions[this.index].Options[0] : ''
    },
    answer2_text: function () {
      if (this.allQuestions[this.index].Options[1]) {
        return this.allQuestions[this.index].Options[1]
      } else {
        return ''
      }
    }
  },
  methods: {
    selected () {
      if (this.ansText[1] === '') {
        this.$refs.input1.focus()
      } else if (this.ansText[2] === '') {
        this.$refs.input2.focus()
      } else if (this.ansText[0] === '') {
        this.$refs.input0.focus()
      }
    },
    selectedDone () {
      console.log('answer - ', this.index, this.ansText[0], this.allQuestions[this.index].Answers[0])
      if (this.allQuestions[this.index].Answers[0]) {
        if (this.ansText[0] !== '') {
          if (this.ansText[0].trim().toUpperCase() === this.allQuestions[this.index].Answers[0].toUpperCase()) {
            this.ansCorrect = true
          } else {
            this.ansCorrect = false
          }
        } else {
          this.ansCorrect = false
        }
      }
      for (let i = 1; i < 3; i++) {
        if (this.ansText[i] !== '' && this.allQuestions[this.index].Answers[i]) {
          if (this.ansText[i].trim().toUpperCase() === this.allQuestions[this.index].Answers[i].toUpperCase()) {
            this.ansCorrect = true
          } else {
            this.ansCorrect = false
            break
          }
        }
      }
      this.$emit('ansCorrect', this.ansCorrect)
      this.$refs.input1.focus()
      this.ansText = ['', '', '', '']
    }
  }
}
</script>

<style>
</style>
