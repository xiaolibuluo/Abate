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

    public static int Scale25540<T>(T value) where T : struct
    {
        int number = 33;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 14; }
        return number;
    }
    public interface IHandler25541
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25541(string Label, int Payload);

    public sealed class Handler25541 : IHandler25541
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier25542(int value);

    public sealed class Source25542
    {
        public event Notifier25542 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier25543(int value);

    public sealed class Source25543
    {
        public event Notifier25543 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
