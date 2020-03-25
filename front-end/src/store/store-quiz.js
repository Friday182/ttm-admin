// import Vue from 'vue'

const state = {
  quiz: [{
    QuizId: '',
    Grade: 0,
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
  addQuiz (state, payload) {
    state.quiz.push(payload)
  },
  setQuizList (state, payload) {
    state.quiz = payload
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
