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

    public delegate void Notifier18100(int value);

    public sealed class Source18100
    {
        public event Notifier18100 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler18101
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18101(string Label, int Payload);

    public sealed class Handler18101 : IHandler18101
    {
        public void Handle(int value) { }
    }
}
