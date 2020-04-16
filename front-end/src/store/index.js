import Vue from 'vue'
import Vuex from 'vuex'
import sessionPersistePlugin from './plugins/sessionPersistePlugin.js'
import currentQuestion from './store-current-question'
import quiz from './store-quiz'
import users from './store-users'
import questions from './store-questions'
import currentUser from './store-current-user'

Vue.use(Vuex)

const persistePlugin = sessionPersistePlugin()

/*
 * If not building with SSR mode, you can
 * directly export the Store instantiation
 */

export default function (/* { ssrContext } */) {
  const Store = new Vuex.Store({
    modules: {
      quiz: quiz,
      users: users,
      currentUser: currentUser,
      questions: questions,
      currentQuestion: currentQuestion
    },
    plugins: [persistePlugin],

    // enable strict mode (adds overhead!)
    // for dev mode only
    strict: process.env.DEV
  })

  return Store
}
