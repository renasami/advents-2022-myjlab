import { Box, Heading } from "@hope-ui/solid";
import { createSignal, Index, JSX } from "solid-js";
import { v4 } from "uuid";
import { Item } from "../../App";
import TodoItem from "../../components/TodoItem";
import { TodoItemProps } from "../../types";

const TodoPage = () => {
  const [isAuthenticated, setIsAuthenticated] = createSignal<boolean>(false);
  const [items, setItems] = createSignal<TodoItemProps[]>([]);
  const [inputValues, setInputValue] = createSignal<string>("");

  const handleOnAdd = (value: any) => {
    const item: TodoItemProps = {
      id: v4().toString(),
      value: value,
      done: false,
    };
    console.log(item);
    setItems([...items(), item]);
    setInputValue("");
  };
  const handleOnDelete = () => {};
  const handleOnEdit = () => {};

  const onInputHandler: JSX.EventHandlerUnion<HTMLInputElement, Event> = (
    e
  ) => {
    setInputValue((e.target as HTMLInputElement).value);
  };

  return (
    <div>
      <Heading size="4xl">Todo</Heading>
      <div>
        <input
          value={inputValues()}
          type="text"
          onInput={(i) => {
            onInputHandler(i);
          }}
        />
        <button onClick={() => handleOnAdd(inputValues())}>追加</button>
      </div>
      <Box ml="45%" mt="8%">
        <ul>
          <Index each={items()}>{(item) => <TodoItem {...item()} />}</Index>
        </ul>
      </Box>
    </div>
  );
};

export default TodoPage;
