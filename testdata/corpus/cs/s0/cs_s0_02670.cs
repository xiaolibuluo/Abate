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

    public interface IHandler24970
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24970(string Label, int Payload);

    public sealed class Handler24970 : IHandler24970
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier24971(int value);

    public sealed class Source24971
    {
        public event Notifier24971 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler24972
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24972(string Label, int Payload);

    public sealed class Handler24972 : IHandler24972
    {
        public void Handle(int value) { }
    }
}
