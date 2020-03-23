<template>
  <q-card
    class="bg-gray text-blue q-px-lg"
    style="width: 70%"
  >
    <q-card-section>
      <div
        v-for="(item,key) in upTexts"
        :key="key"
        class="bg-gray text-blue"
        align="left"
        style="width: 100%; font-size: 17pt"
      >
        {{ item }}
      </div>
    </q-card-section>
    <q-card-section
      align="center"
    >
      <v-chart
        ref="chart"
        :options="chart"
        theme="macarons"
        autoresize
      />
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
    <q-card-section>
      <div
        v-for="(item,key) in downTexts"
        :key="key"
        class="bg-gray text-blue"
        align="left"
        style="width: 100%; font-size: 17pt"
      >
        {{ item }}
      </div>
    </q-card-section>
  </q-card>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'
import { VueMathjax } from 'vue-mathjax'
import ECharts from '../common/ECharts.vue'
import 'echarts/lib/chart/bar'
import 'echarts/lib/chart/line'
import 'echarts/lib/chart/pie'
import 'echarts/lib/chart/scatter'
import 'echarts/lib/component/singleAxis'
import 'echarts/lib/component/tooltip'
import 'echarts/lib/component/legend'
import 'echarts/lib/component/title'
import 'echarts/lib/component/dataset'
import 'zrender/lib/svg/svg'

// built-in theme
import 'echarts/theme/macarons'

export default {
  name: 'QuestionMathChart',
  components: {
    'vue-mathjax': VueMathjax,
    'v-chart': ECharts
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
    ...mapGetters('questions', ['getQuestions']),
    upTexts: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].UpTexts : ''
    },
    downTexts: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].DownTexts : ''
    },
    formula: function () {
      return (this.getQuestions[this.index]) ? this.getQuestions[this.index].Formula : ''
    },
    chart: function () {
      return (this.getQuestions[this.index]) ? JSON.parse(this.getQuestions[this.index].Charts[0]) : ''
    }
  },
  mounted () {
  },

  methods: {
  }
}
</script>

<style>
</style>
