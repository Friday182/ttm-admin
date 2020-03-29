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
      <q-card-section>
        <div
          v-for="(item,key) in formula"
          :key="key"
          align="left"
        >
          <vue-mathjax
            :formula="item"
            :options="{}"
          />
        </div>
      </q-card-section>
      <q-card-section align="left">
        <div
          v-for="(item,key) in downTexts"
          :key="key"
        >
          {{ item }}
        </div>
      </q-card-section>
    </q-card>
  </div>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'
import { VueMathjax } from 'vue-mathjax'

export default {
  name: 'QuestionMathCommon',
  components: {
    'vue-mathjax': VueMathjax
  },
  props: {
    index: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
    }
  },
  computed: {
    ...mapGetters('currentQuestion', ['currentQuestion']),
    upTexts: function () {
      return (this.currentQuestion) ? this.currentQuestion.UpTexts : ''
    },
    downTexts: function () {
      return (this.currentQuestion) ? this.currentQuestion.DownTexts : ''
    },
    formula: function () {
      return (this.currentQuestion) ? this.currentQuestion.Formula : ''
    }
  },
  methods: {
  }
}
</script>

<style>
</style>
