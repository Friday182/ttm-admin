<template>
  <q-card
    class="bg-gray text-black"
    align="center"
    style="width: 80%"
    flat
  >
    <q-card-section>
      <q-btn
        class="q-mt-md"
        style="width: 100%; font-size: 15pt"
        text-color="blue"
        no-caps
      >
        <vue-mathjax
          :formula="formula"
          :options="{}"
        />
      </q-btn>
    </q-card-section>
    <q-card-section>
      <div class="row flex">
        <div
          v-for="(item, key) in moneyPic"
          :key="key"
        >
          <q-btn
            class="q-ma-sm q-pa-xs"
            flat
            :size="item.size"
            :icon="item.path"
            @click="choose(item.value, key)"
          >
            <q-tooltip content-class="bg-white text-primary">
              {{ item.tip }}
            </q-tooltip>
          </q-btn>
        </div>
      </div>
    </q-card-section>
    <q-card-section>
      <q-linear-progress
        :value="status"
        :color="staColor"
        class="q-mt-sm"
      />
    </q-card-section>
    <q-card-section align="left">
      Selected:
    </q-card-section>
    <q-card-section align="left">
      <q-btn
        v-for="n in selection.length"
        :key="n"
        class="q-mx-xs"
        flat
        size="30px"
        :icon="selection[n-1].path"
        @click="removePic(n-1)"
      />
    </q-card-section>
    <q-card-section>
      <q-btn
        push
        label="Ok"
        text-color="white"
        color="blue"
        style="font-size: 12pt"
        @click="selectDone"
      />
    </q-card-section>
  </q-card>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'
import { VueMathjax } from 'vue-mathjax'

export default {
  name: 'QuestionMathMoney',
  components: {
    'vue-mathjax': VueMathjax
  },
  props: {
    index: {
      type: Number,
      default: 0
    },
    status: {
      type: Number,
      default: 0
    },
    staColor: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
      moneyData: [
        {
          path: 'img:statics/money/1.jpg',
          size: '30px',
          tip: '1 Pence Coin',
          value: 1
        },
        {
          path: 'img:statics/money/2.jpg',
          size: '35px',
          tip: '2 Pence Coin',
          value: 2
        },
        {
          path: 'img:statics/money/5.jpg',
          size: '40px',
          tip: '5 Pence Coin',
          value: 5
        },
        {
          path: 'img:statics/money/10.jpg',
          size: '40px',
          tip: '10 Pence Coin',
          value: 10
        },
        {
          path: 'img:statics/money/20.jpg',
          size: '40px',
          tip: '20 Pence Coin',
          value: 20
        },
        {
          path: 'img:statics/money/50.jpg',
          size: '45px',
          tip: '50 Pence Coin',
          value: 50
        },
        {
          path: 'img:statics/money/100.jpg',
          size: '50px',
          tip: '1 Pound Coin',
          value: 100
        },
        {
          path: 'img:statics/money/200.jpg',
          size: '50px',
          tip: '2 Pound Coin',
          value: 200
        },
        {
          path: 'img:statics/money/500.jpg',
          size: '80px',
          tip: '5 Pound Note',
          value: 500
        },
        {
          path: 'img:statics/money/1000.jpg',
          size: '80px',
          tip: '10 Pound Note',
          value: 1000
        },
        {
          path: 'img:statics/money/2000.jpg',
          size: '80px',
          tip: '20 Pound Note',
          value: 2000
        }
      ],
      selection: [],
      ansCorrect: false,
      totalValue: 0
    }
  },
  computed: {
    ...mapGetters('currentQuestion', ['currentQuestion']),
    formula: function () {
      return (this.currentQuestion) ? this.currentQuestion.UpTexts[0] : ''
    },
    moneyPic: function () {
      let tmp = []
      if (this.currentQuestion) {
        console.log('total - ', this.currentQuestion.Options.length)
        for (let i = 0; i < this.currentQuestion.Options.length; i++) {
          console.log('looking - ', this.currentQuestion.Options[i])
          for (let j = 0; j < this.moneyData.length; j++) {
            if (Number(this.currentQuestion.Options[i]) === this.moneyData[j].value) {
              console.log('find - ', this.currentQuestion.Options[i])
              tmp.push(this.moneyData[j])
              break
            }
          }
        }
      }
      if (tmp.length > 0) {
        return tmp
      } else {
        return this.moneyData
      }
    }
  },
  mounted () {
    console.log('mount money comp')
  },
  methods: {
    choose (value, idx) {
      console.log('select - ', idx)
      if (this.selection.length < 15) {
        this.selection.push(this.moneyPic[idx])
        this.totalValue += value
      } else {
        // Alert maximum shoose 15 each time
        const dialog = this.$q.dialog({
          title: 'Alert',
          message: 'Maximum using <span class="text-red text-h5"> 15 </span> items',
          html: true,
          ok: {
            push: true
          }
        }).onDismiss(() => {
          clearTimeout(timer)
        })
        const timer = setTimeout(() => {
          dialog.hide()
        }, 2000)
      }
      console.log('total value: ', this.totalValue)
    },
    selectDone () {
      if (this.totalValue === Number(this.currentQuestion.Answers[0])) {
        this.ansCorrect = true
      } else {
        this.ansCorrect = false
      }
      this.$emit('ansCorrect', this.ansCorrect)
      this.totalValue = 0
      this.selection = []
    },
    removePic (idx) {
      this.totalValue -= this.selection[idx].value
      this.selection.splice(idx, 1)
      console.log('remove pic idx - ', idx, this.totalValue)
    }
  }
}
</script>

<style>
</style>
