import * as React from 'react';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import { UserResponse } from '../type'

type Props = {
  id: string
  contributorId: string
  setContributorId: React.Dispatch<React.SetStateAction<string>>
  point: number
  setPoint: React.Dispatch<React.SetStateAction<number>>
  message: string
  setMessage: React.Dispatch<React.SetStateAction<string>>
  onSubmit: (nameid: string, contributorId: string, point: number, message: string) => void
};

const Update = (props: Props) => {
  const submit = (e: React.MouseEvent<HTMLButtonElement, MouseEvent>) => {
    e.preventDefault()
    props.onSubmit(props.id, props.contributorId, props.point, props.message)
  }
  
  const [users,setUsers] = React.useState<UserResponse[]>([]);

  const handleChange = (event: SelectChangeEvent) => {
    props.setContributorId(event.target.value);
  };

  const fetchUsers = async ()=>{
    try {
      const res = await fetch ("http://localhost:8000/login",
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json"
        },
      });
      if (!res.ok){
        throw Error(`Failed to fetch users: ${res.status}`);
      }
      const users = await res.json();
      setUsers(users);
    } catch(err) {
      console.error(err)
    }
  };

  React.useEffect(() => {
    fetchUsers()
  },[]);

  return (
    <div>
      <form style={{ display: "flex", flexDirection: "column" }}>
      <FormControl sx={{ m: 1, minWidth: 120 }}>
      <InputLabel id="demo-simple-select-helper-label">User</InputLabel>
        <Select
          labelId="demo-simple-select-helper-label"
          id="demo-simple-select-helper"
          value={props.contributorId}
          label="User"
          onChange={handleChange}
        >
          <MenuItem value="">
            <em>None</em>
          </MenuItem>
          {users.map((user,index) => (
            <MenuItem
              key={index}
              value={user.NameId}
            >
              {user.Name} ID:{user.NameId}
            </MenuItem>
          ))}
          </Select>
        </FormControl>
        <label>Point: </label>
        <input
          type={"number"}
          value={props.point}
          onChange={(e) => props.setPoint(e.target.valueAsNumber)}
        ></input>
        <label>Message: </label>
        <input
          type={"text"}
          style={{ marginBottom: 20 }}
          value={props.message}
          onChange={(e) => props.setMessage(e.target.value)}
        ></input>
        <button onClick={(e) => submit(e)}>Submit</button>
      </form>
  </div>
  );
};

export default Update;