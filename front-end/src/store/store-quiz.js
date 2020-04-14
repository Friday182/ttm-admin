// import Vue from 'vue'

const state = {
  quiz: [{
    QuizId: '',
    Grade: '',
    Desc: '',
    Status: '',
    Operator: '',
    Approver: '',
    Comment: '',
    CreatedAt: '',
    UpdatedAt: ''
  }]
}

const mutations = {
  addNewQuiz (state, payload) {
    state.quiz.push(payload)
  },
  setQuizList (state, payload) {
    state.quiz = payload
  },
  removeQuiz (state, payload) {
    for (let i = 0; i < state.quiz.length; i++) {
      if (state.quiz[i].QuizId === payload.QuizId) {
        state.quiz.splice(i, 1)
        break
      }
    }
  }
}

const actions = {
}

const getters = {
  getQuizList: (state) => {
    return state.quiz
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
