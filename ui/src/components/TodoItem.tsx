import {
  Box,
  Button,
  CloseButton,
  ListIcon,
  ListItem,
  Text,
  Radio,
  Input,
  Checkbox,
  IconButton,
} from "@hope-ui/solid";
import {
  Component,
  createEffect,
  createSignal,
  JSX,
  Setter,
  splitProps,
} from "solid-js";
import { del, put } from "../api/api";
import { TodoItemProps } from "../types";
import { AiOutlineDelete } from "solid-icons/ai";

type props = {
  props: TodoItemProps;
  initItems: () => Promise<void>;
  items: TodoItemProps[];
  setItems: Setter<TodoItemProps[]>;
};
const TodoItem: Component<props> = ({ props, initItems, items, setItems }) => {
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

  const setIsDone = () => {
    const current = item();
    if (current == null) return;
    setItem({
      id: current.id,
      value: current.value,
      done: !current.done,
    });
    handleOnEdit();
  };
  const handleOnDelete = async () => {
    const newArr = items.filter((i) => i.id !== item()?.id);
    console.log(newArr);
    setItems(newArr);
    await del(`/api/items/delete?id=${item()?.id}`, {
      "Content-Type": "application/json",
      Authorization: `Bearer ${localStorage.getItem("token")}`,
    });
  };

  const handleOnEdit = () => {
    setIsEdit(false);
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
      <Checkbox
        checked={item()?.done}
        onChange={setIsDone}
        colorScheme="primary"
      />
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
          {item()?.done ? (
            <>
              <Text
                color="$neutral9"
                fontWeight="$bold"
                letterSpacing="$wide"
                fontSize="$2xl"
                textTransform="uppercase"
                textDecoration="line-through"
                m="$2"
              >
                {item()?.value}
              </Text>
              <IconButton
                colorScheme="primary"
                aria-label="Delete"
                icon={<AiOutlineDelete />}
                onClick={handleOnDelete}
              />
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
        </>
      )}
    </Box>
  );
};

export default TodoItem;
