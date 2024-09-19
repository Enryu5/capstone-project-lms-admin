import React, { useState, useContext } from "react";
import { NavLink, useNavigate } from "react-router-dom";
import {
  Box,
  Button,
  Flex,
  FormControl,
  FormLabel,
  Heading,
  Input,
  InputGroup,
  InputRightElement,
  Text,
  useColorModeValue,
  Icon, // Import the Icon component
} from "@chakra-ui/react";
import { MdOutlineRemoveRedEye } from "react-icons/md";
import { RiEyeCloseLine } from "react-icons/ri";
import DefaultAuth from "layouts/auth/Default";
import { AuthContext } from '../../../contexts/AuthContext'; // Correct AuthContext import
import axios from "axios";

function SignIn() {
  const textColor = useColorModeValue("navy.700", "white");
  const textColorSecondary = "gray.400";
  const textColorBrand = useColorModeValue("brand.500", "white");
  const brandStars = useColorModeValue("brand.500", "brand.400");

  const [username, setUserName] = useState("");
  const [password, setPassword] = useState("");
  const [show, setShow] = useState(false);
  const [error, setError] = useState("");
  const handleClick = () => setShow(!show);
  const navigate = useNavigate();

  const { login } = useContext(AuthContext); // Now using login from AuthContext

  const handleSignIn = async (e) => {
    e.preventDefault();
    try {
      // Call login from AuthContext
      await login(username, password);
      // If login is successful, redirect
      navigate("/admin/default"); // Redirect to dashboard
    } catch (err) {
      // If login fails, show error
      setError("Invalid username or password. Please try again.");
    }
  };

  return (
    <DefaultAuth>
      <Flex
        maxW={{ base: "100%", md: "max-content" }}
        w="100%"
        mx={{ base: "auto", lg: "0px" }}
        me="auto"
        h="100%"
        alignItems="start"
        justifyContent="center"
        mb={{ base: "30px", md: "60px" }}
        px={{ base: "25px", md: "0px" }}
        mt={{ base: "40px", md: "14vh" }}
        flexDirection="column"
      >
        <Box me="auto">
          <Heading color={textColor} fontSize="36px" mb="10px">
            Sign In
          </Heading>
          <Text mb="36px" ms="4px" color={textColorSecondary} fontWeight="400" fontSize="md">
            Enter your username and password to sign in!
          </Text>
        </Box>
        <Flex
          zIndex="2"
          direction="column"
          w={{ base: "100%", md: "420px" }}
          maxW="100%"
          background="transparent"
          borderRadius="15px"
          mx={{ base: "auto", lg: "unset" }}
          me="auto"
          mb={{ base: "20px", md: "auto" }}
        >
          <form onSubmit={handleSignIn}>
            <FormControl>
              <FormLabel
                display="flex"
                ms="4px"
                fontSize="sm"
                fontWeight="500"
                color={textColor}
                mb="8px"
              >
                Username<Text color={brandStars}>*</Text>
              </FormLabel>
              <Input
                isRequired
                value={username}
                onChange={(e) => setUserName(e.target.value)}
                variant="auth"
                fontSize="sm"
                type="username"
                placeholder="Username"
                mb="24px"
                size="lg"
              />
              <FormLabel ms="4px" fontSize="sm" fontWeight="500" color={textColor}>
                Password<Text as="span" color={brandStars}>*</Text>
              </FormLabel>
              <InputGroup size="md">
                <Input
                  isRequired
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  type={show ? "text" : "password"}
                  placeholder="Min. 8 characters"
                  size="lg"
                  variant="auth"
                  mb="24px"
                />
                <InputRightElement>
                  <Icon
                    as={show ? RiEyeCloseLine : MdOutlineRemoveRedEye}
                    onClick={handleClick}
                    _hover={{ cursor: "pointer" }}
                  />
                </InputRightElement>
              </InputGroup>
              {error && (
                <Text color="red.500" mb="24px">
                  {error}
                </Text>
              )}
              <Button type="submit" fontSize="sm" variant="brand" w="100%" h="50" mb="24px">
                Sign In
              </Button>
            </FormControl>
          </form>
          <Flex flexDirection="column" justifyContent="center" alignItems="start" maxW="100%">
            <Text color={textColorSecondary} fontWeight="400" fontSize="14px">
              Not registered yet?
              <NavLink to="/auth/register">
                <Text color={textColorBrand} as="span" ms="5px" fontWeight="500">
                  Create an Account
                </Text>
              </NavLink>
            </Text>
          </Flex>
        </Flex>
      </Flex>
    </DefaultAuth>
  );
}

export default SignIn;
