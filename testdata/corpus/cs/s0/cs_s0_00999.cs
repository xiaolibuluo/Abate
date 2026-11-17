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

    public delegate void Notifier9370(int value);

    public sealed class Source9370
    {
        public event Notifier9370 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier9371(int value);

    public sealed class Source9371
    {
        public event Notifier9371 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
