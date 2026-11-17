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

    public delegate void Notifier21670(int value);

    public sealed class Source21670
    {
        public event Notifier21670 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier21671(int value);

    public sealed class Source21671
    {
        public event Notifier21671 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier21672(int value);

    public sealed class Source21672
    {
        public event Notifier21672 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler21673
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21673(string Label, int Payload);

    public sealed class Handler21673 : IHandler21673
    {
        public void Handle(int value) { }
    }
    public interface IHandler21674
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope21674(string Label, int Payload);

    public sealed class Handler21674 : IHandler21674
    {
        public void Handle(int value) { }
    }
}
