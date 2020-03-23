// This is Json object from localstorage
// It keeps the list of saved students

const state = {
  localStudents: []
}

const mutations = {
  updateStudents (state, payload) {
    state.localStudents = payload
  },
  removeStudent (state, key) {
    if (state.localStudents[key]) {
      state.localStudents.splice(key, 1)
    }
  },
  addStudent (state, payload) {
    state.localStudents.push(payload)
  },
  updateStudentDate (state, key, payload) {
    if (state.localStudents[key]) {
      state.localStudents[key].lastSigninDate = Date()
    }
  }
}

const actions = {
  updateLocalStudents ({ commit }, payload) {
    commit('updateStudents', payload)
  },
  removeLocalStudent ({ commit }, key) {
    commit('removeStudent', key)
  },
  addLocalStudent ({ commit }, payload) {
    commit('addStudent', payload)
  },
  updateLocalStudentDate ({ commit }, key) {
    commit('updateStudentDate', key)
  }
}

const getters = {
  localStudents: (state) => {
    return state.localStudents
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
