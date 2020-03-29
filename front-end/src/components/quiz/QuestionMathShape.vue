<template>
  <div>
    <q-card
      class="bg-gray text-blue q-px-lg"
      style="width: 70%; font-size: 17pt"
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
      <q-card-section
        align="center"
      >
        <v-stage
          :config="configKonva"
        >
          <v-layer>
            <v-circle
              v-for="(item,key) in configCircles"
              :key="key"
              :config="item"
            />
            <v-line
              v-for="(item,key) in configLines"
              :key="key"
              :config="item"
            />
            <v-ellipse
              v-for="(item,key) in configEllipses"
              :key="key"
              :config="item"
            />
            <v-rect
              v-for="(item,key) in configRects"
              :key="key"
              :config="item"
            />
            <v-regular-polygon
              v-for="(item,key) in configPolygons"
              :key="key"
              :config="item"
            />
            <v-arrow
              v-for="(item,key) in configArrows"
              :key="key"
              :config="item"
            />
            <v-star
              v-for="(item,key) in configStars"
              :key="key"
              :config="item"
            />
            <v-arc
              v-for="(item,key) in configArcs"
              :key="key"
              :config="item"
            />
            <v-wedge
              v-for="(item,key) in configWedges"
              :key="key"
              :config="item"
            />
          </v-layer>
          <v-layer>
            <v-text
              v-for="(item,key) in configTexts"
              :key="key"
              :config="item"
            />
          </v-layer>
        </v-stage>
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
  name: 'QuestionMathShape',
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
    },
    configKonva: function () {
      // console.log('window size: ', window.innerWidth, ':', window.innerWidth)
      let w = window.innerWidth / 2
      let h = 400
      if (w < 500) {
        w = 500
      }
      return {
        width: w,
        height: h
      }
    },
    configCircles: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'circle') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configLines: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'line') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configEllipses: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'ellipse') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configRects: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'rect') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configPolygons: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'polygon') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configArrows: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'arrow') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configStars: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'star') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configArcs: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'arc') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configWedges: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'wedge') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    },
    configTexts: function () {
      let tmp = []
      if (this.currentQuestion) {
        for (let i = 0; i < this.currentQuestion.Shapes.length; i++) {
          let shape = JSON.parse(this.currentQuestion.Shapes[i])
          if (shape.type === 'text') {
            tmp.push(shape.config)
          }
        }
      }
      return tmp
    }
  },
  methods: {
  }
}
</script>
