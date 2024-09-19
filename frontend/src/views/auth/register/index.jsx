import React, { useState } from "react";
import { NavLink, useNavigate } from "react-router-dom";
import {
  Box,
  Button,
  FormControl,
  FormLabel,
  Flex,
  Heading,
  Icon,
  Input,
  InputGroup,
  InputRightElement,
  Text,
  useColorModeValue,
} from "@chakra-ui/react";
import { MdOutlineRemoveRedEye } from "react-icons/md";
import { RiEyeCloseLine } from "react-icons/ri";
import DefaultAuth from "layouts/auth/Default";
import axios from "axios";

function Register() {
  const textColor = useColorModeValue("navy.700", "white");
  const textColorSecondary = "gray.400";
  const textColorBrand = useColorModeValue("brand.500", "white");
  const brandStars = useColorModeValue("brand.500", "brand.400");

  const [full_name, setFullName] = useState("");
  const [username, setUserName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState("");
  const [success, setSuccess] = useState("");

  const handleClickShowPassword = () => setShowPassword(!showPassword);
  const navigate = useNavigate();

  const handleRegister = async (e) => {
    e.preventDefault();
    if (password !== confirmPassword) {
      setError("Passwords do not match.");
      return;
    }

    try {
      const response = await axios.post("http://localhost:8080/api/register", {
        full_name,
        username,
        email,
        password,
      });

      if (response.status === 201) {
        setSuccess("Registration successful! Redirecting to login...");
        setTimeout(() => {
          navigate("/auth/sign-in");
        }, 2000); // Redirect to sign-in after 2 seconds
      }
    } catch (err) {
      setError("Error during registration. Please try again.");
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
            Register
          </Heading>
          <Text mb="36px" ms="4px" color={textColorSecondary} fontWeight="400" fontSize="md">
            Create an account with your email and password.
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
          <form onSubmit={handleRegister}>
            <FormControl>
              <FormLabel
                display="flex"
                ms="4px"
                fontSize="sm"
                fontWeight="500"
                color={textColor}
                mb="8px"
              >
                Full Name<Text color={brandStars}>*</Text>
              </FormLabel>
              <Input
                isRequired
                value={full_name}
                onChange={(e) => setFullName(e.target.value)}
                variant="auth"
                fontSize="sm"
                type="text"
                placeholder="John Doe"
                mb="24px"
                size="lg"
              />
              <FormLabel
                display="flex"
                ms="4px"
                fontSize="sm"
                fontWeight="500"
                color={textColor}
                mb="8px"
              >
                User Name<Text color={brandStars}>*</Text>
              </FormLabel>
              <Input
                isRequired
                value={username}
                onChange={(e) => setUserName(e.target.value)}
                variant="auth"
                fontSize="sm"
                type="text"
                placeholder="John Doe"
                mb="24px"
                size="lg"
              />
              <FormLabel
                display="flex"
                ms="4px"
                fontSize="sm"
                fontWeight="500"
                color={textColor}
                mb="8px"
              >
                Email<Text color={brandStars}>*</Text>
              </FormLabel>
              <Input
                isRequired
                value={email}
                onChange={(e) => setEmail(e.target.value)}
                variant="auth"
                fontSize="sm"
                type="email"
                placeholder="mail@domain.com"
                mb="24px"
                size="lg"
              />
              <FormLabel ms="4px" fontSize="sm" fontWeight="500" color={textColor}>
                Password<Text as="span" color={brandStars}>*</Text>
              </FormLabel>
              <InputGroup size="md" mb="24px">
                <Input
                  isRequired
                  value={password}
                  onChange={(e) => setPassword(e.target.value)}
                  type={showPassword ? "text" : "password"}
                  placeholder="Min. 8 characters"
                  size="lg"
                  variant="auth"
                />
                <InputRightElement>
                  <Icon
                    as={showPassword ? RiEyeCloseLine : MdOutlineRemoveRedEye}
                    onClick={handleClickShowPassword}
                    _hover={{ cursor: "pointer" }}
                  />
                </InputRightElement>
              </InputGroup>
              <FormLabel ms="4px" fontSize="sm" fontWeight="500" color={textColor}>
                Confirm Password<Text as="span" color={brandStars}>*</Text>
              </FormLabel>
              <InputGroup size="md">
                <Input
                  isRequired
                  value={confirmPassword}
                  onChange={(e) => setConfirmPassword(e.target.value)}
                  type={showPassword ? "text" : "password"}
                  placeholder="Re-enter your password"
                  size="lg"
                  variant="auth"
                />
                <InputRightElement>
                  <Icon
                    as={showPassword ? RiEyeCloseLine : MdOutlineRemoveRedEye}
                    onClick={handleClickShowPassword}
                    _hover={{ cursor: "pointer" }}
                  />
                </InputRightElement>
              </InputGroup>
              {error && (
                <Text color="red.500" mb="24px">
                  {error}
                </Text>
              )}
              {success && (
                <Text color="green.500" mb="24px">
                  {success}
                </Text>
              )}
              <Button type="submit" fontSize="sm" variant="brand" w="100%" h="50" mb="24px">
                Sign Up
              </Button>
            </FormControl>
          </form>
          <Flex flexDirection="column" justifyContent="center" alignItems="start" maxW="100%">
            <Text color={textColorSecondary} fontWeight="400" fontSize="14px">
              Already have an account?
              <NavLink to="/auth/sign-in">
                <Text color={textColorBrand} as="span" ms="5px" fontWeight="500">
                  Sign In
                </Text>
              </NavLink>
            </Text>
          </Flex>
        </Flex>
      </Flex>
    </DefaultAuth>
  );
}

export default Register;
