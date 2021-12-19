import React, { useEffect, useState } from "react";
import { ErrorBoundary } from "react-error-boundary";
import ErrorFallback from "../components/ErrorFallback";
import List from "../components/List";
import axios from "axios";
import UserListItem from "../components/UserListItem";
import { useAuth } from "../context/auth-context";

/**
 * People that follow the user
 * @returns 
 */
export default function UserFollowers() {
  const { user } = useAuth();

  const [state, setState] = useState({
    status: "pending",
    results: null,
    error: null,
  });
  const { status, results, error } = state;

  async function queryDatabase() {
    const url = "http://localhost:8080/api/v1";
    const endpoint = "getFollowingByUser/" + user["id"];

    return axios.get(`${url}/${endpoint}`).then((response) => {
      if (response) {
        return response.data["users"];
      } else {
        const error = {
          message: response?.errors?.map((e) => e.message).join("\n"),
        };
        return Promise.reject(error);
      }
    });
  }

  async function handleFollow(userListItem) {
    const url = "http://localhost:8080/api/v1"
    const endpoint = "followUser"

    const config = {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
    }

    const data = {
      FollowerID: user["id"],
      FollowedID: userListItem["id"]
    }
  
    return axios.post(`${url}/${endpoint}`,JSON.stringify(data), config).then((response) => {
        if (response) {
          return response.data;
        } else {
          const error = {
            message: response?.errors?.map((e) => e.message).join("\n"),
          };
          return Promise.reject(error);
        }
      })
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
              {results.map((listItem) => (
                <UserListItem key={listItem["id"]} user={listItem} showFollow={false} isFollow={false} handleClick={handleFollow} />
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