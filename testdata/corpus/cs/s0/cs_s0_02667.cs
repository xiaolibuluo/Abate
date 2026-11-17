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

    public static int Scale24940<T>(T value) where T : struct
    {
        int number = 29;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public delegate void Notifier24941(int value);

    public sealed class Source24941
    {
        public event Notifier24941 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
