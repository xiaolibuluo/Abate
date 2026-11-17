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

    public interface IHandler14690
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14690(string Label, int Payload);

    public sealed class Handler14690 : IHandler14690
    {
        public void Handle(int value) { }
    }
    public static int Scale14691<T>(T value) where T : struct
    {
        int number = 97;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 55; }
        return number;
    }
}
