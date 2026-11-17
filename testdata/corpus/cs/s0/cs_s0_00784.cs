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

    public delegate void Notifier7350(int value);

    public sealed class Source7350
    {
        public event Notifier7350 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
