import React, { useEffect, useState } from "react";
import { ErrorBoundary } from "react-error-boundary";
import ErrorFallback from "./components/ErrorFallback";
import { Tabs, TabList, Tab, TabPanels, TabPanel } from "@reach/tabs";
import { useAuth } from "./context/auth-context";

function RegisterForm({ onSubmit, submitButton }) {
  function handleSubmit(event) {
    event.preventDefault();
    const { fullname, password, bio, handle } = event.target.elements;

    onSubmit({
      fullname: fullname.value,
      password: password.value,
      bio: bio.value,
      handle: "@" + handle.value,
    });
  }

  return (
    <form onSubmit={handleSubmit} className="p-2 border m-2">
      <div className="flex flex-col space-y-2 relative text-gray-400 focus-within:text-gray-600 block">
        <input
          className="p-2 focus:ring-2 focus:ring-sky-500 focus:outline-none w-full text-sm leading-6 text-gray-900 placeholder-gray-400 rounded-md py-2 pl-10 ring-1 ring-gray-200 shadow-sm"
          id="fullname"
          placeholder="Full Name"
        />
        <input
          className="focus:ring-2 focus:ring-sky-500 focus:outline-none w-full text-sm leading-6 text-gray-900 placeholder-gray-400 rounded-md py-2 pl-10 ring-1 ring-gray-200 shadow-sm"
          id="handle"
          placeholder="Handle"
        />
        <input
          className="focus:ring-2 focus:ring-sky-500 focus:outline-none w-full text-sm leading-6 text-gray-900 placeholder-gray-400 rounded-md py-2 pl-10 ring-1 ring-gray-200 shadow-sm"
          id="password"
          type="password"
          placeholder="Password"
        />
        <textarea
          className="focus:ring-2 focus:ring-sky-500 focus:outline-none w-full text-sm leading-6 text-gray-900 placeholder-gray-400 rounded-md py-2 pl-10 ring-1 ring-gray-200 shadow-sm"
          id="bio"
          placeholder="Bio"
        />
        <div className="p-2 rounded-md bg-sky-500 text-white text-center text-lg font-bold">
          {React.cloneElement(
            submitButton,
            { type: "submit" },
            ...(Array.isArray(submitButton.props.children)
              ? submitButton.props.children
              : [submitButton.props.children])
          )}
        </div>
      </div>
    </form>
  );
}

function LoginForm({ onSubmit, submitButton }) {
  function handleSubmit(event) {
    event.preventDefault();
    const { fullname, password } = event.target.elements;

    onSubmit({
      fullname: fullname.value,
      password: password.value,
    });
  }

  return (
    <form onSubmit={handleSubmit} className="p-2 border m-2">
      <div className="flex flex-col space-y-2 relative text-gray-400 focus-within:text-gray-600 block">
        <input
          className="p-2 focus:ring-2 focus:ring-sky-500 focus:outline-none w-full text-sm leading-6 text-gray-900 placeholder-gray-400 rounded-md py-2 pl-10 ring-1 ring-gray-200 shadow-sm"
          id="fullname"
          placeholder="Full Name"
        />
        <input
          className="focus:ring-2 focus:ring-sky-500 focus:outline-none w-full text-sm leading-6 text-gray-900 placeholder-gray-400 rounded-md py-2 pl-10 ring-1 ring-gray-200 shadow-sm"
          id="password"
          type="password"
          placeholder="Password"
        />
        <div className="p-2 rounded-md bg-sky-500 text-white text-center text-lg font-bold">
          {React.cloneElement(
            submitButton,
            { type: "submit" },
            ...(Array.isArray(submitButton.props.children)
              ? submitButton.props.children
              : [submitButton.props.children])
          )}
        </div>
      </div>
    </form>
  );
}

function UnauthenticatedApp() {
  const { login, register } = useAuth();

  const [state, setState] = useState({
    tabIndex: 0,
  });
  const { tabIndex } = state;

  const tabs = ["Login", "Register"];

  const handleTabsChange = (index) => {
    setState({ ...state, tabIndex: index });
  };

  return (
    <div className="h-screen flex grid grid-cols-6">
      <div className="flex flex-1 flex-col overflow-hidden col-start-2 border-solid border border-slate-200 border-t-0 border-b-0 rounded-none"></div>
      <div className="flex flex-col flex-1 overflow-hidden col-start-3 col-span-2 border-solid border border-slate-200 border-t-0 border-b-0 border-r-0 border-l-0 rounded-none">
        <ErrorBoundary FallbackComponent={ErrorFallback}>
          <div className="flex-1 overflow-y-scroll">
            <Tabs index={tabIndex} onChange={handleTabsChange}>
              <TabList className="flex justify-center space-x-4 bg-white border-b">
                {tabs.map((tab, i) => {
                  return (
                    <Tab
                      key={i}
                      className={`font-semibold px-6 py-1 ${
                        tabIndex === i ? "border-b-2 border-sky-500" : ""
                      }`}
                    >
                      {tab}
                    </Tab>
                  );
                })}
              </TabList>
              <TabPanels className="py-4">
                <TabPanel>
                  <div aria-label="Login form" title="Login">
                    <LoginForm
                      onSubmit={login}
                      submitButton={<button variant="primary">Login</button>}
                    />
                  </div>
                </TabPanel>
                <TabPanel>
                  <div
                    className="p-0.5"
                    aria-label="Registration form"
                    title="Register"
                  >
                    <RegisterForm
                      onSubmit={register}
                      submitButton={
                        <button variant="secondary">Register</button>
                      }
                    />
                  </div>
                </TabPanel>
              </TabPanels>
            </Tabs>
          </div>
        </ErrorBoundary>
      </div>
      <div className="flex flex-col flex-1 overflow-hidden border-solid border border-slate-200 border-t-0 border-b-0 rounded-none"></div>
    </div>
  );
}

export default UnauthenticatedApp;
