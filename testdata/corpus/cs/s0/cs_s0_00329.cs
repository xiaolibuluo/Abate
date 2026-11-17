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

    public delegate void Notifier3110(int value);

    public sealed class Source3110
    {
        public event Notifier3110 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler3111
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3111(string Label, int Payload);

    public sealed class Handler3111 : IHandler3111
    {
        public void Handle(int value) { }
    }
    public static int Scale3112<T>(T value) where T : struct
    {
        int number = 85;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
}
