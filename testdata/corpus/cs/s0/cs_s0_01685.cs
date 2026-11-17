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

    public static int Scale15740<T>(T value) where T : struct
    {
        int number = 97;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public interface IHandler15741
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15741(string Label, int Payload);

    public sealed class Handler15741 : IHandler15741
    {
        public void Handle(int value) { }
    }
    public interface IHandler15742
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15742(string Label, int Payload);

    public sealed class Handler15742 : IHandler15742
    {
        public void Handle(int value) { }
    }
}
