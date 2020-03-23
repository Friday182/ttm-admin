// import Vue from 'vue'

const state = {
  student: {
    gid: 0,
    mentorId: 0,
    name: '',
    age: 0,
    city: '',
    country: '',
    school: '',
    stickers: 0,
    stars: 0,
    isMember: false,
    membershipStart: '',
    membershipStop: '',
    lastLoginDate: '',
    subjects: [],
    quizzes: [],
    quizDone: false,
    stickerLog: [],
    contacts: []
  }
}

const mutations = {
  setStudent (state, payload) {
    Object.assign(state.student, payload)
  },
  setAssignmentDone (state, payload) {
    if (payload.English) {
      state.student.subjects[1].Done = payload.English
    }
    if (payload.Math) {
      console.log('assignmentDone Math - ', payload)
      for (let i = 0; i < state.student.subjects[0].Assignment.length; i++) {
        if (payload.Kp === state.student.subjects[0].Assignment[i].Kp) {
          state.student.subjects[0].Assignment[i].Done = payload.Math
          break
        }
      }
    }
  },
  setQuizDone (state, payload) {
    state.student.quizDone = payload
  }
}

const actions = {
  updateStudent ({ commit }, payload) {
    commit('setStudent', payload)
  },
  updateAssignmentDone ({ commit }, payload) {
    commit('setAssignmentDone', payload)
  },
  updateQuizDone ({ commit }, payload) {
    commit('setQuizDone', payload)
  }
}

const getters = {
  currentStudent: (state) => {
    return state.student
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
