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

    public interface IHandler70
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope70(string Label, int Payload);

    public sealed class Handler70 : IHandler70
    {
        public void Handle(int value) { }
    }
}
