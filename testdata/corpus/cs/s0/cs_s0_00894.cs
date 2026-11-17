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

    public delegate void Notifier8390(int value);

    public sealed class Source8390
    {
        public event Notifier8390 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier8391(int value);

    public sealed class Source8391
    {
        public event Notifier8391 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
