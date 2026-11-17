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

    public delegate void Notifier19100(int value);

    public sealed class Source19100
    {
        public event Notifier19100 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler19101
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19101(string Label, int Payload);

    public sealed class Handler19101 : IHandler19101
    {
        public void Handle(int value) { }
    }
}
