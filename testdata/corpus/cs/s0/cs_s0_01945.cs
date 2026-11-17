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

    public interface IHandler18150
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18150(string Label, int Payload);

    public sealed class Handler18150 : IHandler18150
    {
        public void Handle(int value) { }
    }
    public interface IHandler18151
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18151(string Label, int Payload);

    public sealed class Handler18151 : IHandler18151
    {
        public void Handle(int value) { }
    }
}
