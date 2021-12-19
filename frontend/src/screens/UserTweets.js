import React, { useEffect, useState } from "react";
import { ErrorBoundary } from "react-error-boundary";
import ErrorFallback from "../components/ErrorFallback";
import List from "../components/List";
import FeedListItem from "../components/FeedListItem";
import axios from "axios";
import { useAuth } from "../context/auth-context";

export default function UserTweets() {
  const { user } = useAuth();
  const API = "http://localhost:8080/api/v1/getTweetsByUser/" + user["id"];

  const [state, setState] = useState({
    status: "pending",
    results: null,
    error: null,
  });
  const { status, results, error } = state;

  async function queryDatabase() {
    const request = axios.get(API);

    return axios.all([request]).then(
      axios.spread(async (...responses) => {
        if (responses && responses[0]["data"]["data"]["success"]) {
          return responses[0]["data"]["data"]["tweet"];
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
        <ErrorBoundary FallbackComponent={ErrorFallback}>
          {results && results.length > 0 ? (
            <List>
              {results.map((tweet) => (
                <FeedListItem key={tweet["id"]} tweet={tweet} />
              ))}
            </List>
          ) : (
            <ErrorFallback></ErrorFallback>
          )}
        </ErrorBoundary>
      </>
    );
  }
}
