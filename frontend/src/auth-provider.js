import axios from "axios";

const localStorageKey = '__auth_provider_token__'

async function getUserFromLocalStorage() {
  return JSON.parse(window.localStorage.getItem(localStorageKey))
}

function handleUserResponse(user) {
  window.localStorage.setItem(localStorageKey, JSON.stringify(user))
  return user
}

function login({fullname, password}) {
  return client('login', {fullname, password}).then(handleUserResponse)
}

function register({fullname, password, bio, handle}) {
  return client('register', {fullname, password, bio, handle}).then(handleUserResponse)
}

async function logout() {
  window.localStorage.removeItem(localStorageKey)
}

const authURL = "http://localhost:8080/api/v1"

async function client(endpoint, data) {
  const config = {
    method: 'POST',
    headers: {'Content-Type': 'application/json'},
  }

  return axios.post(`${authURL}/${endpoint}`,JSON.stringify(data), config).then((response) => {
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

export {getUserFromLocalStorage, login, register, logout, localStorageKey}