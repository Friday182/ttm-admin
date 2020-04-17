
const getDefaultState = () => {
  return {
    questions: []
  }
}

const state = getDefaultState()

const mutations = {
  setAllQuestions (state, payload) {
    // state.questions = payload
    state.questions = Object.assign([], payload)
  },
  resetState (state) {
    state.questions = []
  },
  setQueChoice (state, payload) {
    if (state.questions[payload.index]) {
      state.questions[payload.index].Choice = payload.choice
    }
  },
  doAddNewQuestion (state, payload) {
    let newQue = true
    for (let i = 0; i < state.questions.length; i++) {
      if (state.questions[i].QueIdx === payload.QueIdx) {
        state.questions[i] = payload
        newQue = false
      }
    }
    if (newQue === true) {
      state.questions.push(payload)
    }
  }
}

const actions = {
  setQuestions ({ commit }, payload) {
    commit('setAllQuestions', payload)
  },
  clearQuestions ({ commit }) {
    commit('resetState')
  },
  addNewQuestion ({ commit }, payload) {
    commit('doAddNewQuestion', payload)
  }
}

const getters = {
  allQuestions: (state) => {
    return state.questions
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
