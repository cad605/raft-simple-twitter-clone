import React, { useEffect, useState } from "react";
import { ErrorBoundary } from "react-error-boundary";
import ErrorFallback from "../components/ErrorFallback";
import { Tabs, TabList, Tab, TabPanels, TabPanel } from "@reach/tabs";
import axios from "axios";
import UserTweets from "./UserTweets";
import UserFollowers from "./UserFollowers";
import UserFollowing from "./UserFollowing";
import { useAuth } from '../context/auth-context'

export default function Profile() {
  const {user} = useAuth()
  const API = "http://localhost:8080/api/v1/getUser/" + user["id"];

  const [state, setState] = useState({
    status: "pending",
    results: null,
    error: null,
    tabIndex: 0,
  });
  const { status, results, error, tabIndex } = state;

  const tabs = ["Tweets", "Followers", "Following"];

  const handleTabsChange = (index) => {
    setState({ ...state, tabIndex: index });
  };

  async function queryDatabase() {
    const request = axios.get(API);

    return axios.all([request]).then(
      axios.spread(async (...responses) => {
        if (responses && responses[0]["data"]["data"]["success"]) {
          return responses[0]["data"]["data"]["user"][0];
        } else {
          const error = {
            message: responses?.errors?.map((e) => e.message).join("\n"),
          };
          return Promise.reject(error);
        }
      })
    );
  }
  useEffect(() => {
    setState({ ...state, status: "pending" });
    queryDatabase().then(
      (results) => {
        setState({ ...state, status: "resolved", results });
      },
      (error) => {
        setState({ ...state, status: "rejected", error });
      }
    );
  }, []);

  if (status === "pending") {
    return <p>Loading...</p>;
  } else if (status === "rejected") {
    console.log("throwing error");
    throw error;
  } else if (status === "resolved") {
    return (
      <>
        <div className="container h-14 bg-white border-solid border-slate-200 border-b">
          <div className="py-4 px-4">
            <h1 className="font-bold text-lg">Profile</h1>
          </div>
        </div>
        <div className="bg-white container min-h-fit pb-4 border-b">
          <div className="h-full py-4 px-4">
            <div className="font-semibold text-lg">{results["fullname"]}</div>
            <div className="flex-none w-full mt-2 font-normal">
              <div className="text-base font-normal text-gray-500">
                {results["handle"]}
              </div>
            </div>
            <div className="flex-none w-full mt-2 font-normal">
              <div className="text-base">{results["bio"]}</div>
            </div>
            <div className="flex w-full mt-2 text-sm font-normal text-gray-500">
              <svg
                xmlns="http://www.w3.org/2000/svg"
                className="h-5 w-5"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z"
                />
              </svg>
              Joined {new Date(results["created_at"]).toUTCString()}
            </div>
          </div>
        </div>
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
              <ErrorBoundary FallbackComponent={ErrorFallback}>
                <UserTweets></UserTweets>
              </ErrorBoundary>
            </TabPanel>
            <TabPanel>
              <ErrorBoundary FallbackComponent={ErrorFallback}>
                <UserFollowers></UserFollowers>
              </ErrorBoundary>
            </TabPanel>
            <TabPanel>
              <ErrorBoundary FallbackComponent={ErrorFallback}>
                <UserFollowing></UserFollowing>
              </ErrorBoundary>
            </TabPanel>
          </TabPanels>
        </Tabs>
      </>
    );
  }
}
