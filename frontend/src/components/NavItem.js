import { Link } from "react-router-dom";

export default function NavItem({ href, isActive, children }) {
  return (
    <li >
      <div className="border rounded-md hover:border-sky-500 hover:text-sky-500">
        <Link
          to={href}
          className={`flex justify-center space-x-2 font-bold text-lg py-2 rounded-md ${
            isActive ? "bg-sky-500 text-white" : "bg-white"
          }`}
        >
          {children}
        </Link>
      </div>
    </li>
  );
}
