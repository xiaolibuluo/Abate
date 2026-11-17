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

    public static int Scale7410<T>(T value) where T : struct
    {
        int number = 35;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
    public interface IHandler7411
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7411(string Label, int Payload);

    public sealed class Handler7411 : IHandler7411
    {
        public void Handle(int value) { }
    }
}
