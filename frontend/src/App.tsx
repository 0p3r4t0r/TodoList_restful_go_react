import axios from "axios";
import React, { useEffect } from "react";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import "./App.css";
import { useGetTasks } from "./api";

axios.defaults.baseURL = "http://127.0.0.1:8080/api/v1";

function Tasks() {
  const { data: getTasksResponse } = useGetTasks();

  useEffect(() => {
    if (getTasksResponse) {
      console.log(getTasksResponse.data);
    }
  }, [getTasksResponse]);

  return (
    <div>
      <ul>
        {getTasksResponse &&
          getTasksResponse.data.map((task) => (
            <li>{`${task.title}: ${task.points}`}</li>
          ))}
      </ul>
    </div>
  );
}

function App() {
  const queryClient = new QueryClient();
  return (
    <QueryClientProvider client={queryClient}>
      <Tasks />
    </QueryClientProvider>
  );
}

export default App;
