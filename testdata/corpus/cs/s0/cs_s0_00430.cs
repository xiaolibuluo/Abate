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

    public interface IHandler4040
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4040(string Label, int Payload);

    public sealed class Handler4040 : IHandler4040
    {
        public void Handle(int value) { }
    }
    public static int Scale4041<T>(T value) where T : struct
    {
        int number = 26;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 42; }
        return number;
    }
}
