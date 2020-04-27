import VueApollo from 'vue-apollo'
import { ApolloClient } from 'apollo-client'
import { createHttpLink } from 'apollo-link-http'
import { InMemoryCache } from 'apollo-cache-inmemory'

// HTTP connection to the API
const httpLink = createHttpLink({
  // You should use an absolute URL here
  uri: 'http://localhost:8085/query'
  // uri: 'https://ttm2020.pythonanywhere.com/graphql'
  // uri: 'https://www.transfertestmaster.com/admin/graphql/'
})

// Cache implementation
const cache = new InMemoryCache()

// Create the apollo client
const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
  connectToDevTools: true,
  defaultOptions: {
    query: {
      fetchPolicy: 'network-only'
    }
  },
  request: (operation) => {
    operation.setContext({
      ContentType: 'application/x-www-form-urlencoded;charset=UTF-8',
      headers: {
        authorization: 'sessionStorage', // .getItem('token') || null,
        ContentType: 'application/x-www-form-urlencoded;charset=UTF-8'
      }
    })
  }
})

export const apolloProvider = new VueApollo({
  defaultClient: apolloClient,
  errorHandler ({ graphQLErrors, networkError }) {
    if (graphQLErrors) {
      graphQLErrors.map(({ message, locations, path }) =>
        console.log(
          'GraphQL error:', message // `[GraphQL error]: Message: ${message}, Location: ${locations}, Path: ${path}`
        )
      )
    }
    if (networkError) {
      console.log('[Network error]:', networkError)
    }
  }
})

export default ({ app, Vue }) => {
  Vue.use(VueApollo)
  app.apolloProvider = apolloProvider
}
