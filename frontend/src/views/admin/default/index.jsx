import React, { useEffect, useState } from "react";
import { Box, SimpleGrid, useColorModeValue, Text } from "@chakra-ui/react";
import axios from "axios";
import CheckTable from "views/admin/default/components/CheckTable"; // Adjust the path if needed
import {
  columnsDataCheck, // Columns for questionOpts (id, question)
  columnsDataComplex, // Columns for questionEssays (question, answer)
} from "views/admin/default/variables/columnsData"; // Adjust the path if needed

export default function UserReports() {
  const [questionOpts, setQuestionOpts] = useState([]);
  const [questionEssays, setQuestionEssays] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState("");

  // Chakra Color Mode
  const brandColor = useColorModeValue("brand.500", "white");

  // Fetch data from /api/questionOpts and /api/questionEssays
  useEffect(() => {
    const fetchQuestions = async () => {
      try {
        setLoading(true);
        // Fetch question options
        const optsResponse = await axios.get("/api/questionOpts");
        setQuestionOpts(optsResponse.data);

        // Fetch question essays
        const essaysResponse = await axios.get("/api/questionEssays");
        setQuestionEssays(essaysResponse.data);

        setLoading(false);
      } catch (err) {
        setError("Failed to load data. Please try again later.");
        setLoading(false);
      }
    };

    fetchQuestions();
  }, []);

  return (
    <Box pt={{ base: "130px", md: "80px", xl: "80px" }}>
      {loading ? (
        <Text>Loading...</Text>
      ) : error ? (
        <Text color="red.500">{error}</Text>
      ) : (
        <SimpleGrid columns={{ base: 1, md: 1, xl: 2 }} gap="20px" mb="20px">
          {/* Table for Question Options with appropriate columns */}
          <CheckTable
            columnsData={columnsDataCheck} // Columns for questionOpts (id, question)
            tableData={questionOpts}
          />
          {/* Table for Question Essays with appropriate columns */}
          <CheckTable
            columnsData={columnsDataComplex} // Columns for questionEssays (question, answer)
            tableData={questionEssays}
          />
        </SimpleGrid>
      )}
    </Box>
  );
}
