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

    public interface IHandler7800
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7800(string Label, int Payload);

    public sealed class Handler7800 : IHandler7800
    {
        public void Handle(int value) { }
    }
}
