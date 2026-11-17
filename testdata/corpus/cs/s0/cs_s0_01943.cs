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

    public interface IHandler18130
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18130(string Label, int Payload);

    public sealed class Handler18130 : IHandler18130
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier18131(int value);

    public sealed class Source18131
    {
        public event Notifier18131 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
