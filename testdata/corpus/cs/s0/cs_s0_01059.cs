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

    public delegate void Notifier9910(int value);

    public sealed class Source9910
    {
        public event Notifier9910 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier9911(int value);

    public sealed class Source9911
    {
        public event Notifier9911 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler9912
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope9912(string Label, int Payload);

    public sealed class Handler9912 : IHandler9912
    {
        public void Handle(int value) { }
    }
}
