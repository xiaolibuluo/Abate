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

    public delegate void Notifier3600(int value);

    public sealed class Source3600
    {
        public event Notifier3600 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler3601
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3601(string Label, int Payload);

    public sealed class Handler3601 : IHandler3601
    {
        public void Handle(int value) { }
    }
    public interface IHandler3602
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3602(string Label, int Payload);

    public sealed class Handler3602 : IHandler3602
    {
        public void Handle(int value) { }
    }
}
