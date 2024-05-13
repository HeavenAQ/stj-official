import { useQuery } from "@tanstack/react-query";
import { useNavigate } from "react-router-dom";
import { getUser } from "../api/user";

const User = () => {
  const navigate = useNavigate();
  const { data: user } = useQuery({
    queryKey: ["user"],
    queryFn: () => getUser(),
  });

  if (user === undefined) {
    navigate("/login");
  }

  return (
    <div className="mt-20 w-full h-11 text-center bg-slate-300">
      <h1>{user?.data.email}</h1>
    </div>
  );
};

export default User;
