// import React from 'react';
import './App.css';
import Header from './Header';
import MessageForm from './MessageForm';

function App() {
  return (
    <div className="App">
      {/* <header>
        <h1>Simple React Message App</h1>
      </header> */}
      <Header />
      <main>
        <MessageForm />
      </main>
    </div>
  );
}

export default App;
