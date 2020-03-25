<template>
  <q-dialog v-model="active">
    <q-card>
      <q-card-section class="q-px-lg q-ma-md primary">
        <div class="text-h6 text-blue">
          {{ msg }}
        </div>
      </q-card-section>
      <q-card-actions align="center">
        <q-btn
          label="OK"
          color="primary"
          @click="cancelAlert(1)"
        />
      </q-card-actions>
    </q-card>
  </q-dialog>
</template>

<script type="text/javascript">

export default {
  name: 'AlertCom',

  props: {
    active: {
      type: Boolean,
      default: false
    },
    msg: {
      type: String,
      default: ''
    }
  },
  data () {
    return {
    }
  },
  created () {
    document.addEventListener('keyup', this.handleKeyup)
  },
  destroyed () {
    document.removeEventListener('keyup', this.handleKeyup)
  },
  methods: {
    handleKeyup (event) {
      if (event.keyCode === 13) {
        this.cancelAlert()
      }
    },
    cancelAlert (opt) {
      this.$emit('cancelAlert', opt)
    }
  }
}
</script>

<style>
</style>
