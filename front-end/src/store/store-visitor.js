/*
This store save the current information
*/
// import Vue from 'vue'

const state = {
  visitor: {
    visitorLog: [
      {
        name: '马大帅',
        gender: '男',
        ethnic: '汉',
        id: '12345678912345',
        dob: '20010101',
        block: '高大上小区',
        enter: '进入',
        deviceid: 'device-123',
        datetime: '2020-01-15 10:11:50'
      },
      {
        name: '王大拿',
        gender: '男',
        ethnic: '汉',
        id: '12345678912345',
        dob: '19750101',
        block: '高大上小区',
        enter: '进入',
        deviceid: 'device-123',
        datetime: '2020-01-16 12:21:50'
      },
      {
        name: '彪哥',
        gender: '男',
        ethnic: '汉',
        id: '12345678912345',
        dob: '19900101',
        block: '高大上小区',
        enter: '出',
        deviceid: 'device-123',
        datetime: '2020-01-17 13:25:50'
      }
    ]
  }
}

const mutations = {
  addVisitor (state, payload) {
    let valid = true
    // need valid the input
    if (valid === true) {
      console.log('add new visitor in vuex store')
      state.visitor.visitorLog.push(payload)
      // Vue.set(state.visitor.visitorLog, payload, [])
    }
  },

  removeVisitor (state, payload) {
    for (let i = 0; i < state.visitor.visitorLog.length; i++) {
      if (state.visitor.visitorLog[i].datetime === payload.datetime &&
        state.visitor.visitorLog[i].deviceid === payload.deviceid) {
        state.visitor.visitorLog.splice(i, 1)
        console.log('Number of visitors - ', state.visitor.visitorLog.length)
      }
    }
  },

  setFilter (state, payload) {
  }
}

const actions = {
}

const getters = {
  visitor: (state) => {
    return state.visitor
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
