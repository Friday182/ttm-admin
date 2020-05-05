import gql from 'graphql-tag'

export const GET_QUIZ_QUERY = gql`query GetQuizMsg($gid: String!){
  GetQuiz (gid: $gid){
    QuizId
    UsedCount
    Grade
    Desc
    Status
    Operator
    Approver
    Comment
    CreatedAt
    UpdatedAt
    Details {
      Ma
      Mb
      Mc
      Md
      Me
      Mf
      Mg
      Mh
      Mi
      Mj
    }
  }
}`

export const GET_USERS_QUERY = gql`query GetUsers($role: String!){
  GetUsers(role: $role) {
    Gid
    Username
    Password
    Name
    Role
    Email
    Mobile
    Comment
    LastLoginTime
  }
}`

export const TASK_QUESTIONS_QUERY = gql`query GetQuestions($gid: String!, $kp: String!){
  GetQuestions(gid: $gid, kp: $kp) {
    QueIdx
    Kp
    StdSec
    AnswerType
    QuestionType
    UpTexts
    DownTexts
    Formula
    Options
    Answers
    Tags
    Charts
    Clocks
    Tables
    Shapes
    AnswerText
    Helper
    Imgs
    Tips
    Status
  }
}`
