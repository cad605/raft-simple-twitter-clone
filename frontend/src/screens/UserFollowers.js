import React, { useEffect, useState } from "react";
import List from "../components/List";
import axios from "axios";
import UserListItem from "../components/UserListItem";

export default function UserFollowers() {
  const API = "http://localhost:8080/api/v1/getFollowingByUser/1";

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
          console.log(responses[0]["data"]["data"]["user"]);
          return responses[0]["data"]["data"]["user"];
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
        <List>
          {results.map((user) => (
            <UserListItem key={user["id"]} user={user} />
          ))}
        </List>
      </>
    );
  }
}
