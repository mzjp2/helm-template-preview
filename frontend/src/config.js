const local = {
  api: "http://localhost:8001",
};

const prod = {
  api: "https://helm-preview.zainp.com/api",
};

const config = process.env.NODE_ENV === "production" ? prod : local;

export default config;
