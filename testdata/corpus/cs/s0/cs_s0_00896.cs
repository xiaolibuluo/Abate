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

    public delegate void Notifier8410(int value);

    public sealed class Source8410
    {
        public event Notifier8410 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler8411
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8411(string Label, int Payload);

    public sealed class Handler8411 : IHandler8411
    {
        public void Handle(int value) { }
    }
}
