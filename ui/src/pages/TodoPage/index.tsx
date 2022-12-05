import { Box, Button, Heading, Input } from "@hope-ui/solid";
import { createEffect, createSignal, Index, JSX } from "solid-js";
import { v4 } from "uuid";
import { get, post } from "../../api/api";
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
    setItems([...items(), item]);
    setInputValue("");

    post("/api/items/", {
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
      body: JSON.stringify(item),
    });
  };

  const initItems = async () => {
    const data = await get("/api/items/", {
      "Content-Type": "application/json",
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    });
    const items: TodoItemProps[] = data.map((i: any) => {
      return {
        id: i.id,
        value: i.value,
        done: i.done,
      };
    });
    setItems(items);
  };

  createEffect(async () => {
    await initItems();
  });

  const onInputHandler: JSX.EventHandlerUnion<HTMLInputElement, Event> = (
    e
  ) => {
    setInputValue((e.target as HTMLInputElement).value);
  };

  return (
    <div>
      <Heading size="4xl">Todo</Heading>
      <div>
        <Input
          value={inputValues()}
          w="30%"
          type="text"
          onInput={(i) => {
            onInputHandler(i);
          }}
        />
        <Button
          disabled={inputValues().length === 0}
          size="sm"
          onClick={() => handleOnAdd(inputValues())}
        >
          追加
        </Button>
      </div>
      <Box ml="45%" mt="8%">
        <ul>
          <Index each={items()}>
            {(item) => (
              <TodoItem
                props={item()}
                initItems={initItems}
                items={items()}
                setItems={setItems}
              />
            )}
          </Index>
        </ul>
      </Box>
    </div>
  );
};

export default TodoPage;
