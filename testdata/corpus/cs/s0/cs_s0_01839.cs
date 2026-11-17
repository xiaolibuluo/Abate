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

    public delegate void Notifier17170(int value);

    public sealed class Source17170
    {
        public event Notifier17170 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier17171(int value);

    public sealed class Source17171
    {
        public event Notifier17171 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler17172
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17172(string Label, int Payload);

    public sealed class Handler17172 : IHandler17172
    {
        public void Handle(int value) { }
    }
}
