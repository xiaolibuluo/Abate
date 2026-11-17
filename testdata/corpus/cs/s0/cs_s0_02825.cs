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

    public interface IHandler26420
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26420(string Label, int Payload);

    public sealed class Handler26420 : IHandler26420
    {
        public void Handle(int value) { }
    }
    public static int Scale26421<T>(T value) where T : struct
    {
        int number = 20;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 31; }
        return number;
    }
}
