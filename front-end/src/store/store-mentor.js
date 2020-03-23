
const state = {
  mentor: {
    id: '',
    name: '',
    mentorEmail: '',
    refreshStudent: false,
    contacts: [],
    students: []
  }
}

const mutations = {
  setMentor (state, payload) {
    Object.assign(state.mentor, payload)
  },
  mentorAddStudent (state, payload) {
    console.log('add student in vuex: ', payload)
    payload.index = state.mentor.students.length
    state.mentor.students.push(payload)
    for (let i = 0; i < state.mentor.students.length; i++) {
      if (state.mentor.students[i].gid === payload.gid) {
        state.mentor.students[i] = payload
        break
      }
    }
  },
  mentorDelStudent (state, name) {
    for (let i = 0; i < state.mentor.students.length; i++) {
      if (state.mentor.students[i].name === name) {
        state.mentor.students.splice(i, 1)
      }
    }
    // update the index
    for (let i = 0; i < state.mentor.students.length; i++) {
      state.mentor.students.index = i
    }
  },
  setRefreshStudent (state, payload) {
    state.mentor.refreshStudent = payload
  },
  setAssignment (state, payload) {
    console.log('find student - assignment: ', payload.gid, payload.newVal)
    for (let i = 0; i < state.mentor.students.length; i++) {
      if (state.mentor.students[i].gid === payload.gid) {
        state.mentor.students[i].assignment = payload.newVal
      }
    }
  },
  setStudentPw (state, payload) {
    for (let i = 0; i < state.mentor.students.length; i++) {
      if (state.mentor.students[i].gid === payload.gid) {
        state.mentor.students[i].password = payload.newVal
      }
    }
  }
}

const actions = {
}

const getters = {
  currentMentor: (state) => {
    return state.mentor
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
