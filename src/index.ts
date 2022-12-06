// *** Start *** //

import "@shared/start/index";
import "./db/index";

// *** Modules *** //

import logger from "jet-logger";
import mongoose from "mongoose";

// ??? My Import ??? //

import envVars from "@shared/Ienv";
import service from "./server";

// *** Constants *** //

const serverStartMsg = "Express server started on port: ";

const start = async () => {
  try {
    mongoose.connect(
      `mongodb://${envVars.mongo.username}:${envVars.mongo.password}@${envVars.mongo.hostname}:${envVars.mongo.port}/?authSource=${envVars.mongo.db_authsource}`
    );
    logger.info("DB Ok.");
    service.listen(envVars.port, () => {
      logger.imp(serverStartMsg + envVars.port);
    });
  } catch (err: any) {
    console.error(err);
  }
};

start();
