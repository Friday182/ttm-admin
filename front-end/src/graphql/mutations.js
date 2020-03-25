import gql from 'graphql-tag'

export const LOGIN_MUTATION = gql`
mutation UserLogin($username: String!, $password: String!) {
  UserLogin(
    username: $username,
    password: $password
  ){
    Gid
    Name
    Username
    Role
    LastLoginTime
    Comment
  }
}`

export const ADD_QUIZ_MUTATION = gql`
mutation AddQuiz($quiz: Object!) {
  AddQuiz(
    quiz: $quiz
  )
}`
