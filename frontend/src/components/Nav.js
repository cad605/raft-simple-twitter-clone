export default function Nav({ children }) {
  return (
    <div className="container mx-auto w-full h -full">
      <ul className="grid grid-cols-1 gap-4 py-4 px-6 text-sm font-medium w-full h-full">
        {children}
      </ul>
    </div>
  );
}
