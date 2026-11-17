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

    public static int Scale21040<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 36; }
        return number;
    }
    public interface IHandler21041
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21041(string Label, int Payload);

    public sealed class Handler21041 : IHandler21041
    {
        public void Handle(int value) { }
    }
}
