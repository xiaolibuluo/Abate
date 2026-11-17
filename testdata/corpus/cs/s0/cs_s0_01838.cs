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

    public delegate void Notifier17160(int value);

    public sealed class Source17160
    {
        public event Notifier17160 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler17161
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17161(string Label, int Payload);

    public sealed class Handler17161 : IHandler17161
    {
        public void Handle(int value) { }
    }
}
