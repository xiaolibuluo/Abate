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

    public interface IHandler9880
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9880(string Label, int Payload);

    public sealed class Handler9880 : IHandler9880
    {
        public void Handle(int value) { }
    }
    public static int Scale9881<T>(T value) where T : struct
    {
        int number = 52;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
    public interface IHandler9882
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9882(string Label, int Payload);

    public sealed class Handler9882 : IHandler9882
    {
        public void Handle(int value) { }
    }
}
