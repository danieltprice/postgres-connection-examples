# TypeScript PostgreSQL Connection Example using pg-promise

This example shows you how to connect to PostgreSQL using pg-promise in TypeScript programming language.

## Prerequisites

Ensure you have the following installed on your machine:

- Node.js
- TypeScript

## Usage

1. Clone this repository:

```bash
git clone https://github.com/<your_username>/postgresql-connection-examples.git
```

2. Navigate to the TypeScript example directory:

```bash
cd postgresql-connection-examples/typescript/pg-promise
```
3. Install the necessary npm packages:

```bash
npm install pg-promise pg
```

4. Run the `tsc --init` command to create a `tsconfig.json` file in your project root directory, and add the content shown below to allow the required imports:

```bash
tsc --init
```

```json
{
  "compilerOptions": {
    "module": "commonjs",
    "esModuleInterop": true,
    "target": "es6",
    "noImplicitAny": true,
    "moduleResolution": "node",
    "outDir": "dist",
    "strict": true,
    "lib": ["es2017", "dom"]
  },
  "include": ["**/*.ts"],
  "exclude": ["node_modules"]
}

```

The `esModuleInterop` entry allows default imports from modules that use `export =`.

4. Compile the TypeScript file to JavaScript:

```bash
tsc index.ts
```

Or, if you are using `sslmode=verify-full`, run: 

```bash
tsc index-ssl.ts
```

5. Run the program:

```bash
node index.js
```

Or, if you used `index-ssl.ts`, run:

```bash
node index-ssl.js
```

This will print the PostgreSQL version and the current timestamp.

Always remember to handle database credentials securely in your applications. In a real-world scenario, you should not hardcode your credentials in the code. Instead, consider using environment variables or secure configuration files to handle such sensitive data.