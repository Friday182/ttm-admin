<template>
  <div>
    <div
      :class="layoutFormat1"
    >
      <div
        :class="layoutFormat2"
      >
        <AnalogClock
          ref="ac"
          :border-width="borderWidth"
          border-color="#E8E8E8"
          scale-color="#2B2B2e"
          :size="clockSize"
          :scale-type="scaleType"
          :show-shadow="true"
          hand-type="pins"
          minute-hand-color="blue"
          hour-hand-color="black"
          :time="clock1"
        />
        <div
          v-if="isClock2"
        >
          Clock Start
        </div>
      </div>
      <div
        v-if="isClock2"
        :class="layoutFormat2"
      >
        <AnalogClock
          ref="ac"
          :border-width="borderWidth"
          border-color="#E8E8E8"
          scale-color="#2B2B2B"
          :size="clockSize"
          :scale-type="scaleType"
          :show-shadow="true"
          hand-type="pins"
          minute-hand-color="blue"
          hour-hand-color="black"
          :time="clock2"
        />
        <div>
          Clock End
        </div>
      </div>
    </div>
    <q-card
      class="q-mt-xl bg-gray text-blue"
      align="left"
      style="width: 60%; font-size: 15pt"
      flat
    >
      <vue-mathjax
        :formula="formula"
        :options="{}"
      />
    </q-card>
  </div>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'
import { VueMathjax } from 'vue-mathjax'
import { AnalogClock } from 'vue-analog-clock'

export default {
  name: 'QuestionMathClock',
  components: {
    'vue-mathjax': VueMathjax,
    'AnalogClock': AnalogClock
  },
  props: {
    index: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      scaleType: 'arabic'
    }
  },
  computed: {
    ...mapGetters('questions', ['getQuestions']),
    clock1: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].Clocks[0] : ''
    },
    clock2: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].Clocks[1] : ''
    },
    isClock2: function () {
      if (this.getQuestions[this.index].Clocks.length > 1) {
        return true
      } else {
        return false
      }
    },
    formula: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].UpTexts[0] : ''
    },
    clockSize: function () {
      if (this.$q.screen.width > 1020) {
        return (this.$q.screen.width / 5)
      } else if (this.$q.screen.width > 750) {
        return (this.$q.screen.width / 4)
      } else {
        return (180)
      }
    },
    borderWidth: function () {
      return (this.$q.screen.width / 130)
    },
    layoutFormat1: function () {
      if (this.$q.screen.width > 1020) {
        return 'row q-gutter-lg justify-center'
      } else {
        return 'q-gutter-lg justify-center'
      }
    },
    layoutFormat2: function () {
      if (this.$q.screen.width > 1020) {
        return ''
      } else {
        return 'col-2'
      }
    }
  },
  methods: {
  }
}
</script>

<style>
</style>
