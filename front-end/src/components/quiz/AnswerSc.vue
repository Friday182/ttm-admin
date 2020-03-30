<template>
  <div
    v-if="useradio"
  >
    <div class="row">
      <div class="col-2" />
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-radio
          val="op1"
          label="A:"
          @input="selected('1')"
        >
          <vue-mathjax
            :formula="options[0]"
            :options="{}"
          />
        </q-radio>
      </div>
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-radio
          val="op2"
          @input="selected('2')"
        >
          <vue-mathjax
            :formula="options[1]"
            :options="{}"
          />
        </q-radio>
      </div>
    </div>
    <div class="row">
      <div class="col-2" />
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-radio
          val="op3"
          @input="selected('3')"
        >
          <vue-mathjax
            :formula="options[2]"
            :options="{}"
          />
        </q-radio>
      </div>
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-radio
          val="op4"
          @input="selected('4')"
        >
          <vue-mathjax
            :formula="options[3]"
            :options="{}"
          />
        </q-radio>
      </div>
    </div>
  </div>
  <div
    v-else
    class="q-mt-lg"
  >
    <div class="row">
      <div class="col-3" />
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-btn
          outline
          color="primary"
          style="width: 100%; font-size: 15pt"
          no-caps
          label="A:"
          selected
          @click="selected('1')"
        >
          <q-space />
          <vue-mathjax
            :formula="options[0]"
            :options="{}"
          />
          <q-space />
        </q-btn>
      </div>
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-btn
          outline
          color="primary"
          style="width: 100%; font-size: 15pt"
          no-caps
          label="B:"
          @click="selected('2')"
        >
          <q-space />
          <vue-mathjax
            :formula="options[1]"
            :options="{}"
          />
          <q-space />
        </q-btn>
      </div>
    </div>
    <div class="row">
      <div class="col-3" />
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-btn
          outline
          color="primary"
          label="C:"
          style="width: 100%; font-size: 15pt"
          no-caps
          @click="selected('3')"
        >
          <q-space />
          <vue-mathjax
            :formula="options[2]"
            :options="{}"
          />
          <q-space />
        </q-btn>
      </div>
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-btn
          outline
          color="primary"
          label="D:"
          style="width: 100%; font-size: 15pt"
          no-caps
          @click="selected('4')"
        >
          <q-space />
          <vue-mathjax
            :formula="options[3]"
            :options="{}"
          />
          <q-space />
        </q-btn>
      </div>
    </div>
    <div
      class="row"
    >
      <div class="col-3" />
      <div class="col-3 q-pa-sm q-ml-sm">
        <q-btn
          outline
          color="primary"
          label="E:"
          style="width: 100%; font-size: 15pt"
          no-caps
          @click="selected('5')"
        >
          <q-space />
          <vue-mathjax
            :formula="options[4]"
            :options="{}"
          />
          <q-space />
        </q-btn>
      </div>
      <div class="col-3" />
    </div>
  </div>
</template>

<script type="text/javascript">
import { VueMathjax } from 'vue-mathjax'
import { mapGetters } from 'vuex'

export default {
  name: 'SignalChoiceAnswer',
  components: {
    'vue-mathjax': VueMathjax
  },
  props: {
    index: {
      type: Number,
      default: 0
    },
    useradio: {
      type: Boolean,
      default: true
    }
  },
  data () {
    return {
      ansCorrect: false
    }
  },
  computed: {
    ...mapGetters('currentQuestion', ['currentQuestion']),
    options: function () {
      return (this.currentQuestion) ? this.currentQuestion.Options : ['', '', '', '', '']
    }
  },
  methods: {
    selected (opt) {
    }
  }
}
</script>

<style>
</style>
