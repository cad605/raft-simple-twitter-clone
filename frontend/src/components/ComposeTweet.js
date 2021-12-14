export default function ComposeTweet() {
  return (
    <div className="container h-32 pb-4">
      <div className="h-full py-4 px-4 bg-white border-solid border-slate-200 border-b">
        <h1 className="font-bold text-lg pb-2">What's happening?</h1>
        <label
          htmlFor="tweet"
          className="relative text-gray-400 focus-within:text-gray-600 block"
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            className="h-6 w-6 pointer-events-none w-8 h-8 absolute top-1/2 transform -translate-y-1/2 right-3 stroke-sky-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              strokeLinecap="round"
              strokeLinejoin="round"
              strokeWidth={2}
              d="M8 12h.01M12 12h.01M16 12h.01M21 12c0 4.418-4.03 8-9 8a9.863 9.863 0 01-4.255-.949L3 20l1.395-3.72C3.512 15.042 3 13.574 3 12c0-4.418 4.03-8 9-8s9 3.582 9 8z"
            />
          </svg>
          <input
            className="focus:ring-2 focus:ring-sky-500 focus:outline-none w-full text-sm leading-6 text-gray-900 placeholder-gray-400 rounded-md py-2 pl-10 ring-1 ring-gray-200 shadow-sm"
            type="text"
            id="tweet"
            placeholder="Tweet something..."
          ></input>
        </label>
      </div>
    </div>
  );
}
