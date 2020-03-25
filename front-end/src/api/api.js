import Axios from 'axios'
Axios.defaults.baseURL = 'http://127.0.0.1:8082/'
// Axios.defaults.baseURL = 'https://www.transfertestmaster.com/hywt/v1/'

// Add requset pre-process
Axios.interceptors.request.use(
  config => {
    // do something before request is sent
    console.log(config)
    if (localStorage.getItem('access_token')) {
      // Axios.defaults.headers.post['JWT_TOKEN'] = localStorage.getItem('access_token')
    }
    config.headers['Content-Type'] = 'application/x-www-form-urlencoded;charset=UTF-8'
    // config.headers['JWT_TOKEN'] = localStorage.getItem('access_token')
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

/****************************************************************
 * Home page urls
 */

// login
export const login = (params) => {
  return Axios.post('/login', params)
}

// Signup
export const signup = (params) => {
  return Axios.post('/signup', params)
}

// Add visitor record
export const addRecord = (params) => {
  return Axios.post('/api/visitor', params)
}

// Drawer item: Latest News
export const latestNews = (params) => {
  return Axios.get('/latest-news/', params)
}
