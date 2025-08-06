import "./App.css";
import { CreatePostForm } from "./CreatePostForm";

export const API_URL =
  import.meta.env.VITE_API_URL || "http://localhost:9090/v1";

function App() {
  return (
    <>
      <div>App home screen</div>
      <CreatePostForm />
    </>
  );
}

export default App;
