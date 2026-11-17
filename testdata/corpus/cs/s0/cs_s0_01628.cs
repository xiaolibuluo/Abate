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

    public static int Scale15230<T>(T value) where T : struct
    {
        int number = 16;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public delegate void Notifier15231(int value);

    public sealed class Source15231
    {
        public event Notifier15231 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier15232(int value);

    public sealed class Source15232
    {
        public event Notifier15232 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
