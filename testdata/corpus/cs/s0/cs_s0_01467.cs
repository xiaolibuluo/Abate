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

    public static int Scale13750<T>(T value) where T : struct
    {
        int number = 32;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 57; }
        return number;
    }
    public interface IHandler13751
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13751(string Label, int Payload);

    public sealed class Handler13751 : IHandler13751
    {
        public void Handle(int value) { }
    }
}
