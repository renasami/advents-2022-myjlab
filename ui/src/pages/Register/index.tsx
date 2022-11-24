import { Box, Button, Heading, Input } from "@hope-ui/solid";
import { createEffect, createSignal } from "solid-js";
import styles from "../../App.module.css";
import { onInputHandler } from "../../util/onInputHandler";
import "./style.css";
const Register = () => {
  const [isActive, setIsActive] = createSignal<boolean>(false);
  const [email, setEmail] = createSignal<string>("");
  const [password, setPassword] = createSignal<string>("");
  const [confirmPassword, setConfirmPassword] = createSignal<string>("");
  const format = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;

  createEffect(() => {
    const isEmailValid = email().match(format);
    const isPasswordsSame = password() === confirmPassword();
    const lengthIsValid = password().length > 6;
    setIsActive(
      isEmailValid && isPasswordsSame && lengthIsValid ? true : false
    );
  });

  return (
    <Box alignContent="center" justifyContent="center">
      <Heading size="4xl">Register</Heading>
      <Box maxW="$md" m="auto">
        <form>
          <Input
            value={email()}
            onInput={(e) => onInputHandler(e, setEmail)}
            type="email"
            name="register"
            placeholder="Email"
          />
          <Input
            value={password()}
            onInput={(e) => onInputHandler(e, setPassword)}
            type="password"
            name="password"
            placeholder="Password"
          />
          <Input
            value={confirmPassword()}
            onInput={(e) => onInputHandler(e, setConfirmPassword)}
            type="password"
            name="password"
            placeholder="Confirm"
          />
          <Button disabled={!isActive()}>Register</Button>
        </form>
      </Box>
    </Box>
  );
};

export default Register;
