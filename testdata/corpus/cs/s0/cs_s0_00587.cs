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

    public interface IHandler5500
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5500(string Label, int Payload);

    public sealed class Handler5500 : IHandler5500
    {
        public void Handle(int value) { }
    }
    public static int Scale5501<T>(T value) where T : struct
    {
        int number = 81;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 27; }
        return number;
    }
}
