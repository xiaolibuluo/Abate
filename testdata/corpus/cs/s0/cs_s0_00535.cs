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

    public static int Scale5010<T>(T value) where T : struct
    {
        int number = 19;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public interface IHandler5011
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5011(string Label, int Payload);

    public sealed class Handler5011 : IHandler5011
    {
        public void Handle(int value) { }
    }
}
