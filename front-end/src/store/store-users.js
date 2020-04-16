// import Vue from 'vue'

const state = {
  users: [{
    Gid: '',
    Username: '',
    Password: '',
    Email: '',
    Role: '',
    Name: '',
    Mobile: '',
    Comment: '',
    LastLoginTime: ''
  }]
}

const mutations = {
  addNewUser (state, payload) {
    state.users.push(payload)
  },
  setUserList (state, payload) {
    state.users = payload
  },
  removeUser (state, payload) {
    for (let i = 0; i < state.users.length; i++) {
      if (state.users[i].Gid === payload.Gid) {
        state.users.splice(i, 1)
        break
      }
    }
  }
}

const actions = {
}

const getters = {
  getUserList: (state) => {
    return state.users
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
