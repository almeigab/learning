const user = () => {
  return {
    id: '1',
    userName: 'Gabriel',
  };
};

const users = () => {
  return [
    {
      id: '1',
      userName: 'Gabriel 1',
    },
    {
      id: '2',
      userName: 'Gabriel 2',
    },
    {
      id: '3',
      userName: 'Gabriel 3',
    },
  ];
};

export const userResolvers = {
  Query: {
    user,
    users,
  },
};
