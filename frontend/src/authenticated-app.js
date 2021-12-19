import React from "react";
import { ErrorBoundary } from "react-error-boundary";
import ErrorFallback from "./components/ErrorFallback";
import Feed from "./screens/Feed";
import Profile from "./screens/Profile";
import Users from "./screens/Users";
import Nav from "./components/Nav";
import NavItem from "./components/NavItem";
import { Routes, Route } from "react-router-dom";

function AuthenticatedApp() {
  return (
    <div className="h-screen flex grid grid-cols-6">
      <div className="flex flex-1 flex-col overflow-hidden col-start-2 border-solid border border-slate-200 border-t-0 border-b-0 rounded-none">
        <SideNav></SideNav>
      </div>
      <div className="flex flex-col flex-1 overflow-hidden col-start-3 col-span-2 border-solid border border-slate-200 border-t-0 border-b-0 border-r-0 border-l-0 rounded-none">
        <ErrorBoundary FallbackComponent={ErrorFallback}>
          <div className="flex-1 overflow-y-scroll">
            <AppRoutes />
          </div>
        </ErrorBoundary>
      </div>
      <div className="flex flex-col flex-1 overflow-hidden border-solid border border-slate-200 border-t-0 border-b-0 rounded-none">
        <div className="container h-14 bg-white border-solid border-slate-200 border-b">
          <div className="py-4 px-4">
            <h1 className="font-bold text-lg">Who to follow</h1>
          </div>
        </div>
        <ErrorBoundary FallbackComponent={ErrorFallback}>
          <div className="flex-1 overflow-y-scroll py-4">
            <Users></Users>
          </div>
        </ErrorBoundary>
      </div>
    </div>
  );
}

function SideNav(params) {
  return (
    <Nav>
      <NavItem href="/" isActive>
        <svg
          xmlns="http://www.w3.org/2000/svg"
          className="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={2}
            d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
          />
        </svg>
        <p>Home</p>
      </NavItem>
      <NavItem href="/profile">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          className="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            strokeLinecap="round"
            strokeLinejoin="round"
            strokeWidth={2}
            d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
          />
        </svg>
        <p>Profile</p>
      </NavItem>
    </Nav>
  );
}

function AppRoutes() {
  return (
    <Routes>
      <Route path="/" element={<Feed />} />
      <Route path="/profile" element={<Profile />} />
      <Route path="*" element={<ErrorFallback />} />
    </Routes>
  );
}

export default AuthenticatedApp;
