import React from "react";

export default function UserListItem({
  user,
  showFollow,
  isFollow,
  handleClick,
}) {
  function handleEvent() {
    handleClick(user);
  }

  return (
    <article className="flex items-start space-x-6 p-6 bg-white rounded-lg border">
      <div className="min-w-0 relative flex-auto">
        <h2 className="text-base font-semibold text-gray-900 truncate pr-20">
          {user["fullname"]}
        </h2>
        <dl className="mt-2 flex flex-wrap text-sm leading-6 font-medium">
          {showFollow && (
            <div className="absolute top-0 right-0 flex items-center space-x-1">
              <button className="text-sky-500" onClick={handleEvent}>
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  className="h-6 w-6"
                  fill={isFollow ? "currentColor" : "none"}
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M11.049 2.927c.3-.921 1.603-.921 1.902 0l1.519 4.674a1 1 0 00.95.69h4.915c.969 0 1.371 1.24.588 1.81l-3.976 2.888a1 1 0 00-.363 1.118l1.518 4.674c.3.922-.755 1.688-1.538 1.118l-3.976-2.888a1 1 0 00-1.176 0l-3.976 2.888c-.783.57-1.838-.197-1.538-1.118l1.518-4.674a1 1 0 00-.363-1.118l-3.976-2.888c-.784-.57-.38-1.81.588-1.81h4.914a1 1 0 00.951-.69l1.519-4.674z"
                  />
                </svg>
              </button>
            </div>
          )}
          <div>
            <dd className="text-base text-gray-500 px-1.5 ring-1 ring-gray-200 rounded">
              {user["handle"]}
            </dd>
          </div>
          <div className="flex-none w-full mt-2 font-normal">
            <dd className="text-base">{user["bio"]}</dd>
          </div>
        </dl>
      </div>
    </article>
  );
}
