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

    public static int Scale22570<T>(T value) where T : struct
    {
        int number = 52;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public delegate void Notifier22571(int value);

    public sealed class Source22571
    {
        public event Notifier22571 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler22572
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22572(string Label, int Payload);

    public sealed class Handler22572 : IHandler22572
    {
        public void Handle(int value) { }
    }
}
