<template>
  <div>
    <q-input
      ref="input1"
      v-model="ansText[0]"
      :label="answer1_text"
      outlined
      autofocus
      style="width: 300px; font-size: 18px"
      @keyup.enter="selected()"
    />
    <q-input
      v-if="answer2_text !== ''"
      ref="input2"
      v-model="ansText[1]"
      :label="answer2_text"
      outlined
      style="width: 300px; font-size: 18px"
      class="q-mt-md"
      @keyup.enter="selected()"
    />
    <q-btn
      label="Done"
      outlined
      color="primary"
      class="q-mt-md q-mb-xl"
      @click="selected()"
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
    ...mapGetters('questions', ['getQuestions']),
    answer1_text: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].Options[0] : ''
    },
    answer2_text: function () {
      if (this.getQuestions[this.index].Options[1]) {
        return this.getQuestions[this.index].Options[1]
      } else {
        return ''
      }
    }
  },
  methods: {
    selected () {
      let selectDone = false
      if (this.ansText[0] === '') {
        this.$refs.input1.focus()
      } else if (this.getQuestions[this.index].Options[1]) {
        console.log('option1 = ', this.getQuestions[this.index].Options[1])
        if (this.ansText[1] === '') {
          this.$refs.input2.focus()
        } else {
          selectDone = true
        }
      } else {
        // only one input and already done
        selectDone = true
      }

      if (selectDone === true) {
        console.log('answer - ', this.index, this.ansText[0], this.getQuestions[this.index].Answers[0])
        for (let i = 0; i < 4; i++) {
          if (this.ansText[i] !== '' && this.getQuestions[this.index].Answers[i]) {
            let tmp = parseFloat(this.ansText[i])
            let tmp1 = tmp.toString()
            if (tmp1.trim().toUpperCase() === this.getQuestions[this.index].Answers[i].toUpperCase()) {
              this.ansCorrect = true
            } else {
              this.ansCorrect = false
              break
            }
          }
        }
        /*
        if (this.ansText[0].trim().toUpperCase() === this.getQuestions[this.index].Answers[0].toUpperCase()) {
          if (this.ansText[1] !== '' && this.getQuestions[this.index].Answers[1]) {
            if (this.ansText[1].trim().toUpperCase() === this.getQuestions[this.index].Answers[1].toUpperCase()) {
              this.ansCorrect = true
            } else {
              this.ansCorrect = false
            }
          } else {
            this.ansCorrect = true
          }
        } else {
          this.ansCorrect = false
        }
        */
        this.$emit('ansCorrect', this.ansCorrect)
        this.$refs.input1.focus()
        this.ansText[0] = ''
        this.ansText[1] = ''
      }
    }
  }
}
</script>

<style>
</style>
