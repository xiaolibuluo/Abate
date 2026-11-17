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

    public static int Scale24060<T>(T value) where T : struct
    {
        int number = 49;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 15; }
        return number;
    }
    public static int Scale24061<T>(T value) where T : struct
    {
        int number = 56;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 26; }
        return number;
    }
    public interface IHandler24062
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24062(string Label, int Payload);

    public sealed class Handler24062 : IHandler24062
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier24063(int value);

    public sealed class Source24063
    {
        public event Notifier24063 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
