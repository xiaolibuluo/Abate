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

    public interface IHandler15270
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15270(string Label, int Payload);

    public sealed class Handler15270 : IHandler15270
    {
        public void Handle(int value) { }
    }
    public static int Scale15271<T>(T value) where T : struct
    {
        int number = 75;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public delegate void Notifier15272(int value);

    public sealed class Source15272
    {
        public event Notifier15272 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
