// This is Json object from localstorage
// It keeps the list of saved mentors

const state = {
  localMentors: [
    // gid: 0,
    // email: '',
    // token: '',
    // letter: '',
    // lastSigninDate: ''
  ]
}

const mutations = {
  updateMentors (state, payload) {
    state.localMentors = payload
  },
  removeMentor (state, key) {
    if (state.localMentors[key]) {
      state.localMentors.splice(key, 1)
    }
  },
  addMentor (state, payload) {
    state.localMentors.push(payload)
  },
  updateMentorDate (state, key) {
    if (state.localMentors[key]) {
      state.localMentors[key].lastSigninDate = Date()
    }
  }
}

const actions = {
  updateLocalMentors ({ commit }, payload) {
    commit('updateMentors', payload)
  },
  removeLocalMentor ({ commit }, key) {
    commit('removeMentor', key)
  },
  addLocalMentor ({ commit }, payload) {
    commit('addMentor', payload)
  },
  updateLocalMentorDate ({ commit }, key) {
    commit('updateMentorDate', key)
  }
}

const getters = {
  localMentors: (state) => {
    return state.localMentors
  },
  Mentor: (state) => key => {
    if (state.localMentors[key]) {
      return state.localMentors[key]
    } else {
      return ''
    }
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
