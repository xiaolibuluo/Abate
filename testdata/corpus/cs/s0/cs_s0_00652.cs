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

    public interface IHandler6080
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6080(string Label, int Payload);

    public sealed class Handler6080 : IHandler6080
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier6081(int value);

    public sealed class Source6081
    {
        public event Notifier6081 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler6082
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6082(string Label, int Payload);

    public sealed class Handler6082 : IHandler6082
    {
        public void Handle(int value) { }
    }
    public static int Scale6083<T>(T value) where T : struct
    {
        int number = 73;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 28; }
        return number;
    }
}
