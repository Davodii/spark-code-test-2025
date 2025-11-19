<script lang="ts">
  import Todo from "./lib/Todo.svelte";
  import type { TodoItem } from "./lib/types";

  let todos: TodoItem[] = $state([]);
  let title: String = $state("");
  let description: String = $state("");

  async function handleSubmit(e: SubmitEvent) {
    // Prevent default submit behaviour
    e.preventDefault();

    const data = {
        title,
        description,
    };

    // Send a POST request
    const response = await fetch("http://localhost:8080/", {
        method: "POST",
        headers: {
            // Content-Type is application/json as 
            // specififed in openAPI.yaml
            "Content-Type": "application/json",
        },
        body: JSON.stringify(data),
    });

    if (response.status !== 200) {
        console.log("Error adding a new todo item. ")
    }

    const result = await response.json();

    console.log("Added a new item to the list");

    // Add the new item to the client list
    todos.push(result);
  }

  async function fetchTodos() {
    try {
      const response = await fetch("http://localhost:8080/");
      if (response.status !== 200) {
        console.error("Error fetching data. Response status not 200");
        return;
      }

      const result = await response.json();

      // If we set todos to the response directly, then we would set
      // todos to null. This would be a problem when we send POST
      // requests and expect to be able to push() to todos.
      if (result !== null) {
        todos = result;
      }
    } catch (e) {
      console.error("Could not connect to server. Ensure it is running.", e);
    }
  }

  // Initially fetch todos on page load
  $effect(() => {
    fetchTodos();
  });
</script>

<main class="app">
  <header class="app-header">
    <h1>TODO</h1>
  </header>

  <div class="todo-list">
    {#each todos as todo}
      <Todo title={todo.title} description={todo.description} />
    {/each}
  </div>

  <h2 class="todo-list-form-header">Add a Todo</h2>
  <form onsubmit={handleSubmit} class="todo-list-form">
    <input bind:value={title} placeholder="Title" name="title" />
    <input bind:value={description} placeholder="Description" name="description" />
    <button>Add Todo</button>
  </form>
</main>

<style>
  .app {
    color: white;
    background-color: #282c34;

    text-align: center;
    font-size: 24px;

    min-height: 100vh;
    padding: 20px;
  }

  .app-header {
    font-size: calc(10px + 4vmin);
    margin-top: 50px;
  }

  .todo-list {
    margin: 50px 100px 0px 100px;
  }

  .todo-list-form-header {
    margin-top: 100px;
  }

  .todo-list-form {
    margin-top: 10px;
  }
</style>
