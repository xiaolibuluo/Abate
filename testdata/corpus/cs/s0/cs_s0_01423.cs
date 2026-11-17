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

    public delegate void Notifier13330(int value);

    public sealed class Source13330
    {
        public event Notifier13330 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale13331<T>(T value) where T : struct
    {
        int number = 75;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public interface IHandler13332
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13332(string Label, int Payload);

    public sealed class Handler13332 : IHandler13332
    {
        public void Handle(int value) { }
    }
}
