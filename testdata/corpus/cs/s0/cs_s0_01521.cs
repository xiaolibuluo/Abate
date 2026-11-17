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

    public static int Scale14250<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
    public interface IHandler14251
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14251(string Label, int Payload);

    public sealed class Handler14251 : IHandler14251
    {
        public void Handle(int value) { }
    }
}
