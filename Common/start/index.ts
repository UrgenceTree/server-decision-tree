// *** Modules *** //

import logger from "jet-logger";
import path, { dirname } from "path";
import dotenv from "dotenv";
import { parse } from "ts-command-line-args";

// ??? Interface ??? //

interface ParserArguments {
  env: "development" | "production" | string;
}

// *** Variables *** //

const args = parse<ParserArguments>({
  env: {
    alias: "e",
    defaultValue: "development",
    type: String,
  },
});

// *** Get Root Project *** //

if (require.main === undefined) throw new Error("Not Root Project find.");

const appDir = dirname(require.main.filename);

// *** Set environment *** //

const TakeEnvVariable = dotenv.config({
  path: path.join(appDir, `env/${String(args.env)}.env`),
});

if (TakeEnvVariable.error) throw TakeEnvVariable.error;

if (args.env === "development") {
  logger.imp("development mode");
} else logger.imp("production mode");
