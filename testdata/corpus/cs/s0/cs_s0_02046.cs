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

    public interface IHandler19090
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19090(string Label, int Payload);

    public sealed class Handler19090 : IHandler19090
    {
        public void Handle(int value) { }
    }
    public interface IHandler19091
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19091(string Label, int Payload);

    public sealed class Handler19091 : IHandler19091
    {
        public void Handle(int value) { }
    }
}
