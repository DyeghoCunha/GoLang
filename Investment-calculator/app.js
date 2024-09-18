const axios = require("axios");
const https = require("https");
const fs = require("fs");

const url = "https://165.227.140.4:4000/inventory?id=76561198077947014";
const headers = {
  Authorization: "test",
};

const agent = new https.Agent({
  rejectUnauthorized: false,
});

async function makeRequest() {
  try {
    const response = await axios.get(url, { headers, httpsAgent: agent });
    return response.data;
  } catch (error) {
    console.error("Error making request:", error);
    return null;
  }
}

async function fetchItems(count) {
  const items = [];
  for (let i = 0; i < count; i++) {
    const item = await makeRequest();
    if (item) {
      items.push(item[i]);
      // Armazenar o request em um arquivo JSON
    }
  }
  return items;
}

async function main() {
  const items = await fetchItems(100);
  fs.writeFileSync(`request_${2}.json`, JSON.stringify(items, null, 2));
  console.log(items.length);
  items.forEach((item, index) => {
    console.log(JSON.stringify(item, null, 2));
    console.log(
      "----------------------------------------------------------------------"
    );
    console.log(
      "---------------------------------------------------------------------------------"
    );
    console.log(
      "----------------------------------------------------------------------"
    );
  });
}

main();
