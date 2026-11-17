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

    public delegate void Notifier1550(int value);

    public sealed class Source1550
    {
        public event Notifier1550 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier1551(int value);

    public sealed class Source1551
    {
        public event Notifier1551 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
