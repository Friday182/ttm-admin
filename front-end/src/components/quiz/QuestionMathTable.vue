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
      <q-table
        :style="table.style"
        :title="table.tableTitle"
        :data="table.tableData"
        :columns="table.columns"
        :separator="table.separator"
        hide-bottom
        row-key="id"
        :table-header-style="{ backgroundColor: 'lightblue' }"
        no-data-label="Something Wrong. Contact your mentor."
        fixed-center
        bordered
        dense
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

export default {
  name: 'QuestionMathTable',
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
      defWidth: 'width: 30%'
    }
  },
  computed: {
    ...mapGetters('currentQuestion', ['currentQuestion']),
    table: function () {
      return (this.currentQuestion) ? JSON.parse(this.currentQuestion.Tables[0]) : ''
    },
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
