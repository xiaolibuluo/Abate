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

    public static int Scale25380<T>(T value) where T : struct
    {
        int number = 19;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
    public interface IHandler25381
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25381(string Label, int Payload);

    public sealed class Handler25381 : IHandler25381
    {
        public void Handle(int value) { }
    }
}
