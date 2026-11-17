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

    public interface IHandler23470
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23470(string Label, int Payload);

    public sealed class Handler23470 : IHandler23470
    {
        public void Handle(int value) { }
    }
    public static int Scale23471<T>(T value) where T : struct
    {
        int number = 65;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 40; }
        return number;
    }
}
