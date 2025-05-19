seems there is an issue with my dotnet installation:
```
export PATH=$PATH:$HOME/.dotnet
```
Got it! You’ll need to install .NET on your Linux machine first, since NUnit3 is a .NET testing framework. Here's how you can set up NUnit3 and Allure from scratch:

### Step 1: Install .NET SDK
First, install the .NET SDK on your Linux machine.

### Step 2: Create a New .NET Project
Once .NET is installed, create a new console application:

```bash
dotnet new console -n NUnitExample
cd NUnitExample
```

### Step 3: Install NUnit3 and NUnit Console Runner
Add NUnit3 and NUnit Console Runner to your project:

```bash
dotnet add package NUnit
dotnet add package NUnit3TestAdapter
dotnet add package Microsoft.NET.Test.Sdk
```

### Step 4: Install Allure CLI
Allure generates test reports. Install it with:

```bash
sudo apt update && sudo apt install allure
```

### Step 5: Write a Simple NUnit Test
Create a file `Test.cs` inside the `NUnitExample` project folder:

```csharp
using NUnit.Framework;

namespace NUnitExample
{
    public class Tests
    {
        [Test]
        public void TestExample()
        {
            Assert.AreEqual(2 + 2, 4);
        }
    }
}
```


after that based on this chagne the config:
https://allurereport.org/docs/nunit/


change it like this:
```csharp
using Allure.NUnit;
using NUnit.Framework;

[AllureNUnit]
class TestLabels
{
    [Test]
    public void TestCreateLabel()
    {
        // ...
    }
}
```
now run:
```
dotnet test
```
This will save necessary data into allure-results or other directory, according to the allure.directory setting. If the directory already exists, the new files will be added to the existing ones, so that a future report will be based on them all.
