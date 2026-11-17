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

    public interface IHandler690
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope690(string Label, int Payload);

    public sealed class Handler690 : IHandler690
    {
        public void Handle(int value) { }
    }
    public static int Scale691<T>(T value) where T : struct
    {
        int number = 36;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public interface IHandler692
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope692(string Label, int Payload);

    public sealed class Handler692 : IHandler692
    {
        public void Handle(int value) { }
    }
}
