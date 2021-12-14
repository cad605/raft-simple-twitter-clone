export default function FeedListItem({ tweet }) {
  return (
    <article className="flex items-start space-x-6 p-6 bg-white rounded-lg border">
      <div className="min-w-0 relative flex-auto">
        <h2 className="text-base font-semibold text-gray-900 truncate pr-20">
          {tweet["authorName"]}
        </h2>
        <dl className="mt-2 flex flex-wrap text-sm leading-6 font-medium">
          <div className="absolute top-0 right-0 flex items-center space-x-1">
            <dd>{new Date(tweet["created_at"]).toUTCString()}</dd>
          </div>
          <div>
            <dd className="text-base font-normal text-gray-500 px-1.5 ring-1 ring-gray-200 rounded">
              {tweet["authorHandle"]}
            </dd>
          </div>
          <div className="flex-none w-full mt-2 font-normal">
            <dd className="text-base">{tweet["content"]}</dd>
          </div>
        </dl>
      </div>
    </article>
  );
}
