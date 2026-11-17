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

    public static int Scale12300<T>(T value) where T : struct
    {
        int number = 87;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public interface IHandler12301
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12301(string Label, int Payload);

    public sealed class Handler12301 : IHandler12301
    {
        public void Handle(int value) { }
    }
}
