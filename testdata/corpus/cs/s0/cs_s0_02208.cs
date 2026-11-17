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

    public static int Scale20620<T>(T value) where T : struct
    {
        int number = 39;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public delegate void Notifier20621(int value);

    public sealed class Source20621
    {
        public event Notifier20621 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler20622
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20622(string Label, int Payload);

    public sealed class Handler20622 : IHandler20622
    {
        public void Handle(int value) { }
    }
}
