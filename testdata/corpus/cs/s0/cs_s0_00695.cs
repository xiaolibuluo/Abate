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

    public static int Scale6500<T>(T value) where T : struct
    {
        int number = 16;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 45; }
        return number;
    }
    public delegate void Notifier6501(int value);

    public sealed class Source6501
    {
        public event Notifier6501 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler6502
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6502(string Label, int Payload);

    public sealed class Handler6502 : IHandler6502
    {
        public void Handle(int value) { }
    }
}
