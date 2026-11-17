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

    public delegate void Notifier22880(int value);

    public sealed class Source22880
    {
        public event Notifier22880 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
