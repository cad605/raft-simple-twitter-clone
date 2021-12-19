import * as React from "react";
import * as auth from "../auth-provider";

const AuthContext = React.createContext();
AuthContext.displayName = "AuthContext";

function AuthProvider(props) {
  const [user, setUser] = React.useState(null);

  React.useEffect(() => {
    
    async function bootstrapAppData() {
      const userInStorage = await auth.getUserFromLocalStorage();
      console.log(userInStorage)
      if (userInStorage) {
        setUser(userInStorage);
      }
    }

    bootstrapAppData();
  }, []);

  const login = React.useCallback(
    (form) => auth.login(form).then((user) => setUser(user)),
    [setUser]
  );
  
  const register = React.useCallback(
    (form) =>
      auth
        .register(form)
        .then(() =>
          auth.login(form).then((user) => setUser(user))
        ),
    [setUser]
  );
  
  const logout = React.useCallback(() => {
    auth.logout();
    setUser(null);
  }, [setUser]);

  const value = React.useMemo(
    () => ({ user, login, logout, register }),
    [login, logout, register, user]
  );

  return <AuthContext.Provider value={value} {...props} />;

  throw new Error(`Unhandled status:`);
}

function useAuth() {
  const context = React.useContext(AuthContext);

  if (context === undefined) {
    throw new Error(`useAuth must be used within a AuthProvider`);
  }
  return context;
}

export { AuthProvider, useAuth };
