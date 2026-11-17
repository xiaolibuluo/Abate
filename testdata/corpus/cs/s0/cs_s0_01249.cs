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

    public interface IHandler11710
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11710(string Label, int Payload);

    public sealed class Handler11710 : IHandler11710
    {
        public void Handle(int value) { }
    }
}
