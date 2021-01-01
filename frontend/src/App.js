import logo from './logo.svg';
import './App.css';
import NewEndpointForm from './newEndpoint'

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <NewEndpointForm />
      </header>
    </div>
  )
}

export default App;
