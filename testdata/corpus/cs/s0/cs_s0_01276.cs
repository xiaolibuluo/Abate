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

    public interface IHandler11960
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11960(string Label, int Payload);

    public sealed class Handler11960 : IHandler11960
    {
        public void Handle(int value) { }
    }
    public interface IHandler11961
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11961(string Label, int Payload);

    public sealed class Handler11961 : IHandler11961
    {
        public void Handle(int value) { }
    }
    public static int Scale11962<T>(T value) where T : struct
    {
        int number = 47;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 46; }
        return number;
    }
    public interface IHandler11963
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11963(string Label, int Payload);

    public sealed class Handler11963 : IHandler11963
    {
        public void Handle(int value) { }
    }
}
