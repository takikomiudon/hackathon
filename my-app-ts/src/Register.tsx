import { useState } from 'react';
import './App.css';

const Register = () => {
  const [name,setName] = useState("")
  const onSubmit = async (name:string) => {
    if (name.length > 50 || name.length==0){
      alert("Please enter a name between 1 and 50 characters")
      return 
    }
    
    try{
      const response = await fetch(
        "http://localhost:8000/register",
        {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: JSON.stringify({
            name: name
          }),
        });
        
        if (!response.ok){
          throw Error(`Failed to create user ${response.status}`);
        }

        setName("");
      } catch(err) {
        console.error(err);
    }
  };

  const submit = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault()
    onSubmit(name)
  }

  return (
    <div className="App">
      <header className="App-header">
        <p>
          User Register
        </p>
        <input
          type={"text"}
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <button onClick={(e) => submit(e)}>Submit</button>
      </header>
    </div>
  );
}

export default Register;