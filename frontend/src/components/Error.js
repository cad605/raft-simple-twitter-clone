
export default function ErrorFallback({error, resetErrorBoundary}) {
  return (
    <div className="flex p-4 justify-center">
      <p className="text-lg font-semibold text-gray-500">Nothing to see here...</p>
      {/* <pre>{error.message}</pre> */}
    </div>
  )
}