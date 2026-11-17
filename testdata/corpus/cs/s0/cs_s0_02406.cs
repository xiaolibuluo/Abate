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

    public interface IHandler22470
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22470(string Label, int Payload);

    public sealed class Handler22470 : IHandler22470
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier22471(int value);

    public sealed class Source22471
    {
        public event Notifier22471 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
