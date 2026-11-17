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

    public static int Scale22520<T>(T value) where T : struct
    {
        int number = 84;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 35; }
        return number;
    }
    public interface IHandler22521
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22521(string Label, int Payload);

    public sealed class Handler22521 : IHandler22521
    {
        public void Handle(int value) { }
    }
}
