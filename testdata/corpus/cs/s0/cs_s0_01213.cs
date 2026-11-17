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

    public interface IHandler11380
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11380(string Label, int Payload);

    public sealed class Handler11380 : IHandler11380
    {
        public void Handle(int value) { }
    }
    public static int Scale11381<T>(T value) where T : struct
    {
        int number = 62;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public delegate void Notifier11382(int value);

    public sealed class Source11382
    {
        public event Notifier11382 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
