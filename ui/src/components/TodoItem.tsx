import {
  Box,
  Button,
  CloseButton,
  ListIcon,
  ListItem,
  Text,
  Radio,
  Input,
} from "@hope-ui/solid";
import { createEffect, createSignal, JSX, splitProps } from "solid-js";
import { put } from "../api/api";
import { TodoItemProps } from "../types";
const TodoItem = (
  props: TodoItemProps = { id: "", value: "", done: false }
) => {
  const [item, setItem] = createSignal<TodoItemProps | null>(null);
  const [isEdit, setIsEdit] = createSignal<boolean>(false);
  createEffect(() => {
    setItem(props);
  });
  const onInputHandler: JSX.EventHandlerUnion<HTMLInputElement, Event> = (
    e
  ) => {
    const current = item();
    if (current == null) return;
    setItem({
      id: current.id,
      value: (e.target as HTMLInputElement).value,
      done: current.done,
    });
  };

  const handleOnEdit = () => {
    setIsEdit(false);
    console.log(import.meta.env.VITE_HOST);
    console.log(JSON.stringify(item()));
    put("/api/items/edit", {
      body: JSON.stringify(item()),
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });
  };

  return (
    <Box as="li" style={{ display: "flex", "align-items": "center" }} ml="$3">
      <Radio colorScheme="primary" />
      {isEdit() ? (
        <>
          <Input
            type="text"
            value={item()?.value}
            onInput={(v) => onInputHandler(v)}
            style={{ width: "50%" }}
          />
          <Button size="sm" onClick={handleOnEdit}>
            確定
          </Button>
        </>
      ) : (
        <>
          <Text
            color="$neutral9"
            fontWeight="$bold"
            letterSpacing="$wide"
            fontSize="$2xl"
            textTransform="uppercase"
            m="$2"
          >
            {item()?.value}
          </Text>
          <Button size="sm" onClick={() => setIsEdit(true)}>
            編集
          </Button>
        </>
      )}

      {item()?.done && <CloseButton />}
    </Box>
  );
};

export default TodoItem;
