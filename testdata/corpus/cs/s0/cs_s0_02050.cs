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

    public static int Scale19130<T>(T value) where T : struct
    {
        int number = 42;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
    public delegate void Notifier19131(int value);

    public sealed class Source19131
    {
        public event Notifier19131 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler19132
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19132(string Label, int Payload);

    public sealed class Handler19132 : IHandler19132
    {
        public void Handle(int value) { }
    }
}
