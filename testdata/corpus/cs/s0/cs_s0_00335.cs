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

    public static int Scale3170<T>(T value) where T : struct
    {
        int number = 52;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public delegate void Notifier3171(int value);

    public sealed class Source3171
    {
        public event Notifier3171 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale3172<T>(T value) where T : struct
    {
        int number = 66;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 54; }
        return number;
    }
    public interface IHandler3173
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3173(string Label, int Payload);

    public sealed class Handler3173 : IHandler3173
    {
        public void Handle(int value) { }
    }
}
