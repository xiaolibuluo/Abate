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

    public static int Scale3560<T>(T value) where T : struct
    {
        int number = 74;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
    public interface IHandler3561
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3561(string Label, int Payload);

    public sealed class Handler3561 : IHandler3561
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier3562(int value);

    public sealed class Source3562
    {
        public event Notifier3562 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
