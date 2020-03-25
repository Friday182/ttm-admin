/*
This store save the current user information
*/
const state = {
  currentUser: {
    Gid: '',
    Name: '',
    Username: '',
    LastLoginTime: '',
    Role: 'na',
    tabs: [{
      name: '1',
      title: 'Dashboard'
    }],
    activeTab: '1',
    wsOk: false
  }
}

const mutations = {
  updateUser (state, payload) {
    Object.assign(state.currentUser, payload)
  },
  addTab (state, payload) {
    let exist = false
    state.currentUser.tabs.forEach((tab, index) => {
      if (tab.name === payload.name) {
        // tab already in the list
        exist = true
      }
    })
    if (exist === false) {
      state.currentUser.tabs.push(payload)
    }
    state.currentUser.activeTab = payload.name
  },

  doRemoveTab (state, payload) {
    console.log('removetab - ', payload.name)
    for (let i = 0; i < state.currentUser.tabs.length; i++) {
      if (state.currentUser.tabs[i].name === payload.name) {
        state.currentUser.tabs.splice(i, 1)
        console.log('new tabs', JSON.stringify(state.currentUser.tabs))
        if (i > 0) {
          state.currentUser.activeTab = state.currentUser.tabs[i - 1].name
        } else {
          if (state.currentUser.tabs.length > 0) {
            state.currentUser.activeTab = state.currentUser.tabs[0].name
          }
        }
      }
    }
  },

  setUserRole (state, payload) {
    state.currentUser.userRole = payload
  },

  setActiveTab (state, payload) {
    state.currentUser.activeTab = payload
  }
}

const actions = {
  removeTab ({ commit }, payload) {
    commit('doRemoveTab', payload)
  }
}

const getters = {
  currentUser: (state) => {
    return state.currentUser
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
