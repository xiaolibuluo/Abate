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

    public interface IHandler1110
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1110(string Label, int Payload);

    public sealed class Handler1110 : IHandler1110
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier1111(int value);

    public sealed class Source1111
    {
        public event Notifier1111 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
