
### Step 1: Initialize Your Project

If you haven't already, create a new project and navigate into it:

```bash
mkdir my-vitest-project
cd my-vitest-project
npm init -y
```

### Step 2: Install Dependencies

You need to install Vitest and Allure. You can do this using npm:

```bash
npm install --save-dev vitest @vitest/allure
```

### Step 3: Configure Vitest

Create a configuration file for Vitest. You can name it `vitest.config.ts` or `vitest.config.js` depending on whether you're using TypeScript or JavaScript.

Here’s a basic example of a `vitest.config.ts` file:

```typescript
import { defineConfig } from 'vitest/config';
import { allure } from '@vitest/allure';

export default defineConfig({
  test: {
    // Add Allure reporter
    reporters: [
      'default',
      allure({
        outputDir: 'allure-results', // Directory to store allure results
      }),
    ],
    // Other configurations can go here
  },
});
```

### Step 4: Write a Test

Create a test file, for example, `example.test.ts`:

```typescript
import { describe, it, expect } from 'vitest';

describe('Example Test', () => {
  it('should add numbers correctly', () => {
    const sum = 1 + 1;
    expect(sum).toBe(2);
  });
});
```

### Step 5: Run Your Tests

You can run your tests using the following command:

```bash
npx vitest
```

### Step 6: Generate Allure Reports

After running your tests, you will have results in the `allure-results` directory. To generate the Allure report, you need to have Allure Commandline installed. You can install it via Homebrew on macOS:

```bash
brew install allure
```

Or you can download it from the [Allure website](https://docs.qameta.io/allure/).

Once you have Allure installed, you can generate the report with:

```bash
allure serve allure-results
```
or
```
allure serve allure-results --host 0.0.0.0
```

This command will start a local server and open the Allure report in your default web browser.
