import React from 'react';
import ReactDOM from 'react-dom/client';
import App from './App';
import { Wrapper } from './Index.style';

const root = ReactDOM.createRoot(
  document.getElementById('root') as HTMLElement
);
root.render(
  <Wrapper>
    <React.StrictMode>
      <App />
    </React.StrictMode>
  </Wrapper>
);
