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

    public delegate void Notifier8880(int value);

    public sealed class Source8880
    {
        public event Notifier8880 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier8881(int value);

    public sealed class Source8881
    {
        public event Notifier8881 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
