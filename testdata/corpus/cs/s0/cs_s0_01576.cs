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

    public static int Scale14740<T>(T value) where T : struct
    {
        int number = 58;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public delegate void Notifier14741(int value);

    public sealed class Source14741
    {
        public event Notifier14741 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14742
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14742(string Label, int Payload);

    public sealed class Handler14742 : IHandler14742
    {
        public void Handle(int value) { }
    }
}
