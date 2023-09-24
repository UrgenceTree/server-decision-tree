// ??? interface ??? //

interface IenvVars {
  nodeEnv: string;
  /**
   * @brief The port were we lisen
   */
  port: number | string;
  /**
   * @development Shows in which middleware we navigate
   */
  show_route: boolean | string;
  mongo: {
    hostname: string;
    username: string;
    password: string;
    port: number;
    db_authsource: string;
  };
}

const envVars: IenvVars = {
  nodeEnv: process.env.NODE_ENV ?? "",
  port: process.env.PORT ?? 3000,
  show_route: process.env.SHOW_ROUTES ?? false,
  mongo: {
    hostname: process.env.DB_HOST ?? "localhost",
    username: process.env.DB_USR ?? "",
    password: process.env.DB_PWD ?? "",
    port: parseInt(process.env.DB_PORT ?? "27017"),
    db_authsource: process.env.DB_AUTH_SOURCE ?? "",
  },
};

export default envVars;
