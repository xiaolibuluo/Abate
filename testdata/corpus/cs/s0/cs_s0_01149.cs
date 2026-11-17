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

    public static int Scale10780<T>(T value) where T : struct
    {
        int number = 51;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public delegate void Notifier10781(int value);

    public sealed class Source10781
    {
        public event Notifier10781 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
