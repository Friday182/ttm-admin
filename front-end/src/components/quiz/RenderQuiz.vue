<template>
  <q-card
    class="bg-gray text-black"
    align="center"
  >
    <q-card-section>
      Showing the Question: {{ editQuestion.QueIdx }}
    </q-card-section>
    <q-card-section
      v-if="editQuestion.QuestionType === 'M_COM'"
      key="que_type"
    >
      <question-math-common />
    </q-card-section>
    <q-card-section
      v-if="editQuestion.QuestionType === 'M_CLK'"
      key="que_type"
    >
      <question-math-clock />
    </q-card-section>
    <q-card-section
      v-if="editQuestion.QuestionType === 'M_MONEY'"
    >
      <question-math-money />
    </q-card-section>
    <q-card-section
      v-if="editQuestion.QuestionType === 'E_COM'"
    >
      <question-en-common />
    </q-card-section>
    <q-card-section
      v-if="editQuestion.QuestionType === 'M_CHART'"
    >
      <question-math-chart />
    </q-card-section>
    <q-card-section
      v-if="editQuestion.QuestionType === 'M_TABLE'"
    >
      <question-math-table />
    </q-card-section>
    <q-card-section
      v-if="editQuestion.QuestionType === 'M_SHAPE'"
    >
      <question-math-shape />
    </q-card-section>
    <q-card-section
      v-if="editQuestion.AnswerType === 'SC'"
      class="q-mt-lg"
    >
      <answer-sc
        :useradio="false"
      />
    </q-card-section>
    <q-card-section
      v-else-if="editQuestion.AnswerType === 'IT'"
      class="q-mt-lg"
    >
      <answer-it />
    </q-card-section>
    <q-card-section
      v-else-if="editQuestion.AnswerType === 'IT_FRAC'"
      key="answer-type"
      class="q-mt-lg"
    >
      <answer-it-frac />
    </q-card-section>
    <q-card-section
      v-else-if="editQuestion.AnswerType === 'TF'"
      key="answer-type"
      class="q-mt-xl"
    >
      <answer-tf />
    </q-card-section>
    <q-card-section
      v-else
      class="q-mt-xl"
    >
      <q-spinner-hourglass
        color="blue"
        size="4em"
      />
    </q-card-section>
  </q-card>
</template>
<script>
import { mapGetters } from 'vuex'

export default {
  components: {
    'question-math-common': require('components/quiz/QuestionMathCommon.vue').default,
    'question-math-clock': require('components/quiz/QuestionMathClock.vue').default,
    // 'question-math-clock': require('components/quiz/QuestionMathClock1.vue').default,
    'question-math-money': require('components/quiz/QuestionMathMoney.vue').default,
    'question-math-chart': require('components/quiz/QuestionMathChart.vue').default,
    'question-math-table': require('components/quiz/QuestionMathTable.vue').default,
    'question-math-shape': require('components/quiz/QuestionMathShape.vue').default,
    'answer-sc': require('components/quiz/AnswerSc.vue').default,
    'answer-it': require('components/quiz/AnswerIt.vue').default,
    'answer-it-frac': require('components/quiz/AnswerItFrac.vue').default,
    'answer-tf': require('components/quiz/AnswerTf.vue').default,
    'question-en-common': require('components/quiz/QuestionEnCommon.vue').default
  },
  props: {
  },
  data () {
    return {
      numQue: 0
    }
  },
  computed: {
    ...mapGetters('currentQuestion', ['currentQuestion']),
    editQuestion: function () {
      return this.currentQuestion
    }
  },
  methods: {
  }
}
</script>

<style lang="stylus" scoped>
</style>
