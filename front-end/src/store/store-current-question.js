const state = {
  currentQuestion: {
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
    Charts: [],
    Clocks: [],
    Tables: [],
    Shapes: [],
    AnswerText: '',
    Helper: false,
    Imgs: [],
    Tips: [],
    Choice: ''
  }
}

const mutations = {
  doSetCurrentQuestion (state, payload) {
    state.currentQuestion = payload
  }
}

const actions = {
  setCurrentQuestion ({ commit }, payload) {
    commit('doSetCurrentQuestion', payload)
  }
}

const getters = {
  currentQuestion: (state) => {
    return state.currentQuestion
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
