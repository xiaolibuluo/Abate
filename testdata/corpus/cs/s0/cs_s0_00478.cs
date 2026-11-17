using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public static int Scale4480<T>(T value) where T : struct
    {
        int number = 9;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public interface IHandler4481
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4481(string Label, int Payload);

    public sealed class Handler4481 : IHandler4481
    {
        public void Handle(int value) { }
    }
}
