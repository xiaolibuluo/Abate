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

    public delegate void Notifier24000(int value);

    public sealed class Source24000
    {
        public event Notifier24000 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier24001(int value);

    public sealed class Source24001
    {
        public event Notifier24001 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler24002
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24002(string Label, int Payload);

    public sealed class Handler24002 : IHandler24002
    {
        public void Handle(int value) { }
    }
}
