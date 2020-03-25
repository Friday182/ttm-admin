import gql from 'graphql-tag'

export const GET_QUIZ_QUERY = gql`query GetQuizMsg{
  GetQuiz {
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
  }
}`
