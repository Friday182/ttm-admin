
const getDefaultState = () => {
  return {
    logs: [] // The structure was defined in TasklogTable.vue
  }
}

const state = getDefaultState()

const mutations = {
  copyLogs (state, payload) {
    state.logs = payload
  },
  addLog (state, payload) {
    for (let i = 0; i < payload.length; i++) {
      state.logs.push(payload[i])
    }
  },
  delLog (state, payload) {
    for (let i = 0; i < state.logs.length; i++) {
      if (state.logs[i].id === payload) {
        state.logs.splice(i, 1)
      }
    }
  },
  tidyLog (state, payload) {
    if (state.logs.length > 180) {
      state.logs.splice(180, state.logs.length - 180)
    }
    if (state.logs[0].gid === payload) {
      let tmp = JSON.stringify(state.logs)
      // use gid-logs as the key in localstorage
      localStorage.setItem(payload + '-logs', tmp)
    }
  }
}

const getters = {
  getLogs: (state) => {
    return state.logs
  },
  getLastLogId: (state) => {
    if (state.logs.length > 0) {
      return state.logs[state.logs.length - 1].id
    } else {
      return 0
    }
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  getters
}
