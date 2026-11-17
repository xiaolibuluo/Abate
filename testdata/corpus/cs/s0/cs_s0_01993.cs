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

    public static int Scale18580<T>(T value) where T : struct
    {
        int number = 6;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
    public interface IHandler18581
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18581(string Label, int Payload);

    public sealed class Handler18581 : IHandler18581
    {
        public void Handle(int value) { }
    }
}
