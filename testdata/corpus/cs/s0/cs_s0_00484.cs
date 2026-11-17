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

    public static int Scale4540<T>(T value) where T : struct
    {
        int number = 87;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public delegate void Notifier4541(int value);

    public sealed class Source4541
    {
        public event Notifier4541 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler4542
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4542(string Label, int Payload);

    public sealed class Handler4542 : IHandler4542
    {
        public void Handle(int value) { }
    }
}
