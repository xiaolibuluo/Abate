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

    public delegate void Notifier11300(int value);

    public sealed class Source11300
    {
        public event Notifier11300 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier11301(int value);

    public sealed class Source11301
    {
        public event Notifier11301 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
