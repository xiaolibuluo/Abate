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

    public interface IHandler22000
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22000(string Label, int Payload);

    public sealed class Handler22000 : IHandler22000
    {
        public void Handle(int value) { }
    }
    public interface IHandler22001
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22001(string Label, int Payload);

    public sealed class Handler22001 : IHandler22001
    {
        public void Handle(int value) { }
    }
}
