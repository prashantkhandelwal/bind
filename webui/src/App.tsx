import React from 'react';
import { BrowserRouter, Route, Routes } from "react-router-dom";
import Header from './layouts/Header';
import Main from './Main';

interface IAppProps {

}

export const App: React.FC<IAppProps> = (props: IAppProps) => {
  return (
    <div>
      {
        <React.Fragment>
          <BrowserRouter>
            <Header />
            <Routes>
              <Route path="/" element={<Main />} />
              <Route path="/archive" />
            </Routes>
          </BrowserRouter>
        </React.Fragment>
      }
    </div>
  );
}

export default App;
