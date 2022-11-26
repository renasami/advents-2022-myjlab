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
  import { InferType, object, string } from "yup";
  import { createForm } from "@felte/solid";
  const schema = object({
    email: string().email().required(),
    password: string().min(6).max(12).required(),
  });
  
  const Login = () => {
    const [isActive, setIsActive] = createSignal<boolean>(false);
    const [email, setEmail] = createSignal<string>("");
    const [password, setPassword] = createSignal<string>("testtest");
  
    const format = /^\w+([\.-]?\w+)*@\w+([\.-]?\w+)*(\.\w{2,3})+$/;
  
    createEffect(() => {
      const isEmailValid = email().match(format);
      const lengthIsValid = password().length > 6;
      setIsActive(
        isEmailValid && lengthIsValid ? true : false
      );
    });
  
    const { form, errors, data, isValid, setFields } = createForm<
      InferType<typeof schema>
    >({
      extend: validator({ schema }),
      onSubmit: async (values) => {
        console.log(values);
        const postBody = {
          email: values.email,
          password: values.password,
        };

        await fetch("http://127.0.0.1:8080/api/auth/login", {
            method: "POST",
            body: JSON.stringify(postBody),
            headers: { "Content-Type": "Application-JSON" }
          } );
      },
    });
  
    return (
      <Box alignContent="center" justifyContent="center">
        <Heading size="4xl">Login</Heading>
        <Box maxW="$md" m="auto" paddingTop="$10">
          <VStack
            as="form"
            ref={form}
            spacing="$5"
            alignItems="stretch"
            maxW="$96"
            mx="auto"
          >
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
            <Button type="submit" disabled={!isActive()}>
              Login
            </Button>
          </VStack>
        </Box>
      </Box>
    );
  };
  
  export default Login;
  