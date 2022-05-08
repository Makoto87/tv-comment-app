import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

import {
      ApolloClient,
      InMemoryCache,
      ApolloProvider,
      createHttpLink,
} from "@apollo/client";

const link = createHttpLink({
      uri: `http://${process.env.REACT_APP_HOST}:${process.env.REACT_APP_PORT}/${process.env.REACT_APP_PATH}`,
});

const client = new ApolloClient({
      cache: new InMemoryCache(),
      link: link,
});

ReactDOM.render(
  <React.StrictMode>
      <ApolloProvider client={client}>
            <App />
      </ApolloProvider>
  </React.StrictMode>,
  document.getElementById('root')
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
