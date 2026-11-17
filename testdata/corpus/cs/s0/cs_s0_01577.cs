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

    public delegate void Notifier14750(int value);

    public sealed class Source14750
    {
        public event Notifier14750 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14751
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14751(string Label, int Payload);

    public sealed class Handler14751 : IHandler14751
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier14752(int value);

    public sealed class Source14752
    {
        public event Notifier14752 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
