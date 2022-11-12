import * as React from 'react';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import FormControl from '@mui/material/FormControl';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import { UserResponse } from '../type'

type Props = {
  NameId: string
  setNameId: React.Dispatch<React.SetStateAction<string>>
  users: UserResponse[]
  setUsers: React.Dispatch<React.SetStateAction<UserResponse[]>>
};

export default function SelectLabels(props:Props) {
  const handleChange = (event: SelectChangeEvent) => {
    props.setNameId(event.target.value);
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
      props.setUsers(users);
    } catch(err) {
      console.error(err)
    }
  };

  React.useEffect(() => {
      fetchUsers()
  },[]);

  return (
    <div>
      <FormControl sx={{ m: 1, minWidth: 120 }}>
        <InputLabel id="demo-simple-select-helper-label">User</InputLabel>
        <Select
          labelId="demo-simple-select-helper-label"
          id="demo-simple-select-helper"
          value={props.NameId}
          label="User"
          onChange={handleChange}
        >
          <MenuItem value="">
            <em>None</em>
          </MenuItem>
          {props.users.map((user,index) => (
            <MenuItem
              key={index}
              value={user.NameId}
            >
              {user.Name} ID:{user.NameId}
            </MenuItem>
          ))}
        </Select>
      </FormControl>
    </div>
  );
}
