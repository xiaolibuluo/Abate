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

    public interface IHandler25820
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25820(string Label, int Payload);

    public sealed class Handler25820 : IHandler25820
    {
        public void Handle(int value) { }
    }
}
