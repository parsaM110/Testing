using NUnit.Framework;
using Allure.NUnit;

namespace NUnitExample
{
    [AllureNUnit]
    public class Tests
    {
        [Test]
        public void TestExample()
        {
            Assert.That(2 + 2, Is.EqualTo(4));
        }
    }
}
