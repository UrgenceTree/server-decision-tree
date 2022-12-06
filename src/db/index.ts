// *** Modules *** //

import mongoose from "mongoose";

// *** Schema *** //

export const LongTermRental = new mongoose.Schema<any>(
  {
    Name: {
      type: String,
      required: true,
      unique: true,
    },
  },
  {
    timestamps: true,
  }
);

const MongoModLongTermRental = mongoose.model<any>(
  "long-term-rental",
  LongTermRental
);
