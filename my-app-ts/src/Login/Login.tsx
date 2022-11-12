import { useState, useEffect } from 'react';
import '../App.css';
import { UserResponse } from '../type'
import SelectLabels from './SelectLabels';
import { Link } from "react-router-dom";

type Props = {
    NameId: string
    setNameId: React.Dispatch<React.SetStateAction<string>>
};

const Login = (props:Props) => {
  const [users,setUsers] = useState<UserResponse[]>([]);

  useEffect(() => {
    console.log("users",users)
  }, [users])

  //TODO name無しでログインしようとした時にブロックする

  return (
    <div className="App">
      <header className="App-header">
        <p>
          Login
        </p>
        <SelectLabels NameId={props.NameId} setNameId={props.setNameId} users={users} setUsers={setUsers}/>
        <Link to='/pointranking'>Login</Link>
      </header>
    </div>
  );
}

export default Login;