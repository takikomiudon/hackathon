import { useState, useEffect } from 'react';
import './App.css';
import { Contribution, Contributed } from './type';
import { Link } from "react-router-dom";

type Props = {
  nameid: string;
  setId: React.Dispatch<React.SetStateAction<string>>
}

function ListPage(props:Props) {
  const [contribution, setContribution] = useState<Contribution []>([]);
  const [contributed, setContributed] = useState<Contributed []>([]);

  const fetchUsers = async ()=>{
    try {
      const res = await fetch ("http://localhost:8000/mycontribution?nameid=" + props.nameid,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json"
        },
      });
      if (!res.ok){
        throw Error(`Failed to fetch users: ${res.status}`);
      }
      const contribution = await res.json();
      setContribution(contribution);
    } catch(err) {
      console.error(err)
    }

    try {
        const res = await fetch ("http://localhost:8000/mycontributed?nameid=" + props.nameid,
        {
          method: "GET",
          headers: {
            "Content-Type": "application/json"
          },
        });
        if (!res.ok){
          throw Error(`Failed to fetch users: ${res.status}`);
        }
        const contributed = await res.json();
        setContributed(contributed);
      } catch(err) {
        console.error(err)
      }
  };

  const onDelete = async (id: string) =>{
    try{
      const response = await fetch(
        "http://localhost:8000/contributiondelete",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            id: id
          }),
        });
        
        if (!response.ok){
          throw Error(`Failed to delete contribution ${response.status}`);
        }

      } catch(err) {
        console.error(err);
    }
    fetchUsers()
  }

  useEffect(() => {
    fetchUsers()
  },[]);

  return (
    <div className="App">
      <header className="App-header">
        <h1>Your Contribution</h1>
        {contribution.map((c,index) =>
          <section key={index}>
            <p>FROM {c.Contributor}  {c.Point}Pt  Message:{c.Message}</p>
          </section>
        )}
        <h1>Your Contributed</h1>
        {contributed.map((c,index) =>

          <section key={index}>
            <p>TO {c.Contributor}  {c.Point}Pt  Message:{c.Message}</p>
            <Link to='/update'>
              <button onClick={() => props.setId(c.Id)}>Update</button>
            </Link>
            <button
              onClick={(e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
              e.preventDefault()
              onDelete(c.Id)
            }}>
            Delete
            </button>
          </section>
        )}
        <Link to='/contributionpost'>Contribution Post</Link>
        <Link to='/contributionlist'>Contribution List</Link>
        <Link to='/pointranking'>Point Ranking</Link>
        <Link to='/register'>User Register</Link>
        <Link to='/login'>Login Page</Link>
      </header>

    </div>
  );
}

export default ListPage;