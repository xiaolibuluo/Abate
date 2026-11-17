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

    public interface IHandler180
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope180(string Label, int Payload);

    public sealed class Handler180 : IHandler180
    {
        public void Handle(int value) { }
    }
    public static int Scale181<T>(T value) where T : struct
    {
        int number = 52;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
    public static int Scale182<T>(T value) where T : struct
    {
        int number = 59;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 35; }
        return number;
    }
}
