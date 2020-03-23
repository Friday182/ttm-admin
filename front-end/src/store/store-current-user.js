/*
This store save the current user information
*/
const state = {
  currentUser: {
    Gid: '',
    Name: '',
    Username: '',
    LastLoginTime: '',
    Role: 'na'
  }
}

const mutations = {
  updateUser (state, payload) {
    Object.assign(state.currentUser, payload)
  }
}

const actions = {
  updateUser ({ commit }, payload) {
    commit('updateUser', payload)
  }
}

const getters = {
  currentUser: (state) => {
    return state.currentUser
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
