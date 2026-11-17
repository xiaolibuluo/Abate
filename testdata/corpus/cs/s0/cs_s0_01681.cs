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

    public interface IHandler15700
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15700(string Label, int Payload);

    public sealed class Handler15700 : IHandler15700
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier15701(int value);

    public sealed class Source15701
    {
        public event Notifier15701 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
