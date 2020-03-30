<template>
  <div>
    <q-input
      ref="input1"
      v-model="ansText[0]"
      :label="answer1_text"
      outlined
      autofocus
      style="width: 300px; font-size: 18px"
    />
    <q-input
      v-if="answer2_text !== ''"
      ref="input2"
      v-model="ansText[1]"
      :label="answer2_text"
      outlined
      style="width: 300px; font-size: 18px"
      class="q-mt-md"
    />
    <q-btn
      label="Done"
      outlined
      color="primary"
      class="q-mt-md q-mb-xl"
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
      return (this.currentQuestion) ? this.currentQuestion.Options[0] : ''
    },
    answer2_text: function () {
      if (this.currentQuestion.Options[1]) {
        return this.currentQuestion.Options[1]
      } else {
        return ''
      }
    }
  },
  methods: {
  }
}
</script>

<style>
</style>
