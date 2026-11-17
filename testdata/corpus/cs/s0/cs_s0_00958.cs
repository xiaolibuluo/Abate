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

    public static int Scale8980<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
    public interface IHandler8981
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8981(string Label, int Payload);

    public sealed class Handler8981 : IHandler8981
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier8982(int value);

    public sealed class Source8982
    {
        public event Notifier8982 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale8983<T>(T value) where T : struct
    {
        int number = 60;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 45; }
        return number;
    }
}
