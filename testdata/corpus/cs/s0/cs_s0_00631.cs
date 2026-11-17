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

    public delegate void Notifier5890(int value);

    public sealed class Source5890
    {
        public event Notifier5890 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
