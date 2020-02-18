import React, { Component } from 'react';
import './App.css';
import axios from 'axios'

class App extends Component {
  constructor(){
    super()
    this.spellClick = this.spellClick.bind(this)
    this.readClick = this.readClick.bind(this)
  }

  spellClick(){
    //GET
    axios.get('http://0.0.0.0:8080/spell?number='+document.getElementById("spell").value,{  crossDomain: true
    }).then(response => document.getElementById("result").innerHTML = 
      "<span>Spell Result</span> = "+response.data.text)
  }

  readClick(){
    //POST
    axios.post('http://0.0.0.0:8080/read?text='+document.getElementById("read").value,{
      crossDomain: true
    }).then(response => document.getElementById("result").innerHTML = 
    "<span>Read Result</span> = "+response.data.number)
  }
  
  render() {
    return (
      <div className="App">
        <div className="App-header">
          <h2>Welcome to Indonesian Numeral Speller</h2>
        </div>
        <div className="App-container">
          <div className="spell">
            <h4>Spelling (number to text)</h4>
            <p>Insert valid number. Eg: 23, 245, 4587. Maximum number can handle is 2 billion</p>
            <form>
              <input type="number" name="spell" placeholder="Insert number" id="spell"></input>
              <button id="spell-submit" onClick={this.spellClick}>submit</button>
            </form>
          </div>
          <div className="read">
            <h4>Reading (text to number)</h4>
            <p>Insert valid text. Make sure there is no typo on it. Maximum number can handle is 2 billion</p>
            <form>
              <input type="text" name="read" placeholder="Masukkan teks" id="read"></input>
              <button id="read-submit" onClick={this.readClick}>submit</button>
            </form>
          </div>
          <div className="result" id="result" hidden>
          </div>
        </div>
        <div className="App-footer">
          <p>Created by : Juniardi Akbar</p>
        </div>
      </div>
    );
  }
}

export default App;