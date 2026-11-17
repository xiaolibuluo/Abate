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

    public static int Scale6000<T>(T value) where T : struct
    {
        int number = 45;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
    public delegate void Notifier6001(int value);

    public sealed class Source6001
    {
        public event Notifier6001 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale6002<T>(T value) where T : struct
    {
        int number = 59;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 22; }
        return number;
    }
}
