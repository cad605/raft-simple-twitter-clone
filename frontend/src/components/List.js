export default function List({ children }) {
  return (
    <div className="px-4">
        <ul className="space-y-3">{children}</ul>
    </div>
  );
}
