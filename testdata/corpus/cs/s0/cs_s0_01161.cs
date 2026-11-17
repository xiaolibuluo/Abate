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

    public interface IHandler10890
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10890(string Label, int Payload);

    public sealed class Handler10890 : IHandler10890
    {
        public void Handle(int value) { }
    }
    public static int Scale10891<T>(T value) where T : struct
    {
        int number = 7;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public interface IHandler10892
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10892(string Label, int Payload);

    public sealed class Handler10892 : IHandler10892
    {
        public void Handle(int value) { }
    }
}
