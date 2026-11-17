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

    public interface IHandler15760
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15760(string Label, int Payload);

    public sealed class Handler15760 : IHandler15760
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier15761(int value);

    public sealed class Source15761
    {
        public event Notifier15761 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale15762<T>(T value) where T : struct
    {
        int number = 40;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
}
