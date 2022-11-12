import '../App.css';
import { useState, useEffect } from "react";
import Form from './Form';
import { Link } from "react-router-dom";

type Props = {
  nameid: string
}

const FormPage = (props:Props) => {
  const [contributorId, setContributorId] = useState("");
  const [point, setPoint] = useState(0);
  const [message, setMessage] = useState(""); 

  const onSubmit = async (nameid:string, contributorId:string, point:number, message:string) => {
    if (!contributorId) {
      alert("Please enter contributor's ID");
      return;
    }

    if (point <= 0 || point > 100){
      alert("Please enter point between 1 and 100");
      return;
    }

    if (message.length > 100){
      alert("Please enter a message shorter than 100 characters")
    }

    try{
      const response = await fetch(
        "http://localhost:8000/contributionpost",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            nameid: nameid,
            contributorId: contributorId,
            point: String(point),
            message: message
          }),
        });

        console.log(JSON.stringify({
          nameid: nameid,
          contributorId: contributorId,
          point: point,
          message: message}))

        if (!response.ok){
          throw Error(`Failed to create user ${response.status}`);
        }

        setContributorId("");
        setPoint(0);
        setMessage("");
      } catch(err) {
        console.error(err);
    }
  };

  useEffect(() => {
    console.log("point",point)
  }, [point])

  return (
    <div className="App">
      <header className="App-header">
        <Form 
          nameid={props.nameid} 
          contributorId={contributorId} 
          setContributorId={setContributorId} 
          point={point} 
          setPoint={setPoint} 
          message={message} 
          setMessage={setMessage} 
          onSubmit={onSubmit} 
        />
        <Link to='/contributionpost'>Contribution Post</Link>
        <Link to='/contributionlist'>Contribution List</Link>
        <Link to='/pointranking'>Point Ranking</Link>
        <Link to='/register'>User Register</Link>
        <Link to='/login'>Login Page</Link>
      </header>
    </div>
  );
};

export default FormPage;