import { useQuery } from "@tanstack/react-query";
import { getUser } from "../api/user";

const User = () => {
  const { data: user } = useQuery({
    queryKey: ["user"],
    queryFn: () => getUser(),
  });

  return (
    <div className="mt-20 w-full h-11 text-center bg-slate-300">
      {user === undefined ? (
        <a href="/login">Login First</a>
      ) : (
        <h1>{user?.data.email}</h1>
      )}
    </div>
  );
};

export default User;
