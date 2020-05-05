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

export const ADD_USER_MUTATION = gql`
mutation AddUser($name: String!, $role: String!, $email: String!, $mobile: String!, $comment: String!) {
  AddUser(
    user: {
      Username: $role,
      Password: $role,
      Name: $name,
      Role: $role,
      Email: $email,
      Mobile: $mobile,
      Comment: $comment
    }
  )
}`

export const DEL_USER_MUTATION = gql`
mutation DelUser($Gid: String!) {
  DelUser(
    gid: $Gid
  )
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

export const UPDATE_QUIZ_MUTATION = gql`
mutation UpdateQuiz($Gid: String!, $QuizId: String!, $Status: String!) {
  UpdateQuiz(
    gid: $Gid,
    quizId: $QuizId,
    status: $Status,
  )
}`

export const ADD_QUESTION_MUTATION = gql`
mutation AddQuestion($Gid: String!, $QueIdx: Int!, $Kp: String!, $StdSec: Int!, $AnswerType: String!, $QuestionType: String!,
  $UpTexts: String!, $DownTexts: String!, $Formula: String!, $Charts: String!, $Shapes: String!, $Tables: String!, $Clocks: String!,
  $Options: String!, $Answers: String!, $Tags: String!, $Status: Int!) {
  AddQuestion(
    que: {
      Gid: $Gid,
      QueIdx: $QueIdx,
      Kp: $Kp,
      StdSec: $StdSec,
      AnswerType: $AnswerType,
      QuestionType: $QuestionType,
      UpTexts: $UpTexts,
      DownTexts: $DownTexts,
      Formula: $Formula,
      Charts: $Charts,
      Shapes: $Shapes,
      Tables: $Tables,
      Clocks: $Clocks,
      Options: $Options,
      Answers: $Answers,
      Tags: $Tags,
      Status: $Status
    }
  ) {
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
