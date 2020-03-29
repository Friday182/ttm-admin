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
mutation AddQuiz($QuizId: String!, $Grade: Int!, $Desc: String!, $Operator: String!, $Comment: String!) {
  AddQuiz(
    quiz: {
      QuizId: $QuizId,
      Grade: $Grade,
      Desc: $Desc,
      Operator: $Operator,
      Comment: $Comment
    }
  )
}`

export const ADD_TASKLOG_MUTATION = gql`
mutation AddTaskLog($gid: String!, $kps: String!, $totalQue: Int!, $totalCorrect: Int!, $totalSec: Int!, $totalScore: Int!, $details: String!) {
  AddTaskLog(
    gid: $gid
    totalQue: $totalQue,
    totalCorrect: $totalCorrect,
    totalSec: $totalSec,
    totalScore: $totalScore,
    kps: $kps,
    details: $details
  ) {
    Gid
    Stickers
    Stars
  }
}`
