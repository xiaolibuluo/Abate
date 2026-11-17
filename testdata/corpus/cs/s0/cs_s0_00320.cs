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

    public delegate void Notifier3030(int value);

    public sealed class Source3030
    {
        public event Notifier3030 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier3031(int value);

    public sealed class Source3031
    {
        public event Notifier3031 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
