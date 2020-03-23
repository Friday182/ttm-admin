
const getDefaultState = () => {
  return {
    questions: [{
      QueIdx: 0,
      Kp: '',
      StdSec: 0,
      AnswerType: '',
      QuestionType: '',
      UpTexts: [],
      DownTexts: [],
      Formula: [],
      Options: [],
      Answers: [],
      Tags: [],
      ChartType: '',
      Charts: [],
      Clocks: [],
      Tables: [],
      Shapes: [],
      AnswerText: '',
      Helper: false,
      Imgs: [],
      Tips: [],
      Choice: ''
    }]
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
  }
}

const actions = {
  setQuestions ({ commit }, payload) {
    commit('setAllQuestions', payload)
  },
  clearQuestions ({ commit }) {
    commit('resetState')
  }
}

const getters = {
  getQuestions: (state) => {
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
