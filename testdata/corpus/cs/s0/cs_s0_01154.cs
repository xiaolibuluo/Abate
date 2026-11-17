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

    public delegate void Notifier10830(int value);

    public sealed class Source10830
    {
        public event Notifier10830 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier10831(int value);

    public sealed class Source10831
    {
        public event Notifier10831 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
