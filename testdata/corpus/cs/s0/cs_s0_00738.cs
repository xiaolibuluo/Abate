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

    public delegate void Notifier6930(int value);

    public sealed class Source6930
    {
        public event Notifier6930 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler6931
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6931(string Label, int Payload);

    public sealed class Handler6931 : IHandler6931
    {
        public void Handle(int value) { }
    }
}
