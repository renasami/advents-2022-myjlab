import { Component, createSignal, Index, JSX } from "solid-js";
import { v4 } from "uuid";
import logo from "./logo.svg";
import styles from "./App.module.css";

type Item = {
  id: string;
  value: string;
};
const App: Component = () => {
  const [items, setItems] = createSignal<Item[]>([]);
  const [inputValues, setInputValue] = createSignal<string>("");
  const markAsFin = (uuid: string) => {
    const current = items();
    const newValue = current.filter((item) => item.id !== uuid);
    setItems(newValue);
  };

  const handleOnGenerate = (value: any) => {
    const item: Item = {
      id: v4().toString(),
      value: value,
    };
    setItems([...items(), item]);
    setInputValue("");
  };
  const onInputHandler: JSX.EventHandlerUnion<HTMLInputElement, Event> = (
    e
  ) => {
    setInputValue((e.target as HTMLInputElement).value);
  };

  return (
    <div class={styles.App}>
      <header class={styles.header}>
        <img src={logo} class={styles.logo} alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          class={styles.link}
          href="https://github.com/solidjs/solid"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn Solid
        </a>
        <ul>
          <Index each={items()}>
            {(item) => (
              <li>
                <span>{item().value}</span>
                <button onClick={() => markAsFin(item().id)}>完了</button>
              </li>
            )}
          </Index>
        </ul>

        <div>
          <input
            value={inputValues()}
            type="text"
            onInput={(i) => {
              onInputHandler(i);
            }}
          />
          <button onClick={() => handleOnGenerate(inputValues())}>追加</button>
        </div>
      </header>
    </div>
  );
};

export default App;
