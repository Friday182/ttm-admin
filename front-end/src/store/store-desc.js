
const getDefaultState = () => {
  return {
    desc: {
      MA: [],
      MB: [],
      MC: [],
      MD: [],
      ME: [],
      MF: [],
      MG: [],
      MH: [],
      MI: []
    }
  }
}

const state = getDefaultState()

const mutations = {
  copyDesc (state, payload) {
    state.desc = payload
  },
  saveDesc (state) {
    let tmp = JSON.stringify(state.desc)
    localStorage.setItem('kpDescriptions', tmp)
  }
}

const getters = {
  getDesc: (state) => {
    return state.desc
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  getters
}
