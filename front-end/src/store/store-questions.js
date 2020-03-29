
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
  },
  doAddNewQuestion (state, payload) {
    state.questions.push(payload)
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
