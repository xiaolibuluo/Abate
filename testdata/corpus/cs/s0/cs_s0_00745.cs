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

    public delegate void Notifier6990(int value);

    public sealed class Source6990
    {
        public event Notifier6990 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale6991<T>(T value) where T : struct
    {
        int number = 78;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 33; }
        return number;
    }
    public delegate void Notifier6992(int value);

    public sealed class Source6992
    {
        public event Notifier6992 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
