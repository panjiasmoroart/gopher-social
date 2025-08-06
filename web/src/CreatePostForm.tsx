import React, { useState } from "react";
import { API_URL } from "./App";

export const CreatePostForm: React.FC = () => {
  const [title, setTitle] = useState("");
  const [content, setContent] = useState("");

  const handleSubmit = async () => {
    try {
      const response = await fetch(`${API_URL}/posts`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
          Authorization: `Bearer foo`, // Ganti "foo" dengan token yang valid
        },
        body: JSON.stringify({ title, content }),
      });

      if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
      }

      const data = await response.json();
      console.log("Post successful:", data);

      setTitle("");
      setContent("");
    } catch (error) {
      console.error("Error submitting post:", error);
    }
  };

  return (
    <div>
      <label>
        <input
          type="text"
          placeholder="Title..."
          value={title}
          onChange={(e) => setTitle(e.target.value)}
        />
      </label>
      <br /> <br />
      <label>
        <textarea
          placeholder="What's in your mind..."
          value={content}
          onChange={(e) => setContent(e.target.value)}
        ></textarea>
      </label>
      <br />
      <button onClick={handleSubmit}>Share</button>
    </div>
  );
};
