import Vue from 'vue'
import Vuex from 'vuex'
import sessionPersistePlugin from './plugins/sessionPersistePlugin.js'
import student from './store-student'
import mentor from './store-mentor'
import currentUser from './store-current-user'
import localMentors from './store-local-mentors'
import localStudents from './store-local-students'
import questions from './store-questions'
import logs from './store-logs'
import desc from './store-desc'
Vue.use(Vuex)

const persistePlugin = sessionPersistePlugin()

/*
 * If not building with SSR mode, you can
 * directly export the Store instantiation
 */

export default function (/* { ssrContext } */) {
  const Store = new Vuex.Store({
    modules: {
      student: student,
      mentor: mentor,
      currentUser: currentUser,
      localMentors: localMentors,
      localStudents: localStudents,
      questions: questions,
      logs: logs,
      desc: desc
    },
    plugins: [persistePlugin],

    // enable strict mode (adds overhead!)
    // for dev mode only
    strict: process.env.DEV
  })

  return Store
}
