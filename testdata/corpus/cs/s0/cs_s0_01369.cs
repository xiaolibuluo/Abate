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

    public static int Scale12830<T>(T value) where T : struct
    {
        int number = 97;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 23; }
        return number;
    }
    public static int Scale12831<T>(T value) where T : struct
    {
        int number = 7;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public delegate void Notifier12832(int value);

    public sealed class Source12832
    {
        public event Notifier12832 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
