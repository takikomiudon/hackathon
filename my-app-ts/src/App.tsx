import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import {useState} from "react";
import Register from "./Register"
import Login from "./Login/Login"
import List from "./ListPage"
import Ranking from "./RankingPage"
import Form from "./FormPage/FormPage"
import Update from "./UpdatePage/UpdatePage"

function App() {
  const [nameid, setNameId] = useState("");
  const [id, setId] = useState("");
  return (
    <div className="app">
      <Router>
        <Routes>
          <Route path="/register" element={<Register/>}/>

          <Route path="/login" element={<Login NameId={nameid} setNameId={setNameId}/>}/>

          <Route path="/contributionlist" element={<List nameid={nameid} setId={setId}/>}/>

          <Route path="/pointranking" element={<Ranking nameid={nameid}/>}/>

          <Route path="/contributionpost" element={<Form nameid={nameid}/>}/>

          <Route path="/update" element={<Update id={id}/>}/>
        </Routes>
      </Router>
    </div>
  );
}

export default App;