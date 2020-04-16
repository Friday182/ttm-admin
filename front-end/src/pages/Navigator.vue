<template>
  <el-tabs
    v-model="activeTab"
    type="border-card"
    editable
    @edit="handleTabsEdit"
  >
    <el-tab-pane
      v-for="item in tabs"
      :key="item.name"
      :label="item.title"
      :name="item.name"
    >
      <div v-if="item.name==='1'">
        <dashboard />
      </div>
      <div v-else-if="item.name==='2-1'">
        <user-tab
          role="staff"
        />
      </div>
      <div v-else-if="item.name==='2-2'">
        <user-tab
          role="operator"
        />
      </div>
      <div v-else-if="item.name==='3-1'">
        <math-tab />
      </div>
      <div v-else-if="item.name==='3-2'">
        <english-tab />
      </div>
      <div v-else-if="item.name==='4'">
        <not-available />
      </div>
      <div v-else-if="item.name==='5'">
        <not-available />
      </div>
    </el-tab-pane>
  </el-tabs>
</template>

<script>
import { mapGetters, mapActions, mapMutations } from 'vuex'
// import { createSocket } from '../api/websocket'

export default {
  name: 'Navigator',
  components: {
    'dashboard': require('components/admin/Dashboard.vue').default,
    'math-tab': require('components/quiz/MathTab.vue').default,
    'english-tab': require('components/quiz/EnglishTab.vue').default,
    'user-tab': require('components/admin/UserTab.vue').default
  },
  data () {
    return {
    }
  },
  computed: {
    ...mapGetters('currentUser', ['currentUser']),

    loginRole: function () {
      return this.currentUser.Role
    },
    tabs: function () {
      return this.currentUser.tabs
    },
    activeTab: {
      get: function () {
        return this.currentUser.activeTab
      },
      set: function (newValue) {
        console.log('set active tab - ', newValue)
        // this.setActiveTab(newValue)
      }
    }
  },
  mounted () {
    console.log('mounted Navigator')
    // createSocket()
    // listen incoming message
    // window.addEventListener('onmessageWS', this.processData)
  },
  beforeDestroy () {
    // window.removeEventListener('onmessageWS', this.processData)
  },
  methods: {
    ...mapActions('currentUser', ['removeTab']),
    ...mapMutations('visitor', ['addVisitor']),
    ...mapMutations('currentUser', ['setActiveTab']),

    handleTabsEdit (targetName, action) {
      console.log('in edit - ', targetName, action)
      if (action === 'remove') {
        this.tabs.forEach((tab, index) => {
          if (tab.name === targetName && tab.name !== '1') {
            this.removeTab(tab)
          }
        })
      }
    },
    processData (e) {
      console.log('Handle WS message: ', e.detail.data.type)
      if (e.detail.data.type === 'VISITORLOG') {
        let logs = JSON.parse(e.detail.data.data)
        console.log('json data: ', logs)
        for (let i = 0; i < logs.length; i++) {
          console.log(logs[i].dob)
          this.addVisitor({
            name: logs[i].name,
            gender: logs[i].gender,
            ethnic: logs[i].ethnic,
            id: logs[i].id,
            dob: logs[i].dob,
            block: logs[i].block,
            enter: logs[i].enter,
            deviceid: logs[i].deviceid,
            datetime: logs[i].datetime
          })
        }
        this.setActiveTab('1-1')
      } else if (e.detail.data.type === 'ERROR') {
        console.log('Websocket Error message: ', e.detail.data.data)
      }
    }
  }
}
</script>
