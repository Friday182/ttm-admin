<template>
  <div>
    <q-btn
      label="True"
      glossy
      outlined
      color="green"
      class="q-mr-md"
      @click="selected('1')"
    />
    <q-btn
      label="False"
      glossy
      outlined
      color="red"
      @click="selected('0')"
    />
  </div>
</template>

<script type="text/javascript">
import { mapGetters } from 'vuex'

export default {
  name: 'TrueFalseAnswer',
  props: {
    index: {
      type: Number,
      default: 0
    }
  },
  data () {
    return {
      ansCorrect: false
    }
  },
  computed: {
    ...mapGetters('questions', ['getQuestions'])
  },
  methods: {
    selected (opt) {
      if (opt === this.getQuestions[this.index].Answer[0]) {
        this.ansCorrect = true
      } else {
        this.ansCorrect = false
      }
      this.$emit('ansCorrect', this.ansCorrect)
    }
  }
}
</script>

<style>
</style>
