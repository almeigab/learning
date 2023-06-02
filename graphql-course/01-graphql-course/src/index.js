import { ApolloServer, gql } from 'apollo-server';

const server = new ApolloServer({
  typeDefs: gql`
    type Query {
      user: User!
      users: [User!]!
    }

    type User {
      id: ID!
      userName: String!
    }
  `,
  resolvers: {
    Query: {
      user: () => ({
        id: 'ashui343',
        userName: 'Gabriel Almeida',
      }),
      users: () => [
        {
          id: '1',
          userName: 'user 1',
        },
        {
          id: '2',
          userName: 'user 2',
        },
        {
          id: '3',
          userName: 'user 3',
        },
      ],
    },
  },
});

server.listen(4003).then(({ url }) => {
  console.log(`Server listening on url ${url}`);
});
