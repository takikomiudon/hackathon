import { createContext, useState } from "react";

type Props = {
    children: string
}

export const UserContext = createContext({});

export const UserProvider = (props:Props) => {
    const [userInfo, setUserInfo] = useState("");
    return (
        <UserContext.Provider value={{userInfo, setUserInfo}}>
            {props.children}
        </UserContext.Provider>
    )
}