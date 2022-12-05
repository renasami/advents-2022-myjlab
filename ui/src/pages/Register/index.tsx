import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Heading,
  Input,
  VStack,
} from "@hope-ui/solid";
import { createEffect, createSignal } from "solid-js";
import { validator } from "@felte/validator-yup";
import { onInputHandler } from "../../util/onInputHandler";
import "./style.css";
import { InferType, object, string } from "yup";
import { createForm } from "@felte/solid";
import { post } from "../../api/api";
const schema = object({
  username: string().min(3).max(8).required(),
  email: string().email().required(),
  password: string().min(6).max(12).required(),
});

const Register = () => {
  const [isActive, setIsActive] = createSignal<boolean>(false);
  const [email, setEmail] = createSignal<string>("");
  const [username, setUsername] = createSignal<string>("testuser");
  const [password, setPassword] = createSignal<string>("testtest");
  const [confirmPassword, setConfirmPassword] =
    createSignal<string>("testtest");

  const format = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;

  createEffect(() => {
    const isEmailValid = email().match(format);
    const isPasswordsSame = password() === confirmPassword();
    const lengthIsValid = password().length > 6;
    setIsActive(
      isEmailValid && isPasswordsSame && lengthIsValid ? true : false
    );
  });

  const { form, errors, data, isValid, setFields } = createForm<
    InferType<typeof schema>
  >({
    extend: validator({ schema }),
    onSubmit: async (values) => {
      alert("values");
      console.log(values);
      const postBody = {
        username: values.username,
        email: values.email,
        password: values.password,
      };
      const options = {
        body: JSON.stringify(postBody),
        headers: { "Content-Type": "Application-JSON" },
      };
      await post("/api/auth/register", options);
    },
  });

  return (
    <Box alignContent="center" justifyContent="center">
      <Heading size="4xl">Register</Heading>
      <Box maxW="$md" m="auto">
        <VStack
          as="form"
          ref={form}
          spacing="$5"
          alignItems="stretch"
          maxW="$96"
          mx="auto"
        >
          <FormControl required invalid={!!errors("username")}>
            <Input
              value={username()}
              onInput={(e) => onInputHandler(e, setUsername)}
              type="text"
              name="username"
              placeholder="Username"
            />
          </FormControl>
          <FormControl required invalid={!!errors("email")}>
            <Input
              value={email()}
              onInput={(e) => onInputHandler(e, setEmail)}
              type="email"
              name="email"
              placeholder="Email"
            />
          </FormControl>
          <FormControl>
            <Input
              value={password()}
              onInput={(e) => onInputHandler(e, setPassword)}
              type="password"
              name="password"
              placeholder="Password"
            />
          </FormControl>
          <FormControl>
            <Input
              value={confirmPassword()}
              onInput={(e) => onInputHandler(e, setConfirmPassword)}
              type="password"
              name="password"
              placeholder="Confirm"
            />
          </FormControl>
          <Button type="submit" disabled={!isActive()}>
            Register
          </Button>
        </VStack>
      </Box>
    </Box>
  );
};

export default Register;
