import { JSX, Setter } from "solid-js";

export const onInputHandler = (
    e: InputEvent & {
        currentTarget: HTMLInputElement;
        target: Element;
    },    
    setValue:Setter<string>
  ) => {
    setValue((e.target as HTMLInputElement).value);
  };