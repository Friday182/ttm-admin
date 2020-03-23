<template>
  <div class="q-mx-xl q-px-xl">
    <q-separator />
    <div class="row">
      <div
        v-for="(item, key) in stickerData"
        :key="key"
        class="col-2"
      >
        <q-card class="my-card q-ma-md">
          <img
            :src="item.path"
            style="height: 120px; width: 120px"
          >
        </q-card>
      </div>
    </div>
    <stickers
      :active="showStickers"
      :select-num="selectNum"
      @selectSticker="getSticker"
    />
    <alert-msg
      :active="alert"
      :msg="alertMsg"
      @cancelAlert="alertClose"
    />
  </div>
</template>

<script type="text/javascript">
import { mapGetters, mapActions } from 'vuex'
import { UPDATE_STICKERS_MUTATION } from '../../graphql/mutations'

export default {
  name: 'StickerTable',
  components: {
    'alert-msg': require('components/common/AlertMsg.vue').default,
    'stickers': require('components/common/Stickers.vue').default
  },

  data () {
    return {
      alert: false,
      skipQuerySticker: false,
      localStickerLog: [],
      alertMsg: ''
    }
  },
  computed: {
    ...mapGetters('student', ['currentStudent']),
    stickerData: {
      // getter
      get: function () {
        let tmp = []
        for (let i = 0; i < this.currentStudent.stickerLog.length; i++) {
          tmp.push({
            id: i,
            path: 'statics/stickers/' + this.currentStudent.stickerLog[i]
          })
        }
        return tmp
      },
      // setter
      set: function (newValue) {
      }
    },
    stickerNum: function () {
      return this.currentStudent.stickers
    },
    selectNum: function () {
      if (this.currentStudent.stickers > this.currentStudent.stickerLog.length) {
        console.log('stickers, sticklog length - ', this.currentStudent.stickers, this.currentStudent.stickerLog.length)
        return this.currentStudent.stickers - this.currentStudent.stickerLog.length
      } else {
        return 0
      }
    },
    showStickers: function () {
      return (this.selectNum > 0)
    }
  },
  updated () {
  },
  methods: {
    ...mapActions('student', ['updateStudent']),
    updateSticker (newStickers) {
      console.log('update sticker table now - ', newStickers)
      let obj = []
      for (let i = 0; i < this.currentStudent.stickerLog.length; i++) {
        obj.push(this.currentStudent.stickerLog[i])
      }
      for (let j = 0; j < newStickers.length; j++) {
        obj.push(newStickers[j])
      }
      console.log('new sticker log - ', obj)
      this.updateStudent({
        stickerLog: obj
      })
    },
    getSticker (opt) {
      // this.showStickers = false
      console.log('get selected: ' + JSON.stringify(opt), opt.length, this.currentStudent.gid)
      this.$apollo
        .mutate({
          mutation: UPDATE_STICKERS_MUTATION,
          variables: {
            gid: this.currentStudent.gid,
            stickerLog: JSON.stringify(opt)
          }
        })
        .then(response => {
          if (response.data.SetStickerLog === false) {
            this.alertMsg = 'Update stickers failed, please report to your mentor.'
            this.alert = true
          } else {
            this.updateSticker(opt)
          }
        })
    },
    alertClose () {
      this.alert = false
      this.alertMsg = ''
    }
  }
}
</script>

<style lang="stylus" scoped>
.my-card
  width 100%
  max-width 100px
</style>
