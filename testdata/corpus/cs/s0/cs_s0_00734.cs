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

    public delegate void Notifier6890(int value);

    public sealed class Source6890
    {
        public event Notifier6890 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler6891
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6891(string Label, int Payload);

    public sealed class Handler6891 : IHandler6891
    {
        public void Handle(int value) { }
    }
}
