import axios from "axios";
import React from "react";
import "./App.css";

function App() {
  const getTasks = async () => {
    axios({
      method: "GET",
      url: "api/v1/tasks",
      baseURL: "http://127.0.0.1:8080",
    })
      .then((res) => console.log(res.data))
      .catch((e) => console.warn(e));
  };
  return (
    <div className="App">
      <button onClick={() => getTasks()}>Test</button>
    </div>
  );
}

export default App;
